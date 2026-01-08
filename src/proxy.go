package main

import (
	"io"
	"net"
	"log"
	"net/http"
	"time"
)

type Proxy struct{}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.Host, r.URL.String())
	if r.Method == http.MethodConnect {
		p.handlehttps(w, r)
	} else {
		p.handlehttp(w, r)
	}
}

func (p *Proxy) handlehttp(w http.ResponseWriter, r *http.Request) {
	normalizeheaders(r)

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	copyheaders(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func (p *Proxy) handlehttps(w http.ResponseWriter, r *http.Request) {
	destConn, err := net.DialTimeout("tcp", r.Host, 10*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}

	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		return
	}

	clientConn.Write([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))

	go tunnel(destConn, clientConn)
	go tunnel(clientConn, destConn)
}

func tunnel(dst net.Conn, src net.Conn) {
	defer dst.Close()
	defer src.Close()
	io.Copy(dst, src)
}

func copyheaders(dst, src http.Header) {
	for k, v := range src {
		dst[k] = v
	}
}

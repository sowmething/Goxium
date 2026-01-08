package main

import "net/http"

var (
	HeadersToDelete []string
	HeaderSetKey    string
	HeaderSetVal    string
)

func normalizeheaders(r *http.Request) {
	r.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	r.Header.Set("Accept-Language", "en-US,en;q=0.9")
	r.Header.Del("Sec-CH-UA")
	r.Header.Del("Via")
	r.Header.Set("DNT", "1")
	r.Header.Del("Sec-CH-UA-Platform")
	r.Header.Del("Sec-CH-UA-Mobile")
	r.Header.Set("Connection", "keep-alive")


	for _, h := range HeadersToDelete {
		r.Header.Del(h)
	}

	if HeaderSetKey != "" {
		r.Header.Set(HeaderSetKey, HeaderSetVal)
	}
}
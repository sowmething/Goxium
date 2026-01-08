package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"os/exec"
	"syscall"
	"github.com/elazarl/goproxy"
)

const (
	Green = "\033[32m"
	Reset = "\033[0m"
)


var (
	user32           = syscall.NewLazyDLL("user32.dll")
	messageBox       = user32.NewProc("MessageBoxW")
	kernel32         = syscall.NewLazyDLL("kernel32.dll")
	setConsoleCtrl   = kernel32.NewProc("SetConsoleCtrlHandler")
)

func addcert() {
	log.Println(Green + "[+]" + Reset + " Adding certificate")
	exec.Command("certutil", "-addstore", "Root", "goxium_ca.crt").Run()
}

func handler(dwCtrlType uintptr) uintptr {
	log.Println(Green + "[+]" + Reset + " Removing certificate")
	exec.Command("certutil", "-delstore", "Root", "Goxium Root CA").Run()
	os.Exit(0)
	return 1
}

func bannerandeffect() {
	goxium := []string{
		".d8888b.                     d8b",
		"d88P  Y88b                   Y8P",
		"888    888",
		"888         .d88b.  888  888 888 888  888 88888b.d88b.",
		"888  88888 d88  88b  Y8bd8P  888 888  888 888  888  88b",
		"888    888 888  888   X88K   888 888  888 888  888  888",
		"Y88b  d88P Y88..88P .d8  8b. 888 Y88b 888 888  888  888",
		" Y8888P88    Y88P   888  888 888   Y88888 888  888  888",
		"",
		" System-wide browser fingerprint hiding utility",
		" Made by https://github.com/sowmething",
	}

	colors := []int{141, 135, 129, 93, 55}

	for i, line := range goxium {
		color := colors[i*len(colors)/len(goxium)] // hard to type
		log.Printf("\033[38;5;%dm%s\033[0m", color, line)
	}
}


func main() {
	
	bannerandeffect()
	
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-del":
			for j := i + 1; j < len(args); j++ { // omg
				if args[j][0] == '-' {
					break
				}
				HeadersToDelete = append(HeadersToDelete, args[j])
				i = j 
			}
		case "-set":
			if i+2 < len(args) {
				HeaderSetKey = args[i+1]
				HeaderSetVal = args[i+2]
				i += 2 
			}
		}
	}
	
	exec.Command("certutil", "-addstore", "-f", "Root", "goxium_ca.crt").Run()

	setConsoleCtrl.Call(syscall.NewCallback(handler), 1)

	proxy := goproxy.NewProxyHttpServer()

	caCert, err := os.ReadFile("goxium_ca.crt")
	caKey, err := os.ReadFile("goxium_ca.key")
	if err == nil {
		goproxyCa, err := tls.X509KeyPair(caCert, caKey)
		if err == nil {
			goproxy.GoproxyCa = goproxyCa
			log.Println(Green + "[+]" + Reset + " Loaded certificate")
		}
	}

	proxy.Verbose = true
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	proxy.OnRequest().DoFunc(func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		normalizeheaders(r)
		log.Println(r.Method, r.Host, r.URL.String())
		return r, nil
	})

	log.Println(Green + "[+]" + Reset + " Goxium listening on 127.0.0.1:19381")
	log.Fatal(http.ListenAndServe("127.0.0.1:19381", proxy))
}
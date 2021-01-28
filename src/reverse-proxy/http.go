package main

import (
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func startHTTPServer(f *functions) {

	mux := http.NewServeMux()

	director := func(r *http.Request) {
		f.RLock()
		defer f.RUnlock()

		p := r.URL.Path

		for p != "" && p[0] == '/' {
			p = p[1:]
		}

		handler := ""
		for path := range f.hosts {

			if strings.HasPrefix(r.URL.Path, "/"+path) && len(handler) <= len(path) {
				handler = path
			}
		}

		urls, ok := f.hosts[handler]

		if ok {
			dest, _ := url.Parse("http://" + urls[rand.Intn(len(urls))] + ":8000")

			r.URL.Host = dest.Host
			r.URL.Path = strings.Replace(r.URL.Path, "/"+handler, "", 1)
		}
		r.URL.Scheme = "http"
	}

	proxy := &httputil.ReverseProxy{Director: director}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	server := http.Server{Addr: ":7000", Handler: mux}
	server.SetKeepAlivesEnabled(false)
	server.ListenAndServe()

}

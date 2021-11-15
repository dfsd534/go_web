package main

import (
	"github.com/dfsd534/go_web/data"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static", http.StripPrefix("/static", files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8081",
		Handler: mux,
	}

	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	if threads, err := data.Threads(); err == nil {
		_, err := session(w, r)
		public_tmpl_files := []string{"layout", "public.navbar", "index"}
		private_tmpl_files := []string{"layout", "private.navbar", "index"}

		if err != nil {
			generateHTML(w, threads, public_tmpl_files...)
		} else {
			generateHTML(w, threads, private_tmpl_files...)
		}
	}
}

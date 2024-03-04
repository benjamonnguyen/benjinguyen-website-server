package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	//
	addr := flag.String("addr", ":3000", "http server address")
	flag.Parse()

	// router
	http.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/index.html")
	})
	http.HandleFunc("GET /{filepath...}", func(w http.ResponseWriter, r *http.Request) {
		path := "public/" + r.PathValue("filepath")
		if _, err := os.Stat(path); err != nil {
			fmt.Println("nonexistent resource requested at filepath:", path)
			http.Error(w, "", 404)
			return
		}

		http.ServeFile(w, r, path)
	})

	// serve
	s := &http.Server{
		Addr:        *addr,
		ReadTimeout: 10 * time.Second,
	}
	fmt.Println("starting http server at", s.Addr)
	log.Fatal(s.ListenAndServe())
}

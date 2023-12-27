package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	//
	addr := flag.String("addr", ":3000", "http server address")
	flag.Parse()

	//
	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, os.Interrupt)

	router := httprouter.New()
	// pages
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		http.ServeFile(w, r, "public/index.html")
	})
	router.GET("/resume", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		http.ServeFile(w, r, "public/resume.html")
	})
	// assets
	router.GET(
		"/css/:sheet",
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			http.ServeFile(w, r, "public/css/"+p.ByName("sheet"))
		},
	)
	router.GET(
		"/assets/:asset",
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			http.ServeFile(w, r, "public/assets/"+p.ByName("asset"))
		},
	)

	// serve
	srv := http.Server{Addr: *addr, Handler: router}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		srv.Shutdown(ctx)
	}()
	go srv.ListenAndServe()
	log.Println("started http server at", *addr)

	//
	<-interruptCh
}

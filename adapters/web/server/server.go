package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/feliperrpereira/go-hexagonal/adapters/web/handler"
	"github.com/feliperrpereira/go-hexagonal/application"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type Webserver struct {
	Service application.ProductServiceInterface
}

func MakeNewWebserver() *Webserver {
	return &Webserver{}
}

func (w Webserver) Server() {

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	log.Println("Server started on :9000")
}

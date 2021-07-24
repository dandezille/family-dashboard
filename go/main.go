package main

import (
	"log"
	"net/http"
	"time"

	"server/app"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	defer app.Close()

	router, err := app.SetupRouter()
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 15,
		Handler:      router,
	}

	log.Print("Server starting")
	log.Fatal(srv.ListenAndServe())
}

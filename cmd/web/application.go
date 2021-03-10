package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/petrostrak/booking-with-go/pkg/config"
	"github.com/petrostrak/booking-with-go/pkg/handlers"
	"github.com/petrostrak/booking-with-go/pkg/render"
)

const (
	portNumber = ":8000"
)

// StartApp func will initializa the url mapping
// and the server
func startApp() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	if err = serve.ListenAndServe(); err != nil {
		panic(err)
	}
}

package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/petrostrak/booking-with-go/pkg/config"
	"github.com/petrostrak/booking-with-go/pkg/render"
)

const (
	portNumber = ":8000"
)

// StartApp func will initializa the url mapping
// and the server
func StartApp() {

	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc

	urlMapping()

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	if err := http.ListenAndServe(portNumber, nil); err != nil {
		panic(err)
	}
}

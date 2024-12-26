package main

import (
	"log"
	"net/http"

	"github.com/ehsanmsb/Tati/pkg/config"
	"github.com/ehsanmsb/Tati/pkg/handlers"
	"github.com/ehsanmsb/Tati/pkg/render"
)

const ServerPortNumber = ":8080"

func main() {
	var app = config.AppConfig{}

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Println("cannot create template cache")
	}
	app.UseCache = false
	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/contact-us", handlers.Repo.ContactUs)
	http.HandleFunc("/aboute", handlers.Repo.Aboute)
	http.HandleFunc("/divide", handlers.Repo.Divide)
	err = http.ListenAndServe(ServerPortNumber, nil)
	if err != nil {
		log.Println(err)
	}
}

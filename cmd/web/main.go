package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/joe16723/booking/pkg/config"
	"github.com/joe16723/booking/pkg/handlers"
	"github.com/joe16723/booking/pkg/render"
	"log"
	"net/http"
	"time"
)

// global variable in main package

const PortNumber string = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	//change to true when in production
	app.Production = false

	// initialize session
	session = scs.New()
	session.Lifetime = 24 * time.Hour // expire in 1 day
	session.Cookie.Persist = true     // data still valid with persist = true of close browser
	session.Cookie.Secure = app.Production

	//set session to config
	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cant create template cache")
	}

	//load once for template cache
	app.TemplateCache = templateCache
	// dev mode
	app.UseCache = false // set to true if production mode

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	// using http handleFunc
	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)
	//fmt.Printf("Starting on port %s", PortNumber)
	//_ = http.ListenAndServe(PortNumber, nil)

	// using pat or chi lib from GitHub by routes concept
	fmt.Printf("Starting on port %s", PortNumber)
	serve := &http.Server{
		Addr:    PortNumber,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

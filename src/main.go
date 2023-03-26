package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Moataz-Hamed/bookings/pkg/config"
	"github.com/Moataz-Hamed/bookings/pkg/handlers"
	"github.com/Moataz-Hamed/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager
var app config.AppConfig

func main() {

	// change this to true when in production
	app.InProduction = false

	//session store data in cookies by default
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	// data stored in cookies remain on close of the web site
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	// not https
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can't create template cache", err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHanlder(repo)

	render.NewTemplate(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	// _ = http.ListenAndServe(":8080", nil)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}

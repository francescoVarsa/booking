package main

import (
	"booking/internal/config"
	"booking/internal/driver"
	"booking/internal/handlers"
	"booking/internal/helpers"
	"booking/internal/models"
	"booking/internal/render"
	"encoding/gob"
	"os"

	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main is the main application function
func main() {
	db, err := run()

	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()
	defer close(app.MailChan)

	fmt.Println("Starting mail listener...")
	listenForMail()

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	// _ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	// What am I going to put in the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

	//Change this to true when in prod
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database
	log.Println("Please insert your database password or leave it empty if your database has no password")
	var pwd string
	fmt.Scanln(&pwd)
	log.Println("Connection to database...")
	// Replace password with your password db
	db, err := driver.ConnectSQL(fmt.Sprintf("host=localhost port=5432 dbname=bookings user=francesco password=%s", pwd))
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}
	log.Println("Connected to database")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}

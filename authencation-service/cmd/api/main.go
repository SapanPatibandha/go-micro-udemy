package main

import (
	"database/sql"
	"log"
)

const webPoart = "8084"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("starting authentication service")

	// TODO connect to DB

	// set up config
	app := Config{}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort)
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

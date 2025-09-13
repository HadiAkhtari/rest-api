package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"rest-api-in-gin/internal/database"
	"rest-api-in-gin/internal/env"
)

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	connstr := "host=localhost port=5432 user=postgres password=1234 dbname=rest-api sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	models := database.NewModels(db)

	app := &application{
		port:      env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "some-secret-1213123"),
		models:    models,
	}

	if err := serve(app); err != nil {
		log.Fatal(err)
	}

}

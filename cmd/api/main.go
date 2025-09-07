package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	connstr := "host=localhost port=5432 user=postgres password=1234 dbname=rest-api sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}

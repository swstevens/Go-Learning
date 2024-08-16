package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type SquirrelObservation struct {
	SquirrelID      string
	PrimaryFurColor string
	Location        string
	Activities      string
}

func main() {
	// Connect to the database
	connStr := "host=localhost user=raspberry dbname=squirrels password=postgrespassword sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Perform a query
	rows, err := db.Query(`
        SELECT squirrel_id, primary_fur_color, location, activities
        FROM squirrel_data
        WHERE primary_fur_color = $1
        LIMIT 10
    `, "Gray")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Process the results
	var observations []SquirrelObservation
	for rows.Next() {
		var obs SquirrelObservation
		err := rows.Scan(&obs.SquirrelID, &obs.PrimaryFurColor, &obs.Location, &obs.Activities)
		if err != nil {
			log.Fatal(err)
		}
		observations = append(observations, obs)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Print the results
	for _, obs := range observations {
		fmt.Printf("Squirrel ID: %s, Color: %s, Location: %s, Activities: %s\n",
			obs.SquirrelID, obs.PrimaryFurColor, obs.Location, obs.Activities)
	}
}

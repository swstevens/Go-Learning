package main

import (
	"database/sql"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	// set up pointer to database (tentatively postgres)
	// this file won't be committed for safety
	var jsonFile, err = os.Open("secret.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	http.HandleFunc("/", handler)
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/counter", counter)
	http.HandleFunc("/lissajous", func(w http.ResponseWriter, r *http.Request) { lissajous(w, r) })
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	connStr := "host=localhost user=myuser dbname=mydb password=mypassword sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	// Perform a query
	rows, err := db.Query(`
			SELECT squirrel_id, primary_fur_color, location, activities
			FROM squirrel_data
			LIMIT 1
	`)
	if err != nil {
		log.Fatalf("Error querying database: %v", err)
	}
	defer rows.Close()

	// Process the results
	for rows.Next() {
		var id, color, location, activities string
		err := rows.Scan(&id, &color, &location, &activities)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
		fmt.Fprintf(w, "ID: %s, Color: %s, Location: %s, Activities: %s\n", id, color, location, activities)
	}

	if err = rows.Err(); err != nil {
		log.Fatalf("Error after scanning rows: %v", err)
	}

	//fmt.Println("Query executed successfully")
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajous(out io.Writer, r *http.Request) {

	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	var cycles = 5
	if r.URL.Query()["cycles"] != nil {
		cycles, _ = strconv.Atoi(r.URL.Query()["cycles"][0])
	}
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2.0*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

// connect to database
var connectionStr = "user=postgres dbname=postgres password=00000000 host=127.0.0.1 sslmode=disable"
var db, err = sql.Open("postgres", connectionStr)

func handler(w http.ResponseWriter, r *http.Request) {
	// set headers
	w.Header().Set("Content-Type", "application/json")

	// if the path = "database"
	if r.URL.Path[1:] == "database" {

		// get id from url query
		q := r.URL.Query()
		id := q.Get("id")
		limit := q.Get("limit")

		// Row struct
		type Row struct {
			ID        rune
			Username  string
			Text      string
			Created   string
			ProjectID string
		}

		// query the database with the query id
		query := fmt.Sprintf("SELECT * FROM comments WHERE id > %s LIMIT %s", id, limit)
		data, err := db.Query(query)
		if err != nil {
			fmt.Println(err)
		}

		// loop through and print results
		slice := make([]Row, 0)
		for data.Next() {
			var row Row

			err := data.Scan(&row.ID, &row.Username, &row.Text, &row.Created, &row.ProjectID)
			if err != nil {
				fmt.Println(err)
			}
			slice = append(slice, row)
		}

		// encode slice of rows to json and to []byte
		comments, err := json.Marshal(slice)
		if err != nil {
			fmt.Println(err)
		}
		// write to response body
		w.Write(comments)
	}
}

// Main function
func main() {
	fmt.Println("running...")
	err := db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	// request handler
	http.HandleFunc("/database", handler)
	http.Handle("/", http.FileServer(http.Dir("dist")))

	// custom server
	server := &http.Server{
		Addr:         ":3000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// listen for requests
	log.Fatal(server.ListenAndServe())
}

// ** ** ** ** ** ** ** ** ** ** ** ** ** ** **
// Deployed Database stats :)
//
// test: SELECT * FROM comments WHERE id > 5000000 LIMIT 100
// average responce before: "~13sec"
// average response after:  "~0.0007sec"
// ** ** ** ** ** ** ** ** ** ** ** ** ** ** **

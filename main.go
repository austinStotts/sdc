package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request /")
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func database(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request /database")
	connectionStr := "user=postgres dbname=postgres password=00000000 host=127.0.0.1 sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		fmt.Println(err)
	}

	if db != nil {
		fmt.Println("succesully connected to database")
	}

	// query := "SELECT * FROM messages WHERE id = 5000000"
	// db.Query(query)
}

func main() {
	fmt.Println("running...")

	http.HandleFunc("/", handler) // handle home requests
	http.HandleFunc("/database", database) // handle database requests

	log.Fatal(http.ListenAndServe(":3000", nil))
}

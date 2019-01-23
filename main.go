package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"log"
	_ "github.com/lib/pq"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request /")
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func database(w http.ResponseWriter, r *http.Request) {

	type Row struct {
		Id int 
		Username string 
		Text string 
		Created string 
		Project_id string 
	}

	fmt.Println("request /database")
	connectionStr := "user=postgres dbname=postgres password=00000000 host=127.0.0.1 sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	fmt.Println("succesully connected to database")

	query := "SELECT username, text FROM messages"
	data, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}

	for data.Next() {
		var row Row

		err := data.Scan(&row.Username, &row.Text)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("row ->" + row.Username + ": " + row.Text)
	}

}

func main() {
	fmt.Println("running...")

	http.HandleFunc("/", handler) // handle home requests
	http.HandleFunc("/database", database) // handle database requests

	log.Fatal(http.ListenAndServe(":3000", nil))
}

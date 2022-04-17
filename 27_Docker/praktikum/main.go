package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getDBConnection() (*sql.DB, error) {
	return sql.Open("mysql", "sample_user:secret@tcp(db:3306)/sample_db")
}

func getUsers() []*User {
	db, err := getDBConnection()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	users := []*User{}
	for results.Next() {
		u := User{}

		err = results.Scan(&u.ID, &u.Name)
		if err != nil {
			panic(err)
		}

		users = append(users, &u)
	}

	return users
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to homepage")
	fmt.Println("Hit homepage endpoint")
}

func userPage(w http.ResponseWriter, r *http.Request) {
	users := getUsers()

	fmt.Println("Hit userpage endpoint")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", userPage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

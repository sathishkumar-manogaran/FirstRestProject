package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type PeopleShowcase struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	From      string `json:"from"`
	Interests string `json:"interests"`
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	db, err := connectDB()

	saveIntoDB(err, db)

	findAllFromDB(err, db)

	findById(err, db)

}

func findById(err error, db *sql.DB) {
	var peopleShowcase PeopleShowcase
	err = db.QueryRow("SELECT * FROM people_showcase WHERE ID = ?", 1).Scan(&peopleShowcase.Id, &peopleShowcase.Name, &peopleShowcase.From, &peopleShowcase.Interests)
	if err != nil {
		panic(err.Error())
	}

	log.Println(peopleShowcase)

}

func findAllFromDB(err error, db *sql.DB) {
	selectQuery, err := db.Query("SELECT * FROM people_showcase")
	if err != nil {
		panic(err.Error())
	}

	for selectQuery.Next() {
		var peopleShowcase PeopleShowcase
		err := selectQuery.Scan(&peopleShowcase.Id, &peopleShowcase.Name, &peopleShowcase.From, &peopleShowcase.Interests)
		if err != nil {
			panic(err.Error())
		}

		//log.Printf(peopleShowcase.Name)
		log.Println(peopleShowcase)
	}
}

func saveIntoDB(err error, db *sql.DB) {
	insetQuery, err := db.Query("INSERT INTO people_showcase values (1,'Test', 'Chennai', 'Developer')")

	if err != nil {
		panic(err.Error())
	}

	defer insetQuery.Close()
}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/people_showcase")

	if err != nil {
		panic(err.Error())
	}

	//defer db.Close()
	return db, err
}

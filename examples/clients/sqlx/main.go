package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Structを用意
type User struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

const (
	DBMS           = "mysql"
	dataSourceName = "user:password@tcp(localhost:3306)/examples?charset=utf8mb4"
)

func main() {

	//DBに接続する
	db, err := sqlx.Open(DBMS, dataSourceName)

	if err != nil {
		defer db.Close()
		log.Fatal(err)
	}

	// SQL実行 (INSERT)
	user := User{Name: "John Doe", Email: "john@example.com"}
	result, err := db.NamedExec("INSERT INTO users(name, email) VALUES(:name, :email)", user)
	if err != nil {
		log.Fatal(err)
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Last inserted ID:", lastInsertID)

	// SQL実行 (SELECT)
	rows, err := db.Queryx("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {

		var user User

		err := rows.StructScan(&user)

		if err != nil {
			log.Fatal(err)
		}

		results = append(results, user)
	}

	fmt.Println(results)

}

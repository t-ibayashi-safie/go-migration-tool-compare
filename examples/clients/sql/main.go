package main

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Structを用意
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

const (
	DBMS           = "mysql"
	dataSourceName = "user:password@tcp(localhost:3306)/examples?charset=utf8mb4"
)

func main() {

	//DBに接続する
	db, err := sql.Open(DBMS, dataSourceName)

	if err != nil {
		defer db.Close()
		log.Fatal(err)
	}
	// SQL実行 (INSERT)
	stmt, err := db.Prepare("INSERT INTO users(name, email) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec("John Doe", "john@example.com")
	if err != nil {
		log.Fatal(err)
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Last inserted ID:", lastInsertID)

	// SQL実行 (SELECT)
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	// 結果を取得
	results := make([]User, 0)
	for rows.Next() {

		var user User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
		)

		if err != nil {
			log.Fatal(err)
		}

		results = append(results, user)
	}

	fmt.Println(results)

}

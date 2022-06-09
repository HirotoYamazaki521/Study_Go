package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//database + sqlite3
//テーブル作成

var DbConnection *sql.DB

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	/*
			cmd := `CREATE TABLE IF NOT EXISTS persons(
		                name STRING,
						age  INT)`
			_, err := DbConnection.Exec(cmd)
	*/

	cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	_, err := DbConnection.Exec(cmd, "Nancy", 20)

	//データの追加
	if err != nil {
		log.Fatalln(err)
	}

}

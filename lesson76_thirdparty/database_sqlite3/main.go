package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//database + sqlite3
//テーブル作成

var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	/*
			cmd := `CREATE TABLE IF NOT EXISTS persons(
		                name STRING,
						age  INT)`
			_, err := DbConnection.Exec(cmd)
	*/

	//データの追加
	/*
		cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
		_, err := DbConnection.Exec(cmd, "Nancy", 20)

		cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
		_, err := DbConnection.Exec(cmd, "hanako", 19)
	*/

	/*
		//データの更新
		cmd := "UPDATE persons SET age = ? WHERE name = ?"
		_, err := DbConnection.Exec(cmd, 25, "Nancy")
	*/

	//特定のデータを取得
	/*
		cmd := "SELECT * FROM persons where age = ?"
		//QueryRow 1レコード取得
		row := DbConnection.QueryRow(cmd, 25)
		var p Person
		err := row.Scan(&p.Name, &p.Age)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Println("No row")
			} else {
				log.Println(err)
			}
			log.Fatalln(err)
		}
		fmt.Println(p.Name, p.Age)
	*/

	//複数データの取得
	/*
		cmd := "SELECT * FROM persons"
		//Queryは条件に合うものをすべて取得
		rows, _ := DbConnection.Query(cmd)
		defer rows.Close()
		var pp []Person
		for rows.Next() {
			var p Person
			err := rows.Scan(&p.Name, &p.Age)
			if err != nil {
				log.Println(err)
			}
			pp = append(pp, p)
		}

		for _, p := range pp {
			fmt.Println(p.Name, p.Age)
		}
	*/

	//データの削除
	cmd := "DELETE FROM persons WHERE name = ?"
	_, err := DbConnection.Exec(cmd, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}

}

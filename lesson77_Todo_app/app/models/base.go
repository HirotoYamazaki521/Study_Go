package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"lesson77_Todo_app/config"
	"log"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

//テーブル作成のためのコード

//usersテーブルの作成
var Db *sql.DB

var err error

const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)

//テーブルをmain関数の前に作成
func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	//コマンドの作成
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	//コマンドの実行
	Db.Exec(cmdU)

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			content TEXT,
			user_id INTEGER,
			created_at DATETIME)`, tableNameTodo)

	//コマンドの実行
	Db.Exec(cmdT)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		user_id INTEGER,
		created_at DATETIME)`, tableNameSession)

	//コマンドの実行
	Db.Exec(cmdS)
}

func createUUID() (uuidobj uuid.UUID) { //uuidパッケージのUUID型を使っている
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

//パスワードの保存はハッシュ値にする必要がある
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext))) //sha1のアルゴリズムでハッシュ値に変換
	return cryptext
}

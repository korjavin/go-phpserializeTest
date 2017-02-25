package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/korjavin/phpserialize2Json"
)

var db *sql.DB
var stmt *sql.Stmt
var stmt1 *sql.Stmt

const dburl = "root:@unix(/run/mysqld/mysqld.sock)/kino"

func init() {
	var err error
	db, err = sql.Open("mysql", dburl)
	if err != nil {
		log.Fatalf("mysql: %s", err)
	}
	stmt, err = db.Prepare("update kp_people_films set films_json=? where id=?")
	if err != nil {
		log.Fatal(err)
	}
	stmt1, err = db.Prepare("select id, films_ser from kp_people_films where films_json is null limit ?")
	if err != nil {
		log.Fatal(err)
	}
}

func update(i int) {
	rows, err := stmt1.Query(i)
	defer rows.Close()
	var (
		id  int
		ser string
	)
	if err != nil {
		log.Fatalf("mysql: %s", err)
	}
	for rows.Next() {
		if err := rows.Scan(&id, &ser); err != nil {
			log.Fatal(err)
		}
		json1, err := json.DecodeToJSON(ser)
		if err != nil {
			log.Fatal(err)
		}
		_, err = stmt.Exec(json1, id)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%d \n", id)
	}
}
func check(c chan<- bool) {
}
func main() {
	for {
		update(90)
		time.Sleep(time.Second * 1)
	}
}

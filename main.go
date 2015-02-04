package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type cafemenu struct {
	id          int
	category    string
	name        string
	chinesename string
	price       float32
	updeat      time.Time
}

func main() {
	log.Println("Hello world!")
	db, err := sql.Open("postgres", "postgres://XJ:xjg123@/xjdb?sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	var cafelist string
	rows, err := db.Query("Select chinesename from cafe")
	for rows.Next() {
		err := rows.Scan(&cafelist)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("cafe list is %s", cafelist)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	rows, err = db.Query("select * from cafe")
	for rows.Next() {
		menuitem := cafemenu{}
		err := rows.Scan(&menuitem.id, &menuitem.category, &menuitem.name, &menuitem.chinesename, &menuitem.price, &menuitem.updeat)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("the menu is %s %s %s %d %s", menuitem.category, menuitem.name, menuitem.chinesename, menuitem.price, menuitem.updeat)
	}

	now := time.Now()
	res, err := db.Exec("insert into cafe (category,name, chinesename, price,last_update) values($1,$2,$3,$4,$5)", "tea", "Milktea", "奶茶", 4.25, now)
	if err != nil {
		log.Fatal(err)
	}
	affected, _ := res.RowsAffected()
	log.Printf("The number of rows affected is %d", affected)

}

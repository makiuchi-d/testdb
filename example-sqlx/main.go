package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Player struct {
	ID   uint   `db:"id"`
	Name string `db:"name"`
	Age  uint   `db:"age"`
}

func GetUnder21(db *sqlx.DB) []Player {
	var players []Player
	err := db.Select(&players, "SELECT * FROM players WHERE age <= 21")
	if err != nil {
		panic(err)
	}
	return players
}

func IncrementAge(db *sqlx.DB) {
	_, err := db.Exec("UPDATE players SET age = age + 1")
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sqlx.Open("mysql", "testdb:testdb@tcp(localhost:3306)/testdb?parseTime=true")
	if err != nil {
		panic(err)
	}

	u21 := GetUnder21(db)
	fmt.Println("before:\n", u21)

	IncrementAge(db)

	u21 = GetUnder21(db)
	fmt.Println("after:\n", u21)
}

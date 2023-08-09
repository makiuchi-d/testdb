package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Player struct {
	ID   uint `gorm:"primarykey"`
	Name string
	Age  uint
}

func GetUnder21(db *gorm.DB) []Player {
	var players []Player
	db.Where("age <= 21").Find(&players)
	return players
}

func IncrementAge(db *gorm.DB) {
	db.Table("players").Where("1=1").UpdateColumn("age", gorm.Expr("age + ?", 1))
}

func main() {
	db, err := gorm.Open(mysql.Open("testdb:testdb@tcp(localhost:3306)/testdb?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	u21 := GetUnder21(db)
	fmt.Println("before:\n", u21)

	IncrementAge(db)

	u21 = GetUnder21(db)
	fmt.Println("after:\n", u21)
}

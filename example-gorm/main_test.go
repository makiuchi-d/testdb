package main_test

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/makiuchi-d/testdb"
	main "github.com/makiuchi-d/testdb/example-gorm"
)

var testdata = map[string]uint{
	"Alice": 20,
	"Bob":   23,
	"Carol": 21,
}

func newTestDb(dbname string) *gorm.DB {
	db := testdb.Must1(
		gorm.Open(mysql.New(mysql.Config{Conn: testdb.New(dbname)}), nil))

	db.AutoMigrate(&main.Player{})

	for name, age := range testdata {
		db.Exec("INSERT INTO players (name, age) VALUES (?, ?)", name, age)
	}

	return db
}

func TestGetUnder21(t *testing.T) {
	db := newTestDb("test_under_21")

	u21 := main.GetUnder21(db)

	if len(u21) != 2 {
		t.Fatalf("must be 2 players: %v", u21)
	}

	for _, p := range u21 {
		if p.Name != "Alice" && p.Name != "Carol" {
			t.Fatalf("unexpected player: %v", p)
		}
	}
}

func TestIncrementAge(t *testing.T) {
	db := newTestDb("test_increment_age")

	main.IncrementAge(db)

	var all []main.Player
	db.Find(&all)

	if len(all) != len(testdata) {
		t.Fatalf("length must be %v: %v", len(testdata), all)
	}

	for _, p := range all {
		age, ok := testdata[p.Name]
		if !ok {
			t.Fatalf("unexpected player: %v", p)
		}
		if p.Age != age+1 {
			t.Fatalf("age of %q = %v, wants %v", p.Name, p.Age, age+1)
		}
	}
}

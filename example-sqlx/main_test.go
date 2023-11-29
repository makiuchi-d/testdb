package main_test

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"

	"github.com/makiuchi-d/testdb"
	main "github.com/makiuchi-d/testdb/example-sqlx"
)

var testdata = map[string]uint{
	"Alice": 20,
	"Bob":   23,
	"Carol": 21,
}

func newTestDB(dbname string) *sqlx.DB {
	db := testdb.New(dbname)

	testdb.Must1(db.Exec(
		string(testdb.Must1(os.ReadFile("../testdata/00_schema.sql")))))

	for name, age := range testdata {
		testdb.Must1(db.Exec("INSERT INTO players (name, age) VALUES (?, ?)", name, age))
	}

	return sqlx.NewDb(db, "mysql")
}

func TestGetUnder21(t *testing.T) {
	db := newTestDB("test_get_under21")

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
	db := newTestDB("test_increment_age")

	main.IncrementAge(db)

	var all []main.Player
	testdb.Must(db.Select(&all, "SELECT * FROM players"))

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

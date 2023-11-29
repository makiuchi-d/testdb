package main_test

import (
	"context"
	"math"
	"os"
	"testing"

	"github.com/makiuchi-d/testdb"
	main "github.com/makiuchi-d/testdb/example-sqlc"
	"github.com/makiuchi-d/testdb/example-sqlc/sqlc"
)

var testdata = map[string]uint32{
	"Alice": 20,
	"Bob":   23,
	"Carol": 21,
}

func newTestQuery(dbname string) *sqlc.Queries {
	db := testdb.New(dbname)

	testdb.Must1(db.Exec(
		string(testdb.Must1(os.ReadFile("../testdata/00_schema.sql")))))

	for name, age := range testdata {
		testdb.Must1(db.Exec("INSERT INTO players (name, age) VALUES (?, ?)", name, age))
	}

	return sqlc.New(db)
}

func TestGetUnder21(t *testing.T) {
	ctx := context.Background()

	query := newTestQuery("test_get_under21")

	u21 := main.GetUnder21(ctx, query)

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
	ctx := context.Background()

	query := newTestQuery("test_increment_age")

	main.IncrementAge(ctx, query)

	all := testdb.Must1(query.GetUnderAge(ctx, math.MaxUint32))

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

package main_test

import (
	"context"
	"testing"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	"github.com/makiuchi-d/testdb"
	main "github.com/makiuchi-d/testdb/example-ent"
	"github.com/makiuchi-d/testdb/example-ent/ent"
	"github.com/makiuchi-d/testdb/example-ent/ent/enttest"
)

var testdata = map[string]uint{
	"Alice": 20,
	"Bob":   23,
	"Carol": 21,
}

func newTestClient(t *testing.T, dbname string) *ent.Client {
	db := testdb.New(dbname)

	drv := entsql.OpenDB(dialect.MySQL, db)
	client := enttest.NewClient(t, enttest.WithOptions(ent.Driver(drv)))

	for name, age := range testdata {
		testdb.Must1(db.Exec("INSERT INTO players (name, age) VALUES (?, ?)", name, age))
	}

	return client
}

func TestGetUnder21(t *testing.T) {
	ctx := context.Background()

	client := newTestClient(t, "test_get_under21")

	u21 := main.GetUnder21(ctx, client)

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

	client := newTestClient(t, "test_increment_age")

	main.IncrementAge(ctx, client)

	all := client.Player.Query().AllX(ctx)

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

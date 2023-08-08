package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/makiuchi-d/testdb/example-sqlc/sqlc"
)

func GetUnder21(ctx context.Context, query *sqlc.Queries) []sqlc.Player {
	u21, err := query.GetUnderAge(ctx, 21)
	if err != nil {
		panic(err)
	}
	return u21
}

func IncrementAge(ctx context.Context, query *sqlc.Queries) {
	err := query.AddAge(ctx, 1)
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()

	db, err := sql.Open("mysql", "testdb:testdb@tcp(localhost:3306)/testdb?parseTime=true")
	if err != nil {
		panic(err)
	}
	query := sqlc.New(db)

	u21 := GetUnder21(ctx, query)
	fmt.Println("before:\n", u21)

	IncrementAge(ctx, query)

	u21 = GetUnder21(ctx, query)
	fmt.Println("after:\n", u21)
}

package main

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/makiuchi-d/testdb/example-ent/ent"
	"github.com/makiuchi-d/testdb/example-ent/ent/player"
)

func GetUnder21(ctx context.Context, client *ent.Client) []*ent.Player {
	return client.Player.Query().Where(player.AgeLTE(21)).AllX(ctx)
}

func IncrementAge(ctx context.Context, client *ent.Client) {
	client.Player.Update().AddAge(1).ExecX(ctx)
}

func main() {
	client, err := ent.Open("mysql", "testdb:testdb@tcp(localhost:3306)/testdb?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.TODO()

	u21 := GetUnder21(ctx, client)
	fmt.Println("before:\n", u21)

	IncrementAge(ctx, client)

	u21 = GetUnder21(ctx, client)
	fmt.Println("after:\n", u21)
}

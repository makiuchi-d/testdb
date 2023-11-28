package testdb_test

import (
	"fmt"

	"github.com/makiuchi-d/testdb"
)

func Example() {
	db := testdb.New("testdb1", "testdb2")

	testdb.Must1(
		db.Exec("CREATE TABLE user (id INTEGER AUTO_INCREMENT, name VARCHAR(15), PRIMARY KEY (id))"))
	testdb.Must1(
		db.Exec("INSERT INTO user (id, name) VALUES (1, 'Alice'), (2, 'Bob')"))
	testdb.Must1(
		db.Exec("CREATE TABLE testdb2.score" +
			" (user_id INTEGER, stage_id INTEGER, score INTEGER, PRIMARY KEY (user_id, stage_id))"))
	testdb.Must1(
		db.Exec("INSERT INTO testdb2.score (user_id, stage_id, score) VALUES" +
			" (1, 1, 30), (1, 2, 40), (1, 4, 50), (2, 1, 10), (2, 3, 40)"))

	var (
		id    int
		name  string
		stage int
		score int
	)
	rows := testdb.Must1(
		db.Query("SELECT id, name, stage_id, score" +
			" FROM user JOIN testdb2.score ON user.id = score.user_id" +
			" WHERE stage_id = 1 ORDER BY user.id"))
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&id, &name, &stage, &score)
		fmt.Println(id, name, stage, score)
	}
	// output:
	// 1 Alice 1 30
	// 2 Bob 1 10
}

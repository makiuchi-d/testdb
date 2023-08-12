// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package sqlc

import (
	"context"
)

const addAge = `-- name: AddAge :exec
UPDATE players SET age = age + ?
`

func (q *Queries) AddAge(ctx context.Context, age uint32) error {
	_, err := q.db.ExecContext(ctx, addAge, age)
	return err
}

const getUnderAge = `-- name: GetUnderAge :many
SELECT id, name, age FROM players WHERE age <= ?
`

func (q *Queries) GetUnderAge(ctx context.Context, age uint32) ([]Player, error) {
	rows, err := q.db.QueryContext(ctx, getUnderAge, age)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Player
	for rows.Next() {
		var i Player
		if err := rows.Scan(&i.ID, &i.Name, &i.Age); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

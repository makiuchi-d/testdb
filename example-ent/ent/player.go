// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/makiuchi-d/testdb/example-ent/ent/player"
)

// Player is the model entity for the Player schema.
type Player struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Age holds the value of the "age" field.
	Age          uint `json:"age,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Player) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case player.FieldID, player.FieldAge:
			values[i] = new(sql.NullInt64)
		case player.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Player fields.
func (pl *Player) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case player.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pl.ID = int(value.Int64)
		case player.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pl.Name = value.String
			}
		case player.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				pl.Age = uint(value.Int64)
			}
		default:
			pl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Player.
// This includes values selected through modifiers, order, etc.
func (pl *Player) Value(name string) (ent.Value, error) {
	return pl.selectValues.Get(name)
}

// Update returns a builder for updating this Player.
// Note that you need to call Player.Unwrap() before calling this method if this Player
// was returned from a transaction, and the transaction was committed or rolled back.
func (pl *Player) Update() *PlayerUpdateOne {
	return NewPlayerClient(pl.config).UpdateOne(pl)
}

// Unwrap unwraps the Player entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pl *Player) Unwrap() *Player {
	_tx, ok := pl.config.driver.(*txDriver)
	if !ok {
		panic("ent: Player is not a transactional entity")
	}
	pl.config.driver = _tx.drv
	return pl
}

// String implements the fmt.Stringer.
func (pl *Player) String() string {
	var builder strings.Builder
	builder.WriteString("Player(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pl.ID))
	builder.WriteString("name=")
	builder.WriteString(pl.Name)
	builder.WriteString(", ")
	builder.WriteString("age=")
	builder.WriteString(fmt.Sprintf("%v", pl.Age))
	builder.WriteByte(')')
	return builder.String()
}

// Players is a parsable slice of Player.
type Players []*Player
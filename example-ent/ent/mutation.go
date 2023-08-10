// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/makiuchi-d/testdb/example-ent/ent/player"
	"github.com/makiuchi-d/testdb/example-ent/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypePlayer = "Player"
)

// PlayerMutation represents an operation that mutates the Player nodes in the graph.
type PlayerMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	age           *uint
	addage        *int
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Player, error)
	predicates    []predicate.Player
}

var _ ent.Mutation = (*PlayerMutation)(nil)

// playerOption allows management of the mutation configuration using functional options.
type playerOption func(*PlayerMutation)

// newPlayerMutation creates new mutation for the Player entity.
func newPlayerMutation(c config, op Op, opts ...playerOption) *PlayerMutation {
	m := &PlayerMutation{
		config:        c,
		op:            op,
		typ:           TypePlayer,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withPlayerID sets the ID field of the mutation.
func withPlayerID(id int) playerOption {
	return func(m *PlayerMutation) {
		var (
			err   error
			once  sync.Once
			value *Player
		)
		m.oldValue = func(ctx context.Context) (*Player, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Player.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withPlayer sets the old Player of the mutation.
func withPlayer(node *Player) playerOption {
	return func(m *PlayerMutation) {
		m.oldValue = func(context.Context) (*Player, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PlayerMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PlayerMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *PlayerMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *PlayerMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Player.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *PlayerMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *PlayerMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Player entity.
// If the Player object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PlayerMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *PlayerMutation) ResetName() {
	m.name = nil
}

// SetAge sets the "age" field.
func (m *PlayerMutation) SetAge(u uint) {
	m.age = &u
	m.addage = nil
}

// Age returns the value of the "age" field in the mutation.
func (m *PlayerMutation) Age() (r uint, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// OldAge returns the old "age" field's value of the Player entity.
// If the Player object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PlayerMutation) OldAge(ctx context.Context) (v uint, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAge is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAge requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAge: %w", err)
	}
	return oldValue.Age, nil
}

// AddAge adds u to the "age" field.
func (m *PlayerMutation) AddAge(u int) {
	if m.addage != nil {
		*m.addage += u
	} else {
		m.addage = &u
	}
}

// AddedAge returns the value that was added to the "age" field in this mutation.
func (m *PlayerMutation) AddedAge() (r int, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ResetAge resets all changes to the "age" field.
func (m *PlayerMutation) ResetAge() {
	m.age = nil
	m.addage = nil
}

// Where appends a list predicates to the PlayerMutation builder.
func (m *PlayerMutation) Where(ps ...predicate.Player) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the PlayerMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *PlayerMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Player, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *PlayerMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *PlayerMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Player).
func (m *PlayerMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *PlayerMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.name != nil {
		fields = append(fields, player.FieldName)
	}
	if m.age != nil {
		fields = append(fields, player.FieldAge)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *PlayerMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case player.FieldName:
		return m.Name()
	case player.FieldAge:
		return m.Age()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *PlayerMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case player.FieldName:
		return m.OldName(ctx)
	case player.FieldAge:
		return m.OldAge(ctx)
	}
	return nil, fmt.Errorf("unknown Player field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PlayerMutation) SetField(name string, value ent.Value) error {
	switch name {
	case player.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case player.FieldAge:
		v, ok := value.(uint)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	}
	return fmt.Errorf("unknown Player field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *PlayerMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, player.FieldAge)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *PlayerMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case player.FieldAge:
		return m.AddedAge()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PlayerMutation) AddField(name string, value ent.Value) error {
	switch name {
	case player.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	}
	return fmt.Errorf("unknown Player numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *PlayerMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *PlayerMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *PlayerMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Player nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *PlayerMutation) ResetField(name string) error {
	switch name {
	case player.FieldName:
		m.ResetName()
		return nil
	case player.FieldAge:
		m.ResetAge()
		return nil
	}
	return fmt.Errorf("unknown Player field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *PlayerMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *PlayerMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *PlayerMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *PlayerMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *PlayerMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *PlayerMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *PlayerMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Player unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *PlayerMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Player edge %s", name)
}

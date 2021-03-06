// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sign-in/app/record/service/internal/data/ent/predicate"
	"sign-in/app/record/service/internal/data/ent/record"
	"sync"
	"time"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeRecord = "Record"
)

// RecordMutation represents an operation that mutates the Record nodes in the graph.
type RecordMutation struct {
	config
	op               Op
	typ              string
	id               *int
	user_id          *int64
	adduser_id       *int64
	sign_in_index    *int
	addsign_in_index *int
	reward           *float64
	addreward        *float64
	sign_in_day      *string
	created_at       *time.Time
	updated_at       *time.Time
	clearedFields    map[string]struct{}
	done             bool
	oldValue         func(context.Context) (*Record, error)
	predicates       []predicate.Record
}

var _ ent.Mutation = (*RecordMutation)(nil)

// recordOption allows management of the mutation configuration using functional options.
type recordOption func(*RecordMutation)

// newRecordMutation creates new mutation for the Record entity.
func newRecordMutation(c config, op Op, opts ...recordOption) *RecordMutation {
	m := &RecordMutation{
		config:        c,
		op:            op,
		typ:           TypeRecord,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withRecordID sets the ID field of the mutation.
func withRecordID(id int) recordOption {
	return func(m *RecordMutation) {
		var (
			err   error
			once  sync.Once
			value *Record
		)
		m.oldValue = func(ctx context.Context) (*Record, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Record.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withRecord sets the old Record of the mutation.
func withRecord(node *Record) recordOption {
	return func(m *RecordMutation) {
		m.oldValue = func(context.Context) (*Record, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m RecordMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m RecordMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *RecordMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *RecordMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Record.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUserID sets the "user_id" field.
func (m *RecordMutation) SetUserID(i int64) {
	m.user_id = &i
	m.adduser_id = nil
}

// UserID returns the value of the "user_id" field in the mutation.
func (m *RecordMutation) UserID() (r int64, exists bool) {
	v := m.user_id
	if v == nil {
		return
	}
	return *v, true
}

// OldUserID returns the old "user_id" field's value of the Record entity.
// If the Record object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RecordMutation) OldUserID(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserID: %w", err)
	}
	return oldValue.UserID, nil
}

// AddUserID adds i to the "user_id" field.
func (m *RecordMutation) AddUserID(i int64) {
	if m.adduser_id != nil {
		*m.adduser_id += i
	} else {
		m.adduser_id = &i
	}
}

// AddedUserID returns the value that was added to the "user_id" field in this mutation.
func (m *RecordMutation) AddedUserID() (r int64, exists bool) {
	v := m.adduser_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetUserID resets all changes to the "user_id" field.
func (m *RecordMutation) ResetUserID() {
	m.user_id = nil
	m.adduser_id = nil
}

// SetSignInIndex sets the "sign_in_index" field.
func (m *RecordMutation) SetSignInIndex(i int) {
	m.sign_in_index = &i
	m.addsign_in_index = nil
}

// SignInIndex returns the value of the "sign_in_index" field in the mutation.
func (m *RecordMutation) SignInIndex() (r int, exists bool) {
	v := m.sign_in_index
	if v == nil {
		return
	}
	return *v, true
}

// OldSignInIndex returns the old "sign_in_index" field's value of the Record entity.
// If the Record object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RecordMutation) OldSignInIndex(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSignInIndex is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSignInIndex requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSignInIndex: %w", err)
	}
	return oldValue.SignInIndex, nil
}

// AddSignInIndex adds i to the "sign_in_index" field.
func (m *RecordMutation) AddSignInIndex(i int) {
	if m.addsign_in_index != nil {
		*m.addsign_in_index += i
	} else {
		m.addsign_in_index = &i
	}
}

// AddedSignInIndex returns the value that was added to the "sign_in_index" field in this mutation.
func (m *RecordMutation) AddedSignInIndex() (r int, exists bool) {
	v := m.addsign_in_index
	if v == nil {
		return
	}
	return *v, true
}

// ResetSignInIndex resets all changes to the "sign_in_index" field.
func (m *RecordMutation) ResetSignInIndex() {
	m.sign_in_index = nil
	m.addsign_in_index = nil
}

// SetReward sets the "reward" field.
func (m *RecordMutation) SetReward(f float64) {
	m.reward = &f
	m.addreward = nil
}

// Reward returns the value of the "reward" field in the mutation.
func (m *RecordMutation) Reward() (r float64, exists bool) {
	v := m.reward
	if v == nil {
		return
	}
	return *v, true
}

// OldReward returns the old "reward" field's value of the Record entity.
// If the Record object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RecordMutation) OldReward(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldReward is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldReward requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldReward: %w", err)
	}
	return oldValue.Reward, nil
}

// AddReward adds f to the "reward" field.
func (m *RecordMutation) AddReward(f float64) {
	if m.addreward != nil {
		*m.addreward += f
	} else {
		m.addreward = &f
	}
}

// AddedReward returns the value that was added to the "reward" field in this mutation.
func (m *RecordMutation) AddedReward() (r float64, exists bool) {
	v := m.addreward
	if v == nil {
		return
	}
	return *v, true
}

// ResetReward resets all changes to the "reward" field.
func (m *RecordMutation) ResetReward() {
	m.reward = nil
	m.addreward = nil
}

// SetSignInDay sets the "sign_in_day" field.
func (m *RecordMutation) SetSignInDay(s string) {
	m.sign_in_day = &s
}

// SignInDay returns the value of the "sign_in_day" field in the mutation.
func (m *RecordMutation) SignInDay() (r string, exists bool) {
	v := m.sign_in_day
	if v == nil {
		return
	}
	return *v, true
}

// OldSignInDay returns the old "sign_in_day" field's value of the Record entity.
// If the Record object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RecordMutation) OldSignInDay(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSignInDay is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSignInDay requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSignInDay: %w", err)
	}
	return oldValue.SignInDay, nil
}

// ResetSignInDay resets all changes to the "sign_in_day" field.
func (m *RecordMutation) ResetSignInDay() {
	m.sign_in_day = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *RecordMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *RecordMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Record entity.
// If the Record object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RecordMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *RecordMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *RecordMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *RecordMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Record entity.
// If the Record object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RecordMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *RecordMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the RecordMutation builder.
func (m *RecordMutation) Where(ps ...predicate.Record) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *RecordMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Record).
func (m *RecordMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *RecordMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.user_id != nil {
		fields = append(fields, record.FieldUserID)
	}
	if m.sign_in_index != nil {
		fields = append(fields, record.FieldSignInIndex)
	}
	if m.reward != nil {
		fields = append(fields, record.FieldReward)
	}
	if m.sign_in_day != nil {
		fields = append(fields, record.FieldSignInDay)
	}
	if m.created_at != nil {
		fields = append(fields, record.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, record.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *RecordMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case record.FieldUserID:
		return m.UserID()
	case record.FieldSignInIndex:
		return m.SignInIndex()
	case record.FieldReward:
		return m.Reward()
	case record.FieldSignInDay:
		return m.SignInDay()
	case record.FieldCreatedAt:
		return m.CreatedAt()
	case record.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *RecordMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case record.FieldUserID:
		return m.OldUserID(ctx)
	case record.FieldSignInIndex:
		return m.OldSignInIndex(ctx)
	case record.FieldReward:
		return m.OldReward(ctx)
	case record.FieldSignInDay:
		return m.OldSignInDay(ctx)
	case record.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case record.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Record field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RecordMutation) SetField(name string, value ent.Value) error {
	switch name {
	case record.FieldUserID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserID(v)
		return nil
	case record.FieldSignInIndex:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSignInIndex(v)
		return nil
	case record.FieldReward:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetReward(v)
		return nil
	case record.FieldSignInDay:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSignInDay(v)
		return nil
	case record.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case record.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Record field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *RecordMutation) AddedFields() []string {
	var fields []string
	if m.adduser_id != nil {
		fields = append(fields, record.FieldUserID)
	}
	if m.addsign_in_index != nil {
		fields = append(fields, record.FieldSignInIndex)
	}
	if m.addreward != nil {
		fields = append(fields, record.FieldReward)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *RecordMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case record.FieldUserID:
		return m.AddedUserID()
	case record.FieldSignInIndex:
		return m.AddedSignInIndex()
	case record.FieldReward:
		return m.AddedReward()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RecordMutation) AddField(name string, value ent.Value) error {
	switch name {
	case record.FieldUserID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddUserID(v)
		return nil
	case record.FieldSignInIndex:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddSignInIndex(v)
		return nil
	case record.FieldReward:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddReward(v)
		return nil
	}
	return fmt.Errorf("unknown Record numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *RecordMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *RecordMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *RecordMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Record nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *RecordMutation) ResetField(name string) error {
	switch name {
	case record.FieldUserID:
		m.ResetUserID()
		return nil
	case record.FieldSignInIndex:
		m.ResetSignInIndex()
		return nil
	case record.FieldReward:
		m.ResetReward()
		return nil
	case record.FieldSignInDay:
		m.ResetSignInDay()
		return nil
	case record.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case record.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Record field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *RecordMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *RecordMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *RecordMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *RecordMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *RecordMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *RecordMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *RecordMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Record unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *RecordMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Record edge %s", name)
}

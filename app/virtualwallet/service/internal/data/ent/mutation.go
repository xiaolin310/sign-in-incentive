// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sign-in/app/virtualwallet/service/internal/data/ent/predicate"
	"sign-in/app/virtualwallet/service/internal/data/ent/virtualwallet"
	"sync"
	"time"

	"github.com/google/uuid"

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
	TypeVirtualWallet = "VirtualWallet"
)

// VirtualWalletMutation represents an operation that mutates the VirtualWallet nodes in the graph.
type VirtualWalletMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	user_id       *int64
	adduser_id    *int64
	balance       *float64
	addbalance    *float64
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*VirtualWallet, error)
	predicates    []predicate.VirtualWallet
}

var _ ent.Mutation = (*VirtualWalletMutation)(nil)

// virtualwalletOption allows management of the mutation configuration using functional options.
type virtualwalletOption func(*VirtualWalletMutation)

// newVirtualWalletMutation creates new mutation for the VirtualWallet entity.
func newVirtualWalletMutation(c config, op Op, opts ...virtualwalletOption) *VirtualWalletMutation {
	m := &VirtualWalletMutation{
		config:        c,
		op:            op,
		typ:           TypeVirtualWallet,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withVirtualWalletID sets the ID field of the mutation.
func withVirtualWalletID(id uuid.UUID) virtualwalletOption {
	return func(m *VirtualWalletMutation) {
		var (
			err   error
			once  sync.Once
			value *VirtualWallet
		)
		m.oldValue = func(ctx context.Context) (*VirtualWallet, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().VirtualWallet.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withVirtualWallet sets the old VirtualWallet of the mutation.
func withVirtualWallet(node *VirtualWallet) virtualwalletOption {
	return func(m *VirtualWalletMutation) {
		m.oldValue = func(context.Context) (*VirtualWallet, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m VirtualWalletMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m VirtualWalletMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of VirtualWallet entities.
func (m *VirtualWalletMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *VirtualWalletMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *VirtualWalletMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().VirtualWallet.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUserID sets the "user_id" field.
func (m *VirtualWalletMutation) SetUserID(i int64) {
	m.user_id = &i
	m.adduser_id = nil
}

// UserID returns the value of the "user_id" field in the mutation.
func (m *VirtualWalletMutation) UserID() (r int64, exists bool) {
	v := m.user_id
	if v == nil {
		return
	}
	return *v, true
}

// OldUserID returns the old "user_id" field's value of the VirtualWallet entity.
// If the VirtualWallet object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VirtualWalletMutation) OldUserID(ctx context.Context) (v int64, err error) {
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
func (m *VirtualWalletMutation) AddUserID(i int64) {
	if m.adduser_id != nil {
		*m.adduser_id += i
	} else {
		m.adduser_id = &i
	}
}

// AddedUserID returns the value that was added to the "user_id" field in this mutation.
func (m *VirtualWalletMutation) AddedUserID() (r int64, exists bool) {
	v := m.adduser_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetUserID resets all changes to the "user_id" field.
func (m *VirtualWalletMutation) ResetUserID() {
	m.user_id = nil
	m.adduser_id = nil
}

// SetBalance sets the "balance" field.
func (m *VirtualWalletMutation) SetBalance(f float64) {
	m.balance = &f
	m.addbalance = nil
}

// Balance returns the value of the "balance" field in the mutation.
func (m *VirtualWalletMutation) Balance() (r float64, exists bool) {
	v := m.balance
	if v == nil {
		return
	}
	return *v, true
}

// OldBalance returns the old "balance" field's value of the VirtualWallet entity.
// If the VirtualWallet object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VirtualWalletMutation) OldBalance(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBalance is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBalance requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBalance: %w", err)
	}
	return oldValue.Balance, nil
}

// AddBalance adds f to the "balance" field.
func (m *VirtualWalletMutation) AddBalance(f float64) {
	if m.addbalance != nil {
		*m.addbalance += f
	} else {
		m.addbalance = &f
	}
}

// AddedBalance returns the value that was added to the "balance" field in this mutation.
func (m *VirtualWalletMutation) AddedBalance() (r float64, exists bool) {
	v := m.addbalance
	if v == nil {
		return
	}
	return *v, true
}

// ResetBalance resets all changes to the "balance" field.
func (m *VirtualWalletMutation) ResetBalance() {
	m.balance = nil
	m.addbalance = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *VirtualWalletMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *VirtualWalletMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the VirtualWallet entity.
// If the VirtualWallet object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VirtualWalletMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
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
func (m *VirtualWalletMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *VirtualWalletMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *VirtualWalletMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the VirtualWallet entity.
// If the VirtualWallet object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VirtualWalletMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
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
func (m *VirtualWalletMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the VirtualWalletMutation builder.
func (m *VirtualWalletMutation) Where(ps ...predicate.VirtualWallet) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *VirtualWalletMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (VirtualWallet).
func (m *VirtualWalletMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *VirtualWalletMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.user_id != nil {
		fields = append(fields, virtualwallet.FieldUserID)
	}
	if m.balance != nil {
		fields = append(fields, virtualwallet.FieldBalance)
	}
	if m.created_at != nil {
		fields = append(fields, virtualwallet.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, virtualwallet.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *VirtualWalletMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case virtualwallet.FieldUserID:
		return m.UserID()
	case virtualwallet.FieldBalance:
		return m.Balance()
	case virtualwallet.FieldCreatedAt:
		return m.CreatedAt()
	case virtualwallet.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *VirtualWalletMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case virtualwallet.FieldUserID:
		return m.OldUserID(ctx)
	case virtualwallet.FieldBalance:
		return m.OldBalance(ctx)
	case virtualwallet.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case virtualwallet.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown VirtualWallet field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *VirtualWalletMutation) SetField(name string, value ent.Value) error {
	switch name {
	case virtualwallet.FieldUserID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserID(v)
		return nil
	case virtualwallet.FieldBalance:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBalance(v)
		return nil
	case virtualwallet.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case virtualwallet.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown VirtualWallet field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *VirtualWalletMutation) AddedFields() []string {
	var fields []string
	if m.adduser_id != nil {
		fields = append(fields, virtualwallet.FieldUserID)
	}
	if m.addbalance != nil {
		fields = append(fields, virtualwallet.FieldBalance)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *VirtualWalletMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case virtualwallet.FieldUserID:
		return m.AddedUserID()
	case virtualwallet.FieldBalance:
		return m.AddedBalance()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *VirtualWalletMutation) AddField(name string, value ent.Value) error {
	switch name {
	case virtualwallet.FieldUserID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddUserID(v)
		return nil
	case virtualwallet.FieldBalance:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddBalance(v)
		return nil
	}
	return fmt.Errorf("unknown VirtualWallet numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *VirtualWalletMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *VirtualWalletMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *VirtualWalletMutation) ClearField(name string) error {
	return fmt.Errorf("unknown VirtualWallet nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *VirtualWalletMutation) ResetField(name string) error {
	switch name {
	case virtualwallet.FieldUserID:
		m.ResetUserID()
		return nil
	case virtualwallet.FieldBalance:
		m.ResetBalance()
		return nil
	case virtualwallet.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case virtualwallet.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown VirtualWallet field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *VirtualWalletMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *VirtualWalletMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *VirtualWalletMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *VirtualWalletMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *VirtualWalletMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *VirtualWalletMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *VirtualWalletMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown VirtualWallet unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *VirtualWalletMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown VirtualWallet edge %s", name)
}
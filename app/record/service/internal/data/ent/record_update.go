// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sign-in/app/record/service/internal/data/ent/predicate"
	"sign-in/app/record/service/internal/data/ent/record"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RecordUpdate is the builder for updating Record entities.
type RecordUpdate struct {
	config
	hooks    []Hook
	mutation *RecordMutation
}

// Where appends a list predicates to the RecordUpdate builder.
func (ru *RecordUpdate) Where(ps ...predicate.Record) *RecordUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUserID sets the "user_id" field.
func (ru *RecordUpdate) SetUserID(i int64) *RecordUpdate {
	ru.mutation.ResetUserID()
	ru.mutation.SetUserID(i)
	return ru
}

// AddUserID adds i to the "user_id" field.
func (ru *RecordUpdate) AddUserID(i int64) *RecordUpdate {
	ru.mutation.AddUserID(i)
	return ru
}

// SetSignInIndex sets the "sign_in_index" field.
func (ru *RecordUpdate) SetSignInIndex(i int) *RecordUpdate {
	ru.mutation.ResetSignInIndex()
	ru.mutation.SetSignInIndex(i)
	return ru
}

// AddSignInIndex adds i to the "sign_in_index" field.
func (ru *RecordUpdate) AddSignInIndex(i int) *RecordUpdate {
	ru.mutation.AddSignInIndex(i)
	return ru
}

// SetReward sets the "reward" field.
func (ru *RecordUpdate) SetReward(f float64) *RecordUpdate {
	ru.mutation.ResetReward()
	ru.mutation.SetReward(f)
	return ru
}

// AddReward adds f to the "reward" field.
func (ru *RecordUpdate) AddReward(f float64) *RecordUpdate {
	ru.mutation.AddReward(f)
	return ru
}

// SetSignInDay sets the "sign_in_day" field.
func (ru *RecordUpdate) SetSignInDay(s string) *RecordUpdate {
	ru.mutation.SetSignInDay(s)
	return ru
}

// SetCreatedAt sets the "created_at" field.
func (ru *RecordUpdate) SetCreatedAt(t time.Time) *RecordUpdate {
	ru.mutation.SetCreatedAt(t)
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *RecordUpdate) SetUpdatedAt(t time.Time) *RecordUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ru *RecordUpdate) SetNillableUpdatedAt(t *time.Time) *RecordUpdate {
	if t != nil {
		ru.SetUpdatedAt(*t)
	}
	return ru
}

// Mutation returns the RecordMutation object of the builder.
func (ru *RecordUpdate) Mutation() *RecordMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RecordUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RecordUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RecordUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RecordUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RecordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   record.Table,
			Columns: record.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: record.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: record.FieldUserID,
		})
	}
	if value, ok := ru.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: record.FieldUserID,
		})
	}
	if value, ok := ru.mutation.SignInIndex(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: record.FieldSignInIndex,
		})
	}
	if value, ok := ru.mutation.AddedSignInIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: record.FieldSignInIndex,
		})
	}
	if value, ok := ru.mutation.Reward(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: record.FieldReward,
		})
	}
	if value, ok := ru.mutation.AddedReward(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: record.FieldReward,
		})
	}
	if value, ok := ru.mutation.SignInDay(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: record.FieldSignInDay,
		})
	}
	if value, ok := ru.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: record.FieldCreatedAt,
		})
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: record.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{record.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RecordUpdateOne is the builder for updating a single Record entity.
type RecordUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RecordMutation
}

// SetUserID sets the "user_id" field.
func (ruo *RecordUpdateOne) SetUserID(i int64) *RecordUpdateOne {
	ruo.mutation.ResetUserID()
	ruo.mutation.SetUserID(i)
	return ruo
}

// AddUserID adds i to the "user_id" field.
func (ruo *RecordUpdateOne) AddUserID(i int64) *RecordUpdateOne {
	ruo.mutation.AddUserID(i)
	return ruo
}

// SetSignInIndex sets the "sign_in_index" field.
func (ruo *RecordUpdateOne) SetSignInIndex(i int) *RecordUpdateOne {
	ruo.mutation.ResetSignInIndex()
	ruo.mutation.SetSignInIndex(i)
	return ruo
}

// AddSignInIndex adds i to the "sign_in_index" field.
func (ruo *RecordUpdateOne) AddSignInIndex(i int) *RecordUpdateOne {
	ruo.mutation.AddSignInIndex(i)
	return ruo
}

// SetReward sets the "reward" field.
func (ruo *RecordUpdateOne) SetReward(f float64) *RecordUpdateOne {
	ruo.mutation.ResetReward()
	ruo.mutation.SetReward(f)
	return ruo
}

// AddReward adds f to the "reward" field.
func (ruo *RecordUpdateOne) AddReward(f float64) *RecordUpdateOne {
	ruo.mutation.AddReward(f)
	return ruo
}

// SetSignInDay sets the "sign_in_day" field.
func (ruo *RecordUpdateOne) SetSignInDay(s string) *RecordUpdateOne {
	ruo.mutation.SetSignInDay(s)
	return ruo
}

// SetCreatedAt sets the "created_at" field.
func (ruo *RecordUpdateOne) SetCreatedAt(t time.Time) *RecordUpdateOne {
	ruo.mutation.SetCreatedAt(t)
	return ruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *RecordUpdateOne) SetUpdatedAt(t time.Time) *RecordUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ruo *RecordUpdateOne) SetNillableUpdatedAt(t *time.Time) *RecordUpdateOne {
	if t != nil {
		ruo.SetUpdatedAt(*t)
	}
	return ruo
}

// Mutation returns the RecordMutation object of the builder.
func (ruo *RecordUpdateOne) Mutation() *RecordMutation {
	return ruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RecordUpdateOne) Select(field string, fields ...string) *RecordUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Record entity.
func (ruo *RecordUpdateOne) Save(ctx context.Context) (*Record, error) {
	var (
		err  error
		node *Record
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RecordUpdateOne) SaveX(ctx context.Context) *Record {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RecordUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RecordUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RecordUpdateOne) sqlSave(ctx context.Context) (_node *Record, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   record.Table,
			Columns: record.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: record.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Record.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, record.FieldID)
		for _, f := range fields {
			if !record.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != record.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: record.FieldUserID,
		})
	}
	if value, ok := ruo.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: record.FieldUserID,
		})
	}
	if value, ok := ruo.mutation.SignInIndex(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: record.FieldSignInIndex,
		})
	}
	if value, ok := ruo.mutation.AddedSignInIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: record.FieldSignInIndex,
		})
	}
	if value, ok := ruo.mutation.Reward(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: record.FieldReward,
		})
	}
	if value, ok := ruo.mutation.AddedReward(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: record.FieldReward,
		})
	}
	if value, ok := ruo.mutation.SignInDay(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: record.FieldSignInDay,
		})
	}
	if value, ok := ruo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: record.FieldCreatedAt,
		})
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: record.FieldUpdatedAt,
		})
	}
	_node = &Record{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{record.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sign-in/app/virtualwallet/service/internal/data/ent/virtualwallet"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// VirtualWalletCreate is the builder for creating a VirtualWallet entity.
type VirtualWalletCreate struct {
	config
	mutation *VirtualWalletMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (vwc *VirtualWalletCreate) SetUserID(i int64) *VirtualWalletCreate {
	vwc.mutation.SetUserID(i)
	return vwc
}

// SetBalance sets the "balance" field.
func (vwc *VirtualWalletCreate) SetBalance(f float64) *VirtualWalletCreate {
	vwc.mutation.SetBalance(f)
	return vwc
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (vwc *VirtualWalletCreate) SetNillableBalance(f *float64) *VirtualWalletCreate {
	if f != nil {
		vwc.SetBalance(*f)
	}
	return vwc
}

// SetCreatedAt sets the "created_at" field.
func (vwc *VirtualWalletCreate) SetCreatedAt(t time.Time) *VirtualWalletCreate {
	vwc.mutation.SetCreatedAt(t)
	return vwc
}

// SetUpdatedAt sets the "updated_at" field.
func (vwc *VirtualWalletCreate) SetUpdatedAt(t time.Time) *VirtualWalletCreate {
	vwc.mutation.SetUpdatedAt(t)
	return vwc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (vwc *VirtualWalletCreate) SetNillableUpdatedAt(t *time.Time) *VirtualWalletCreate {
	if t != nil {
		vwc.SetUpdatedAt(*t)
	}
	return vwc
}

// SetID sets the "id" field.
func (vwc *VirtualWalletCreate) SetID(u uuid.UUID) *VirtualWalletCreate {
	vwc.mutation.SetID(u)
	return vwc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (vwc *VirtualWalletCreate) SetNillableID(u *uuid.UUID) *VirtualWalletCreate {
	if u != nil {
		vwc.SetID(*u)
	}
	return vwc
}

// Mutation returns the VirtualWalletMutation object of the builder.
func (vwc *VirtualWalletCreate) Mutation() *VirtualWalletMutation {
	return vwc.mutation
}

// Save creates the VirtualWallet in the database.
func (vwc *VirtualWalletCreate) Save(ctx context.Context) (*VirtualWallet, error) {
	var (
		err  error
		node *VirtualWallet
	)
	vwc.defaults()
	if len(vwc.hooks) == 0 {
		if err = vwc.check(); err != nil {
			return nil, err
		}
		node, err = vwc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VirtualWalletMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vwc.check(); err != nil {
				return nil, err
			}
			vwc.mutation = mutation
			if node, err = vwc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(vwc.hooks) - 1; i >= 0; i-- {
			if vwc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vwc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vwc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vwc *VirtualWalletCreate) SaveX(ctx context.Context) *VirtualWallet {
	v, err := vwc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vwc *VirtualWalletCreate) Exec(ctx context.Context) error {
	_, err := vwc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vwc *VirtualWalletCreate) ExecX(ctx context.Context) {
	if err := vwc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vwc *VirtualWalletCreate) defaults() {
	if _, ok := vwc.mutation.Balance(); !ok {
		v := virtualwallet.DefaultBalance
		vwc.mutation.SetBalance(v)
	}
	if _, ok := vwc.mutation.UpdatedAt(); !ok {
		v := virtualwallet.DefaultUpdatedAt()
		vwc.mutation.SetUpdatedAt(v)
	}
	if _, ok := vwc.mutation.ID(); !ok {
		v := virtualwallet.DefaultID()
		vwc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vwc *VirtualWalletCreate) check() error {
	if _, ok := vwc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "VirtualWallet.user_id"`)}
	}
	if _, ok := vwc.mutation.Balance(); !ok {
		return &ValidationError{Name: "balance", err: errors.New(`ent: missing required field "VirtualWallet.balance"`)}
	}
	if _, ok := vwc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "VirtualWallet.created_at"`)}
	}
	if _, ok := vwc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "VirtualWallet.updated_at"`)}
	}
	return nil
}

func (vwc *VirtualWalletCreate) sqlSave(ctx context.Context) (*VirtualWallet, error) {
	_node, _spec := vwc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vwc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (vwc *VirtualWalletCreate) createSpec() (*VirtualWallet, *sqlgraph.CreateSpec) {
	var (
		_node = &VirtualWallet{config: vwc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: virtualwallet.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: virtualwallet.FieldID,
			},
		}
	)
	if id, ok := vwc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := vwc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: virtualwallet.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := vwc.mutation.Balance(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: virtualwallet.FieldBalance,
		})
		_node.Balance = value
	}
	if value, ok := vwc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: virtualwallet.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := vwc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: virtualwallet.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// VirtualWalletCreateBulk is the builder for creating many VirtualWallet entities in bulk.
type VirtualWalletCreateBulk struct {
	config
	builders []*VirtualWalletCreate
}

// Save creates the VirtualWallet entities in the database.
func (vwcb *VirtualWalletCreateBulk) Save(ctx context.Context) ([]*VirtualWallet, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vwcb.builders))
	nodes := make([]*VirtualWallet, len(vwcb.builders))
	mutators := make([]Mutator, len(vwcb.builders))
	for i := range vwcb.builders {
		func(i int, root context.Context) {
			builder := vwcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VirtualWalletMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vwcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vwcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vwcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vwcb *VirtualWalletCreateBulk) SaveX(ctx context.Context) []*VirtualWallet {
	v, err := vwcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vwcb *VirtualWalletCreateBulk) Exec(ctx context.Context) error {
	_, err := vwcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vwcb *VirtualWalletCreateBulk) ExecX(ctx context.Context) {
	if err := vwcb.Exec(ctx); err != nil {
		panic(err)
	}
}

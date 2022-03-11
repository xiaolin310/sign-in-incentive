// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"sign-in/app/virtualwallet/service/internal/data/ent/predicate"
	"sign-in/app/virtualwallet/service/internal/data/ent/virtualwallet"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// VirtualWalletQuery is the builder for querying VirtualWallet entities.
type VirtualWalletQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.VirtualWallet
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VirtualWalletQuery builder.
func (vwq *VirtualWalletQuery) Where(ps ...predicate.VirtualWallet) *VirtualWalletQuery {
	vwq.predicates = append(vwq.predicates, ps...)
	return vwq
}

// Limit adds a limit step to the query.
func (vwq *VirtualWalletQuery) Limit(limit int) *VirtualWalletQuery {
	vwq.limit = &limit
	return vwq
}

// Offset adds an offset step to the query.
func (vwq *VirtualWalletQuery) Offset(offset int) *VirtualWalletQuery {
	vwq.offset = &offset
	return vwq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vwq *VirtualWalletQuery) Unique(unique bool) *VirtualWalletQuery {
	vwq.unique = &unique
	return vwq
}

// Order adds an order step to the query.
func (vwq *VirtualWalletQuery) Order(o ...OrderFunc) *VirtualWalletQuery {
	vwq.order = append(vwq.order, o...)
	return vwq
}

// First returns the first VirtualWallet entity from the query.
// Returns a *NotFoundError when no VirtualWallet was found.
func (vwq *VirtualWalletQuery) First(ctx context.Context) (*VirtualWallet, error) {
	nodes, err := vwq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{virtualwallet.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vwq *VirtualWalletQuery) FirstX(ctx context.Context) *VirtualWallet {
	node, err := vwq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VirtualWallet ID from the query.
// Returns a *NotFoundError when no VirtualWallet ID was found.
func (vwq *VirtualWalletQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vwq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{virtualwallet.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vwq *VirtualWalletQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := vwq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VirtualWallet entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one VirtualWallet entity is found.
// Returns a *NotFoundError when no VirtualWallet entities are found.
func (vwq *VirtualWalletQuery) Only(ctx context.Context) (*VirtualWallet, error) {
	nodes, err := vwq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{virtualwallet.Label}
	default:
		return nil, &NotSingularError{virtualwallet.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vwq *VirtualWalletQuery) OnlyX(ctx context.Context) *VirtualWallet {
	node, err := vwq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VirtualWallet ID in the query.
// Returns a *NotSingularError when more than one VirtualWallet ID is found.
// Returns a *NotFoundError when no entities are found.
func (vwq *VirtualWalletQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vwq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{virtualwallet.Label}
	default:
		err = &NotSingularError{virtualwallet.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vwq *VirtualWalletQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := vwq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VirtualWallets.
func (vwq *VirtualWalletQuery) All(ctx context.Context) ([]*VirtualWallet, error) {
	if err := vwq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return vwq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (vwq *VirtualWalletQuery) AllX(ctx context.Context) []*VirtualWallet {
	nodes, err := vwq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VirtualWallet IDs.
func (vwq *VirtualWalletQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := vwq.Select(virtualwallet.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vwq *VirtualWalletQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := vwq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vwq *VirtualWalletQuery) Count(ctx context.Context) (int, error) {
	if err := vwq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return vwq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (vwq *VirtualWalletQuery) CountX(ctx context.Context) int {
	count, err := vwq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vwq *VirtualWalletQuery) Exist(ctx context.Context) (bool, error) {
	if err := vwq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return vwq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (vwq *VirtualWalletQuery) ExistX(ctx context.Context) bool {
	exist, err := vwq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VirtualWalletQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vwq *VirtualWalletQuery) Clone() *VirtualWalletQuery {
	if vwq == nil {
		return nil
	}
	return &VirtualWalletQuery{
		config:     vwq.config,
		limit:      vwq.limit,
		offset:     vwq.offset,
		order:      append([]OrderFunc{}, vwq.order...),
		predicates: append([]predicate.VirtualWallet{}, vwq.predicates...),
		// clone intermediate query.
		sql:    vwq.sql.Clone(),
		path:   vwq.path,
		unique: vwq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID int64 `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.VirtualWallet.Query().
//		GroupBy(virtualwallet.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (vwq *VirtualWalletQuery) GroupBy(field string, fields ...string) *VirtualWalletGroupBy {
	group := &VirtualWalletGroupBy{config: vwq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := vwq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return vwq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID int64 `json:"user_id,omitempty"`
//	}
//
//	client.VirtualWallet.Query().
//		Select(virtualwallet.FieldUserID).
//		Scan(ctx, &v)
//
func (vwq *VirtualWalletQuery) Select(fields ...string) *VirtualWalletSelect {
	vwq.fields = append(vwq.fields, fields...)
	return &VirtualWalletSelect{VirtualWalletQuery: vwq}
}

func (vwq *VirtualWalletQuery) prepareQuery(ctx context.Context) error {
	for _, f := range vwq.fields {
		if !virtualwallet.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vwq.path != nil {
		prev, err := vwq.path(ctx)
		if err != nil {
			return err
		}
		vwq.sql = prev
	}
	return nil
}

func (vwq *VirtualWalletQuery) sqlAll(ctx context.Context) ([]*VirtualWallet, error) {
	var (
		nodes = []*VirtualWallet{}
		_spec = vwq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &VirtualWallet{config: vwq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, vwq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (vwq *VirtualWalletQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vwq.querySpec()
	_spec.Node.Columns = vwq.fields
	if len(vwq.fields) > 0 {
		_spec.Unique = vwq.unique != nil && *vwq.unique
	}
	return sqlgraph.CountNodes(ctx, vwq.driver, _spec)
}

func (vwq *VirtualWalletQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := vwq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (vwq *VirtualWalletQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   virtualwallet.Table,
			Columns: virtualwallet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: virtualwallet.FieldID,
			},
		},
		From:   vwq.sql,
		Unique: true,
	}
	if unique := vwq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := vwq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, virtualwallet.FieldID)
		for i := range fields {
			if fields[i] != virtualwallet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vwq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vwq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vwq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vwq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vwq *VirtualWalletQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vwq.driver.Dialect())
	t1 := builder.Table(virtualwallet.Table)
	columns := vwq.fields
	if len(columns) == 0 {
		columns = virtualwallet.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vwq.sql != nil {
		selector = vwq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vwq.unique != nil && *vwq.unique {
		selector.Distinct()
	}
	for _, p := range vwq.predicates {
		p(selector)
	}
	for _, p := range vwq.order {
		p(selector)
	}
	if offset := vwq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vwq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VirtualWalletGroupBy is the group-by builder for VirtualWallet entities.
type VirtualWalletGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vwgb *VirtualWalletGroupBy) Aggregate(fns ...AggregateFunc) *VirtualWalletGroupBy {
	vwgb.fns = append(vwgb.fns, fns...)
	return vwgb
}

// Scan applies the group-by query and scans the result into the given value.
func (vwgb *VirtualWalletGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := vwgb.path(ctx)
	if err != nil {
		return err
	}
	vwgb.sql = query
	return vwgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (vwgb *VirtualWalletGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := vwgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (vwgb *VirtualWalletGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(vwgb.fields) > 1 {
		return nil, errors.New("ent: VirtualWalletGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := vwgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (vwgb *VirtualWalletGroupBy) StringsX(ctx context.Context) []string {
	v, err := vwgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vwgb *VirtualWalletGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = vwgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{virtualwallet.Label}
	default:
		err = fmt.Errorf("ent: VirtualWalletGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (vwgb *VirtualWalletGroupBy) StringX(ctx context.Context) string {
	v, err := vwgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (vwgb *VirtualWalletGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(vwgb.fields) > 1 {
		return nil, errors.New("ent: VirtualWalletGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := vwgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (vwgb *VirtualWalletGroupBy) IntsX(ctx context.Context) []int {
	v, err := vwgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vwgb *VirtualWalletGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = vwgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{virtualwallet.Label}
	default:
		err = fmt.Errorf("ent: VirtualWalletGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (vwgb *VirtualWalletGroupBy) IntX(ctx context.Context) int {
	v, err := vwgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (vwgb *VirtualWalletGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(vwgb.fields) > 1 {
		return nil, errors.New("ent: VirtualWalletGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := vwgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (vwgb *VirtualWalletGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := vwgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vwgb *VirtualWalletGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = vwgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{virtualwallet.Label}
	default:
		err = fmt.Errorf("ent: VirtualWalletGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (vwgb *VirtualWalletGroupBy) Float64X(ctx context.Context) float64 {
	v, err := vwgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (vwgb *VirtualWalletGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(vwgb.fields) > 1 {
		return nil, errors.New("ent: VirtualWalletGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := vwgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (vwgb *VirtualWalletGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := vwgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vwgb *VirtualWalletGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = vwgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{virtualwallet.Label}
	default:
		err = fmt.Errorf("ent: VirtualWalletGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (vwgb *VirtualWalletGroupBy) BoolX(ctx context.Context) bool {
	v, err := vwgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (vwgb *VirtualWalletGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range vwgb.fields {
		if !virtualwallet.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := vwgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vwgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (vwgb *VirtualWalletGroupBy) sqlQuery() *sql.Selector {
	selector := vwgb.sql.Select()
	aggregation := make([]string, 0, len(vwgb.fns))
	for _, fn := range vwgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(vwgb.fields)+len(vwgb.fns))
		for _, f := range vwgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(vwgb.fields...)...)
}

// VirtualWalletSelect is the builder for selecting fields of VirtualWallet entities.
type VirtualWalletSelect struct {
	*VirtualWalletQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (vws *VirtualWalletSelect) Scan(ctx context.Context, v interface{}) error {
	if err := vws.prepareQuery(ctx); err != nil {
		return err
	}
	vws.sql = vws.VirtualWalletQuery.sqlQuery(ctx)
	return vws.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (vws *VirtualWalletSelect) ScanX(ctx context.Context, v interface{}) {
	if err := vws.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (vws *VirtualWalletSelect) Strings(ctx context.Context) ([]string, error) {
	if len(vws.fields) > 1 {
		return nil, errors.New("ent: VirtualWalletSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := vws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (vws *VirtualWalletSelect) StringsX(ctx context.Context) []string {
	v, err := vws.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (vws *VirtualWalletSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = vws.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{virtualwallet.Label}
	default:
		err = fmt.Errorf("ent: VirtualWalletSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (vws *VirtualWalletSelect) StringX(ctx context.Context) string {
	v, err := vws.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (vws *VirtualWalletSelect) Ints(ctx context.Context) ([]int, error) {
	if len(vws.fields) > 1 {
		return nil, errors.New("ent: VirtualWalletSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := vws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (vws *VirtualWalletSelect) IntsX(ctx context.Context) []int {
	v, err := vws.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (vws *VirtualWalletSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = vws.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{virtualwallet.Label}
	default:
		err = fmt.Errorf("ent: VirtualWalletSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (vws *VirtualWalletSelect) IntX(ctx context.Context) int {
	v, err := vws.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (vws *VirtualWalletSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(vws.fields) > 1 {
		return nil, errors.New("ent: VirtualWalletSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := vws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (vws *VirtualWalletSelect) Float64sX(ctx context.Context) []float64 {
	v, err := vws.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (vws *VirtualWalletSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = vws.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{virtualwallet.Label}
	default:
		err = fmt.Errorf("ent: VirtualWalletSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (vws *VirtualWalletSelect) Float64X(ctx context.Context) float64 {
	v, err := vws.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (vws *VirtualWalletSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(vws.fields) > 1 {
		return nil, errors.New("ent: VirtualWalletSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := vws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (vws *VirtualWalletSelect) BoolsX(ctx context.Context) []bool {
	v, err := vws.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (vws *VirtualWalletSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = vws.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{virtualwallet.Label}
	default:
		err = fmt.Errorf("ent: VirtualWalletSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (vws *VirtualWalletSelect) BoolX(ctx context.Context) bool {
	v, err := vws.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (vws *VirtualWalletSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := vws.sql.Query()
	if err := vws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
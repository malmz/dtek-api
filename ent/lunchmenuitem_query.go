// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dtekcth/dtek-api/ent/lunchmenu"
	"github.com/dtekcth/dtek-api/ent/lunchmenuitem"
	"github.com/dtekcth/dtek-api/ent/predicate"
)

// LunchMenuItemQuery is the builder for querying LunchMenuItem entities.
type LunchMenuItemQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.LunchMenuItem
	// eager-loading edges.
	withMenu *LunchMenuQuery
	withFKs  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LunchMenuItemQuery builder.
func (lmiq *LunchMenuItemQuery) Where(ps ...predicate.LunchMenuItem) *LunchMenuItemQuery {
	lmiq.predicates = append(lmiq.predicates, ps...)
	return lmiq
}

// Limit adds a limit step to the query.
func (lmiq *LunchMenuItemQuery) Limit(limit int) *LunchMenuItemQuery {
	lmiq.limit = &limit
	return lmiq
}

// Offset adds an offset step to the query.
func (lmiq *LunchMenuItemQuery) Offset(offset int) *LunchMenuItemQuery {
	lmiq.offset = &offset
	return lmiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lmiq *LunchMenuItemQuery) Unique(unique bool) *LunchMenuItemQuery {
	lmiq.unique = &unique
	return lmiq
}

// Order adds an order step to the query.
func (lmiq *LunchMenuItemQuery) Order(o ...OrderFunc) *LunchMenuItemQuery {
	lmiq.order = append(lmiq.order, o...)
	return lmiq
}

// QueryMenu chains the current query on the "menu" edge.
func (lmiq *LunchMenuItemQuery) QueryMenu() *LunchMenuQuery {
	query := &LunchMenuQuery{config: lmiq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lmiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lmiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(lunchmenuitem.Table, lunchmenuitem.FieldID, selector),
			sqlgraph.To(lunchmenu.Table, lunchmenu.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, lunchmenuitem.MenuTable, lunchmenuitem.MenuColumn),
		)
		fromU = sqlgraph.SetNeighbors(lmiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LunchMenuItem entity from the query.
// Returns a *NotFoundError when no LunchMenuItem was found.
func (lmiq *LunchMenuItemQuery) First(ctx context.Context) (*LunchMenuItem, error) {
	nodes, err := lmiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{lunchmenuitem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lmiq *LunchMenuItemQuery) FirstX(ctx context.Context) *LunchMenuItem {
	node, err := lmiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LunchMenuItem ID from the query.
// Returns a *NotFoundError when no LunchMenuItem ID was found.
func (lmiq *LunchMenuItemQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lmiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{lunchmenuitem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lmiq *LunchMenuItemQuery) FirstIDX(ctx context.Context) int {
	id, err := lmiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LunchMenuItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LunchMenuItem entity is found.
// Returns a *NotFoundError when no LunchMenuItem entities are found.
func (lmiq *LunchMenuItemQuery) Only(ctx context.Context) (*LunchMenuItem, error) {
	nodes, err := lmiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{lunchmenuitem.Label}
	default:
		return nil, &NotSingularError{lunchmenuitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lmiq *LunchMenuItemQuery) OnlyX(ctx context.Context) *LunchMenuItem {
	node, err := lmiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LunchMenuItem ID in the query.
// Returns a *NotSingularError when more than one LunchMenuItem ID is found.
// Returns a *NotFoundError when no entities are found.
func (lmiq *LunchMenuItemQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lmiq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{lunchmenuitem.Label}
	default:
		err = &NotSingularError{lunchmenuitem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lmiq *LunchMenuItemQuery) OnlyIDX(ctx context.Context) int {
	id, err := lmiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LunchMenuItems.
func (lmiq *LunchMenuItemQuery) All(ctx context.Context) ([]*LunchMenuItem, error) {
	if err := lmiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return lmiq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (lmiq *LunchMenuItemQuery) AllX(ctx context.Context) []*LunchMenuItem {
	nodes, err := lmiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LunchMenuItem IDs.
func (lmiq *LunchMenuItemQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := lmiq.Select(lunchmenuitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lmiq *LunchMenuItemQuery) IDsX(ctx context.Context) []int {
	ids, err := lmiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lmiq *LunchMenuItemQuery) Count(ctx context.Context) (int, error) {
	if err := lmiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return lmiq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (lmiq *LunchMenuItemQuery) CountX(ctx context.Context) int {
	count, err := lmiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lmiq *LunchMenuItemQuery) Exist(ctx context.Context) (bool, error) {
	if err := lmiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return lmiq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (lmiq *LunchMenuItemQuery) ExistX(ctx context.Context) bool {
	exist, err := lmiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LunchMenuItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lmiq *LunchMenuItemQuery) Clone() *LunchMenuItemQuery {
	if lmiq == nil {
		return nil
	}
	return &LunchMenuItemQuery{
		config:     lmiq.config,
		limit:      lmiq.limit,
		offset:     lmiq.offset,
		order:      append([]OrderFunc{}, lmiq.order...),
		predicates: append([]predicate.LunchMenuItem{}, lmiq.predicates...),
		withMenu:   lmiq.withMenu.Clone(),
		// clone intermediate query.
		sql:    lmiq.sql.Clone(),
		path:   lmiq.path,
		unique: lmiq.unique,
	}
}

// WithMenu tells the query-builder to eager-load the nodes that are connected to
// the "menu" edge. The optional arguments are used to configure the query builder of the edge.
func (lmiq *LunchMenuItemQuery) WithMenu(opts ...func(*LunchMenuQuery)) *LunchMenuItemQuery {
	query := &LunchMenuQuery{config: lmiq.config}
	for _, opt := range opts {
		opt(query)
	}
	lmiq.withMenu = query
	return lmiq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LunchMenuItem.Query().
//		GroupBy(lunchmenuitem.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (lmiq *LunchMenuItemQuery) GroupBy(field string, fields ...string) *LunchMenuItemGroupBy {
	grbuild := &LunchMenuItemGroupBy{config: lmiq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := lmiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return lmiq.sqlQuery(ctx), nil
	}
	grbuild.label = lunchmenuitem.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.LunchMenuItem.Query().
//		Select(lunchmenuitem.FieldTitle).
//		Scan(ctx, &v)
//
func (lmiq *LunchMenuItemQuery) Select(fields ...string) *LunchMenuItemSelect {
	lmiq.fields = append(lmiq.fields, fields...)
	selbuild := &LunchMenuItemSelect{LunchMenuItemQuery: lmiq}
	selbuild.label = lunchmenuitem.Label
	selbuild.flds, selbuild.scan = &lmiq.fields, selbuild.Scan
	return selbuild
}

func (lmiq *LunchMenuItemQuery) prepareQuery(ctx context.Context) error {
	for _, f := range lmiq.fields {
		if !lunchmenuitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lmiq.path != nil {
		prev, err := lmiq.path(ctx)
		if err != nil {
			return err
		}
		lmiq.sql = prev
	}
	return nil
}

func (lmiq *LunchMenuItemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LunchMenuItem, error) {
	var (
		nodes       = []*LunchMenuItem{}
		withFKs     = lmiq.withFKs
		_spec       = lmiq.querySpec()
		loadedTypes = [1]bool{
			lmiq.withMenu != nil,
		}
	)
	if lmiq.withMenu != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, lunchmenuitem.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*LunchMenuItem).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &LunchMenuItem{config: lmiq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lmiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := lmiq.withMenu; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*LunchMenuItem)
		for i := range nodes {
			if nodes[i].lunch_menu_items == nil {
				continue
			}
			fk := *nodes[i].lunch_menu_items
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(lunchmenu.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "lunch_menu_items" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Menu = n
			}
		}
	}

	return nodes, nil
}

func (lmiq *LunchMenuItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lmiq.querySpec()
	_spec.Node.Columns = lmiq.fields
	if len(lmiq.fields) > 0 {
		_spec.Unique = lmiq.unique != nil && *lmiq.unique
	}
	return sqlgraph.CountNodes(ctx, lmiq.driver, _spec)
}

func (lmiq *LunchMenuItemQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := lmiq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (lmiq *LunchMenuItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lunchmenuitem.Table,
			Columns: lunchmenuitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lunchmenuitem.FieldID,
			},
		},
		From:   lmiq.sql,
		Unique: true,
	}
	if unique := lmiq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := lmiq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lunchmenuitem.FieldID)
		for i := range fields {
			if fields[i] != lunchmenuitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lmiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lmiq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lmiq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lmiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lmiq *LunchMenuItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lmiq.driver.Dialect())
	t1 := builder.Table(lunchmenuitem.Table)
	columns := lmiq.fields
	if len(columns) == 0 {
		columns = lunchmenuitem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lmiq.sql != nil {
		selector = lmiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lmiq.unique != nil && *lmiq.unique {
		selector.Distinct()
	}
	for _, p := range lmiq.predicates {
		p(selector)
	}
	for _, p := range lmiq.order {
		p(selector)
	}
	if offset := lmiq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lmiq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LunchMenuItemGroupBy is the group-by builder for LunchMenuItem entities.
type LunchMenuItemGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lmigb *LunchMenuItemGroupBy) Aggregate(fns ...AggregateFunc) *LunchMenuItemGroupBy {
	lmigb.fns = append(lmigb.fns, fns...)
	return lmigb
}

// Scan applies the group-by query and scans the result into the given value.
func (lmigb *LunchMenuItemGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := lmigb.path(ctx)
	if err != nil {
		return err
	}
	lmigb.sql = query
	return lmigb.sqlScan(ctx, v)
}

func (lmigb *LunchMenuItemGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range lmigb.fields {
		if !lunchmenuitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := lmigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lmigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (lmigb *LunchMenuItemGroupBy) sqlQuery() *sql.Selector {
	selector := lmigb.sql.Select()
	aggregation := make([]string, 0, len(lmigb.fns))
	for _, fn := range lmigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(lmigb.fields)+len(lmigb.fns))
		for _, f := range lmigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(lmigb.fields...)...)
}

// LunchMenuItemSelect is the builder for selecting fields of LunchMenuItem entities.
type LunchMenuItemSelect struct {
	*LunchMenuItemQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (lmis *LunchMenuItemSelect) Scan(ctx context.Context, v interface{}) error {
	if err := lmis.prepareQuery(ctx); err != nil {
		return err
	}
	lmis.sql = lmis.LunchMenuItemQuery.sqlQuery(ctx)
	return lmis.sqlScan(ctx, v)
}

func (lmis *LunchMenuItemSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := lmis.sql.Query()
	if err := lmis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
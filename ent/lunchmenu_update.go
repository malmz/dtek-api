// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/dtekcth/dtek-api/ent/lunchmenu"
	"github.com/dtekcth/dtek-api/ent/predicate"
	"github.com/dtekcth/dtek-api/ent/schema"
)

// LunchMenuUpdate is the builder for updating LunchMenu entities.
type LunchMenuUpdate struct {
	config
	hooks    []Hook
	mutation *LunchMenuMutation
}

// Where appends a list predicates to the LunchMenuUpdate builder.
func (lmu *LunchMenuUpdate) Where(ps ...predicate.LunchMenu) *LunchMenuUpdate {
	lmu.mutation.Where(ps...)
	return lmu
}

// SetUpdateTime sets the "update_time" field.
func (lmu *LunchMenuUpdate) SetUpdateTime(t time.Time) *LunchMenuUpdate {
	lmu.mutation.SetUpdateTime(t)
	return lmu
}

// SetResturant sets the "resturant" field.
func (lmu *LunchMenuUpdate) SetResturant(s string) *LunchMenuUpdate {
	lmu.mutation.SetResturant(s)
	return lmu
}

// SetDate sets the "date" field.
func (lmu *LunchMenuUpdate) SetDate(t time.Time) *LunchMenuUpdate {
	lmu.mutation.SetDate(t)
	return lmu
}

// SetLanguage sets the "language" field.
func (lmu *LunchMenuUpdate) SetLanguage(l lunchmenu.Language) *LunchMenuUpdate {
	lmu.mutation.SetLanguage(l)
	return lmu
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (lmu *LunchMenuUpdate) SetNillableLanguage(l *lunchmenu.Language) *LunchMenuUpdate {
	if l != nil {
		lmu.SetLanguage(*l)
	}
	return lmu
}

// ClearLanguage clears the value of the "language" field.
func (lmu *LunchMenuUpdate) ClearLanguage() *LunchMenuUpdate {
	lmu.mutation.ClearLanguage()
	return lmu
}

// SetName sets the "name" field.
func (lmu *LunchMenuUpdate) SetName(s string) *LunchMenuUpdate {
	lmu.mutation.SetName(s)
	return lmu
}

// SetMenu sets the "menu" field.
func (lmu *LunchMenuUpdate) SetMenu(smi []schema.LunchMenuItem) *LunchMenuUpdate {
	lmu.mutation.SetMenu(smi)
	return lmu
}

// AppendMenu appends smi to the "menu" field.
func (lmu *LunchMenuUpdate) AppendMenu(smi []schema.LunchMenuItem) *LunchMenuUpdate {
	lmu.mutation.AppendMenu(smi)
	return lmu
}

// Mutation returns the LunchMenuMutation object of the builder.
func (lmu *LunchMenuUpdate) Mutation() *LunchMenuMutation {
	return lmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lmu *LunchMenuUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	lmu.defaults()
	if len(lmu.hooks) == 0 {
		if err = lmu.check(); err != nil {
			return 0, err
		}
		affected, err = lmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LunchMenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lmu.check(); err != nil {
				return 0, err
			}
			lmu.mutation = mutation
			affected, err = lmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lmu.hooks) - 1; i >= 0; i-- {
			if lmu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lmu *LunchMenuUpdate) SaveX(ctx context.Context) int {
	affected, err := lmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lmu *LunchMenuUpdate) Exec(ctx context.Context) error {
	_, err := lmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lmu *LunchMenuUpdate) ExecX(ctx context.Context) {
	if err := lmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lmu *LunchMenuUpdate) defaults() {
	if _, ok := lmu.mutation.UpdateTime(); !ok {
		v := lunchmenu.UpdateDefaultUpdateTime()
		lmu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lmu *LunchMenuUpdate) check() error {
	if v, ok := lmu.mutation.Language(); ok {
		if err := lunchmenu.LanguageValidator(v); err != nil {
			return &ValidationError{Name: "language", err: fmt.Errorf(`ent: validator failed for field "LunchMenu.language": %w`, err)}
		}
	}
	return nil
}

func (lmu *LunchMenuUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lunchmenu.Table,
			Columns: lunchmenu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lunchmenu.FieldID,
			},
		},
	}
	if ps := lmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lmu.mutation.UpdateTime(); ok {
		_spec.SetField(lunchmenu.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := lmu.mutation.Resturant(); ok {
		_spec.SetField(lunchmenu.FieldResturant, field.TypeString, value)
	}
	if value, ok := lmu.mutation.Date(); ok {
		_spec.SetField(lunchmenu.FieldDate, field.TypeTime, value)
	}
	if value, ok := lmu.mutation.Language(); ok {
		_spec.SetField(lunchmenu.FieldLanguage, field.TypeEnum, value)
	}
	if lmu.mutation.LanguageCleared() {
		_spec.ClearField(lunchmenu.FieldLanguage, field.TypeEnum)
	}
	if value, ok := lmu.mutation.Name(); ok {
		_spec.SetField(lunchmenu.FieldName, field.TypeString, value)
	}
	if value, ok := lmu.mutation.Menu(); ok {
		_spec.SetField(lunchmenu.FieldMenu, field.TypeJSON, value)
	}
	if value, ok := lmu.mutation.AppendedMenu(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, lunchmenu.FieldMenu, value)
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lunchmenu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// LunchMenuUpdateOne is the builder for updating a single LunchMenu entity.
type LunchMenuUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LunchMenuMutation
}

// SetUpdateTime sets the "update_time" field.
func (lmuo *LunchMenuUpdateOne) SetUpdateTime(t time.Time) *LunchMenuUpdateOne {
	lmuo.mutation.SetUpdateTime(t)
	return lmuo
}

// SetResturant sets the "resturant" field.
func (lmuo *LunchMenuUpdateOne) SetResturant(s string) *LunchMenuUpdateOne {
	lmuo.mutation.SetResturant(s)
	return lmuo
}

// SetDate sets the "date" field.
func (lmuo *LunchMenuUpdateOne) SetDate(t time.Time) *LunchMenuUpdateOne {
	lmuo.mutation.SetDate(t)
	return lmuo
}

// SetLanguage sets the "language" field.
func (lmuo *LunchMenuUpdateOne) SetLanguage(l lunchmenu.Language) *LunchMenuUpdateOne {
	lmuo.mutation.SetLanguage(l)
	return lmuo
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (lmuo *LunchMenuUpdateOne) SetNillableLanguage(l *lunchmenu.Language) *LunchMenuUpdateOne {
	if l != nil {
		lmuo.SetLanguage(*l)
	}
	return lmuo
}

// ClearLanguage clears the value of the "language" field.
func (lmuo *LunchMenuUpdateOne) ClearLanguage() *LunchMenuUpdateOne {
	lmuo.mutation.ClearLanguage()
	return lmuo
}

// SetName sets the "name" field.
func (lmuo *LunchMenuUpdateOne) SetName(s string) *LunchMenuUpdateOne {
	lmuo.mutation.SetName(s)
	return lmuo
}

// SetMenu sets the "menu" field.
func (lmuo *LunchMenuUpdateOne) SetMenu(smi []schema.LunchMenuItem) *LunchMenuUpdateOne {
	lmuo.mutation.SetMenu(smi)
	return lmuo
}

// AppendMenu appends smi to the "menu" field.
func (lmuo *LunchMenuUpdateOne) AppendMenu(smi []schema.LunchMenuItem) *LunchMenuUpdateOne {
	lmuo.mutation.AppendMenu(smi)
	return lmuo
}

// Mutation returns the LunchMenuMutation object of the builder.
func (lmuo *LunchMenuUpdateOne) Mutation() *LunchMenuMutation {
	return lmuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lmuo *LunchMenuUpdateOne) Select(field string, fields ...string) *LunchMenuUpdateOne {
	lmuo.fields = append([]string{field}, fields...)
	return lmuo
}

// Save executes the query and returns the updated LunchMenu entity.
func (lmuo *LunchMenuUpdateOne) Save(ctx context.Context) (*LunchMenu, error) {
	var (
		err  error
		node *LunchMenu
	)
	lmuo.defaults()
	if len(lmuo.hooks) == 0 {
		if err = lmuo.check(); err != nil {
			return nil, err
		}
		node, err = lmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LunchMenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lmuo.check(); err != nil {
				return nil, err
			}
			lmuo.mutation = mutation
			node, err = lmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(lmuo.hooks) - 1; i >= 0; i-- {
			if lmuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lmuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, lmuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*LunchMenu)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from LunchMenuMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (lmuo *LunchMenuUpdateOne) SaveX(ctx context.Context) *LunchMenu {
	node, err := lmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lmuo *LunchMenuUpdateOne) Exec(ctx context.Context) error {
	_, err := lmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lmuo *LunchMenuUpdateOne) ExecX(ctx context.Context) {
	if err := lmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lmuo *LunchMenuUpdateOne) defaults() {
	if _, ok := lmuo.mutation.UpdateTime(); !ok {
		v := lunchmenu.UpdateDefaultUpdateTime()
		lmuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lmuo *LunchMenuUpdateOne) check() error {
	if v, ok := lmuo.mutation.Language(); ok {
		if err := lunchmenu.LanguageValidator(v); err != nil {
			return &ValidationError{Name: "language", err: fmt.Errorf(`ent: validator failed for field "LunchMenu.language": %w`, err)}
		}
	}
	return nil
}

func (lmuo *LunchMenuUpdateOne) sqlSave(ctx context.Context) (_node *LunchMenu, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lunchmenu.Table,
			Columns: lunchmenu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lunchmenu.FieldID,
			},
		},
	}
	id, ok := lmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LunchMenu.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lunchmenu.FieldID)
		for _, f := range fields {
			if !lunchmenu.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != lunchmenu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lmuo.mutation.UpdateTime(); ok {
		_spec.SetField(lunchmenu.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := lmuo.mutation.Resturant(); ok {
		_spec.SetField(lunchmenu.FieldResturant, field.TypeString, value)
	}
	if value, ok := lmuo.mutation.Date(); ok {
		_spec.SetField(lunchmenu.FieldDate, field.TypeTime, value)
	}
	if value, ok := lmuo.mutation.Language(); ok {
		_spec.SetField(lunchmenu.FieldLanguage, field.TypeEnum, value)
	}
	if lmuo.mutation.LanguageCleared() {
		_spec.ClearField(lunchmenu.FieldLanguage, field.TypeEnum)
	}
	if value, ok := lmuo.mutation.Name(); ok {
		_spec.SetField(lunchmenu.FieldName, field.TypeString, value)
	}
	if value, ok := lmuo.mutation.Menu(); ok {
		_spec.SetField(lunchmenu.FieldMenu, field.TypeJSON, value)
	}
	if value, ok := lmuo.mutation.AppendedMenu(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, lunchmenu.FieldMenu, value)
		})
	}
	_node = &LunchMenu{config: lmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lunchmenu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

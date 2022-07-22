// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dtekcth/dtek-api/ent/lunchmenu"
	"github.com/dtekcth/dtek-api/ent/lunchmenuitem"
	"github.com/dtekcth/dtek-api/ent/predicate"
	"github.com/dtekcth/dtek-api/model"
)

// LunchMenuItemUpdate is the builder for updating LunchMenuItem entities.
type LunchMenuItemUpdate struct {
	config
	hooks    []Hook
	mutation *LunchMenuItemMutation
}

// Where appends a list predicates to the LunchMenuItemUpdate builder.
func (lmiu *LunchMenuItemUpdate) Where(ps ...predicate.LunchMenuItem) *LunchMenuItemUpdate {
	lmiu.mutation.Where(ps...)
	return lmiu
}

// SetTitle sets the "title" field.
func (lmiu *LunchMenuItemUpdate) SetTitle(s string) *LunchMenuItemUpdate {
	lmiu.mutation.SetTitle(s)
	return lmiu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (lmiu *LunchMenuItemUpdate) SetNillableTitle(s *string) *LunchMenuItemUpdate {
	if s != nil {
		lmiu.SetTitle(*s)
	}
	return lmiu
}

// ClearTitle clears the value of the "title" field.
func (lmiu *LunchMenuItemUpdate) ClearTitle() *LunchMenuItemUpdate {
	lmiu.mutation.ClearTitle()
	return lmiu
}

// SetBody sets the "body" field.
func (lmiu *LunchMenuItemUpdate) SetBody(s string) *LunchMenuItemUpdate {
	lmiu.mutation.SetBody(s)
	return lmiu
}

// SetLanguage sets the "language" field.
func (lmiu *LunchMenuItemUpdate) SetLanguage(l lunchmenuitem.Language) *LunchMenuItemUpdate {
	lmiu.mutation.SetLanguage(l)
	return lmiu
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (lmiu *LunchMenuItemUpdate) SetNillableLanguage(l *lunchmenuitem.Language) *LunchMenuItemUpdate {
	if l != nil {
		lmiu.SetLanguage(*l)
	}
	return lmiu
}

// ClearLanguage clears the value of the "language" field.
func (lmiu *LunchMenuItemUpdate) ClearLanguage() *LunchMenuItemUpdate {
	lmiu.mutation.ClearLanguage()
	return lmiu
}

// SetPreformatted sets the "preformatted" field.
func (lmiu *LunchMenuItemUpdate) SetPreformatted(b bool) *LunchMenuItemUpdate {
	lmiu.mutation.SetPreformatted(b)
	return lmiu
}

// SetNillablePreformatted sets the "preformatted" field if the given value is not nil.
func (lmiu *LunchMenuItemUpdate) SetNillablePreformatted(b *bool) *LunchMenuItemUpdate {
	if b != nil {
		lmiu.SetPreformatted(*b)
	}
	return lmiu
}

// SetAllergens sets the "allergens" field.
func (lmiu *LunchMenuItemUpdate) SetAllergens(m []model.Allergen) *LunchMenuItemUpdate {
	lmiu.mutation.SetAllergens(m)
	return lmiu
}

// ClearAllergens clears the value of the "allergens" field.
func (lmiu *LunchMenuItemUpdate) ClearAllergens() *LunchMenuItemUpdate {
	lmiu.mutation.ClearAllergens()
	return lmiu
}

// SetEmission sets the "emission" field.
func (lmiu *LunchMenuItemUpdate) SetEmission(f float64) *LunchMenuItemUpdate {
	lmiu.mutation.ResetEmission()
	lmiu.mutation.SetEmission(f)
	return lmiu
}

// SetNillableEmission sets the "emission" field if the given value is not nil.
func (lmiu *LunchMenuItemUpdate) SetNillableEmission(f *float64) *LunchMenuItemUpdate {
	if f != nil {
		lmiu.SetEmission(*f)
	}
	return lmiu
}

// AddEmission adds f to the "emission" field.
func (lmiu *LunchMenuItemUpdate) AddEmission(f float64) *LunchMenuItemUpdate {
	lmiu.mutation.AddEmission(f)
	return lmiu
}

// ClearEmission clears the value of the "emission" field.
func (lmiu *LunchMenuItemUpdate) ClearEmission() *LunchMenuItemUpdate {
	lmiu.mutation.ClearEmission()
	return lmiu
}

// SetPrice sets the "price" field.
func (lmiu *LunchMenuItemUpdate) SetPrice(s string) *LunchMenuItemUpdate {
	lmiu.mutation.SetPrice(s)
	return lmiu
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (lmiu *LunchMenuItemUpdate) SetNillablePrice(s *string) *LunchMenuItemUpdate {
	if s != nil {
		lmiu.SetPrice(*s)
	}
	return lmiu
}

// ClearPrice clears the value of the "price" field.
func (lmiu *LunchMenuItemUpdate) ClearPrice() *LunchMenuItemUpdate {
	lmiu.mutation.ClearPrice()
	return lmiu
}

// SetMenuID sets the "menu" edge to the LunchMenu entity by ID.
func (lmiu *LunchMenuItemUpdate) SetMenuID(id int) *LunchMenuItemUpdate {
	lmiu.mutation.SetMenuID(id)
	return lmiu
}

// SetMenu sets the "menu" edge to the LunchMenu entity.
func (lmiu *LunchMenuItemUpdate) SetMenu(l *LunchMenu) *LunchMenuItemUpdate {
	return lmiu.SetMenuID(l.ID)
}

// Mutation returns the LunchMenuItemMutation object of the builder.
func (lmiu *LunchMenuItemUpdate) Mutation() *LunchMenuItemMutation {
	return lmiu.mutation
}

// ClearMenu clears the "menu" edge to the LunchMenu entity.
func (lmiu *LunchMenuItemUpdate) ClearMenu() *LunchMenuItemUpdate {
	lmiu.mutation.ClearMenu()
	return lmiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lmiu *LunchMenuItemUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lmiu.hooks) == 0 {
		if err = lmiu.check(); err != nil {
			return 0, err
		}
		affected, err = lmiu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LunchMenuItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lmiu.check(); err != nil {
				return 0, err
			}
			lmiu.mutation = mutation
			affected, err = lmiu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lmiu.hooks) - 1; i >= 0; i-- {
			if lmiu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lmiu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lmiu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lmiu *LunchMenuItemUpdate) SaveX(ctx context.Context) int {
	affected, err := lmiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lmiu *LunchMenuItemUpdate) Exec(ctx context.Context) error {
	_, err := lmiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lmiu *LunchMenuItemUpdate) ExecX(ctx context.Context) {
	if err := lmiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lmiu *LunchMenuItemUpdate) check() error {
	if v, ok := lmiu.mutation.Language(); ok {
		if err := lunchmenuitem.LanguageValidator(v); err != nil {
			return &ValidationError{Name: "language", err: fmt.Errorf(`ent: validator failed for field "LunchMenuItem.language": %w`, err)}
		}
	}
	if _, ok := lmiu.mutation.MenuID(); lmiu.mutation.MenuCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LunchMenuItem.menu"`)
	}
	return nil
}

func (lmiu *LunchMenuItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lunchmenuitem.Table,
			Columns: lunchmenuitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lunchmenuitem.FieldID,
			},
		},
	}
	if ps := lmiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lmiu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lunchmenuitem.FieldTitle,
		})
	}
	if lmiu.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: lunchmenuitem.FieldTitle,
		})
	}
	if value, ok := lmiu.mutation.Body(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lunchmenuitem.FieldBody,
		})
	}
	if value, ok := lmiu.mutation.Language(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: lunchmenuitem.FieldLanguage,
		})
	}
	if lmiu.mutation.LanguageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: lunchmenuitem.FieldLanguage,
		})
	}
	if value, ok := lmiu.mutation.Preformatted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: lunchmenuitem.FieldPreformatted,
		})
	}
	if value, ok := lmiu.mutation.Allergens(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: lunchmenuitem.FieldAllergens,
		})
	}
	if lmiu.mutation.AllergensCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: lunchmenuitem.FieldAllergens,
		})
	}
	if value, ok := lmiu.mutation.Emission(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: lunchmenuitem.FieldEmission,
		})
	}
	if value, ok := lmiu.mutation.AddedEmission(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: lunchmenuitem.FieldEmission,
		})
	}
	if lmiu.mutation.EmissionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: lunchmenuitem.FieldEmission,
		})
	}
	if value, ok := lmiu.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lunchmenuitem.FieldPrice,
		})
	}
	if lmiu.mutation.PriceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: lunchmenuitem.FieldPrice,
		})
	}
	if lmiu.mutation.MenuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lunchmenuitem.MenuTable,
			Columns: []string{lunchmenuitem.MenuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lunchmenu.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmiu.mutation.MenuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lunchmenuitem.MenuTable,
			Columns: []string{lunchmenuitem.MenuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lunchmenu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lmiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lunchmenuitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// LunchMenuItemUpdateOne is the builder for updating a single LunchMenuItem entity.
type LunchMenuItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LunchMenuItemMutation
}

// SetTitle sets the "title" field.
func (lmiuo *LunchMenuItemUpdateOne) SetTitle(s string) *LunchMenuItemUpdateOne {
	lmiuo.mutation.SetTitle(s)
	return lmiuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (lmiuo *LunchMenuItemUpdateOne) SetNillableTitle(s *string) *LunchMenuItemUpdateOne {
	if s != nil {
		lmiuo.SetTitle(*s)
	}
	return lmiuo
}

// ClearTitle clears the value of the "title" field.
func (lmiuo *LunchMenuItemUpdateOne) ClearTitle() *LunchMenuItemUpdateOne {
	lmiuo.mutation.ClearTitle()
	return lmiuo
}

// SetBody sets the "body" field.
func (lmiuo *LunchMenuItemUpdateOne) SetBody(s string) *LunchMenuItemUpdateOne {
	lmiuo.mutation.SetBody(s)
	return lmiuo
}

// SetLanguage sets the "language" field.
func (lmiuo *LunchMenuItemUpdateOne) SetLanguage(l lunchmenuitem.Language) *LunchMenuItemUpdateOne {
	lmiuo.mutation.SetLanguage(l)
	return lmiuo
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (lmiuo *LunchMenuItemUpdateOne) SetNillableLanguage(l *lunchmenuitem.Language) *LunchMenuItemUpdateOne {
	if l != nil {
		lmiuo.SetLanguage(*l)
	}
	return lmiuo
}

// ClearLanguage clears the value of the "language" field.
func (lmiuo *LunchMenuItemUpdateOne) ClearLanguage() *LunchMenuItemUpdateOne {
	lmiuo.mutation.ClearLanguage()
	return lmiuo
}

// SetPreformatted sets the "preformatted" field.
func (lmiuo *LunchMenuItemUpdateOne) SetPreformatted(b bool) *LunchMenuItemUpdateOne {
	lmiuo.mutation.SetPreformatted(b)
	return lmiuo
}

// SetNillablePreformatted sets the "preformatted" field if the given value is not nil.
func (lmiuo *LunchMenuItemUpdateOne) SetNillablePreformatted(b *bool) *LunchMenuItemUpdateOne {
	if b != nil {
		lmiuo.SetPreformatted(*b)
	}
	return lmiuo
}

// SetAllergens sets the "allergens" field.
func (lmiuo *LunchMenuItemUpdateOne) SetAllergens(m []model.Allergen) *LunchMenuItemUpdateOne {
	lmiuo.mutation.SetAllergens(m)
	return lmiuo
}

// ClearAllergens clears the value of the "allergens" field.
func (lmiuo *LunchMenuItemUpdateOne) ClearAllergens() *LunchMenuItemUpdateOne {
	lmiuo.mutation.ClearAllergens()
	return lmiuo
}

// SetEmission sets the "emission" field.
func (lmiuo *LunchMenuItemUpdateOne) SetEmission(f float64) *LunchMenuItemUpdateOne {
	lmiuo.mutation.ResetEmission()
	lmiuo.mutation.SetEmission(f)
	return lmiuo
}

// SetNillableEmission sets the "emission" field if the given value is not nil.
func (lmiuo *LunchMenuItemUpdateOne) SetNillableEmission(f *float64) *LunchMenuItemUpdateOne {
	if f != nil {
		lmiuo.SetEmission(*f)
	}
	return lmiuo
}

// AddEmission adds f to the "emission" field.
func (lmiuo *LunchMenuItemUpdateOne) AddEmission(f float64) *LunchMenuItemUpdateOne {
	lmiuo.mutation.AddEmission(f)
	return lmiuo
}

// ClearEmission clears the value of the "emission" field.
func (lmiuo *LunchMenuItemUpdateOne) ClearEmission() *LunchMenuItemUpdateOne {
	lmiuo.mutation.ClearEmission()
	return lmiuo
}

// SetPrice sets the "price" field.
func (lmiuo *LunchMenuItemUpdateOne) SetPrice(s string) *LunchMenuItemUpdateOne {
	lmiuo.mutation.SetPrice(s)
	return lmiuo
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (lmiuo *LunchMenuItemUpdateOne) SetNillablePrice(s *string) *LunchMenuItemUpdateOne {
	if s != nil {
		lmiuo.SetPrice(*s)
	}
	return lmiuo
}

// ClearPrice clears the value of the "price" field.
func (lmiuo *LunchMenuItemUpdateOne) ClearPrice() *LunchMenuItemUpdateOne {
	lmiuo.mutation.ClearPrice()
	return lmiuo
}

// SetMenuID sets the "menu" edge to the LunchMenu entity by ID.
func (lmiuo *LunchMenuItemUpdateOne) SetMenuID(id int) *LunchMenuItemUpdateOne {
	lmiuo.mutation.SetMenuID(id)
	return lmiuo
}

// SetMenu sets the "menu" edge to the LunchMenu entity.
func (lmiuo *LunchMenuItemUpdateOne) SetMenu(l *LunchMenu) *LunchMenuItemUpdateOne {
	return lmiuo.SetMenuID(l.ID)
}

// Mutation returns the LunchMenuItemMutation object of the builder.
func (lmiuo *LunchMenuItemUpdateOne) Mutation() *LunchMenuItemMutation {
	return lmiuo.mutation
}

// ClearMenu clears the "menu" edge to the LunchMenu entity.
func (lmiuo *LunchMenuItemUpdateOne) ClearMenu() *LunchMenuItemUpdateOne {
	lmiuo.mutation.ClearMenu()
	return lmiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lmiuo *LunchMenuItemUpdateOne) Select(field string, fields ...string) *LunchMenuItemUpdateOne {
	lmiuo.fields = append([]string{field}, fields...)
	return lmiuo
}

// Save executes the query and returns the updated LunchMenuItem entity.
func (lmiuo *LunchMenuItemUpdateOne) Save(ctx context.Context) (*LunchMenuItem, error) {
	var (
		err  error
		node *LunchMenuItem
	)
	if len(lmiuo.hooks) == 0 {
		if err = lmiuo.check(); err != nil {
			return nil, err
		}
		node, err = lmiuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LunchMenuItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lmiuo.check(); err != nil {
				return nil, err
			}
			lmiuo.mutation = mutation
			node, err = lmiuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(lmiuo.hooks) - 1; i >= 0; i-- {
			if lmiuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lmiuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, lmiuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*LunchMenuItem)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from LunchMenuItemMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (lmiuo *LunchMenuItemUpdateOne) SaveX(ctx context.Context) *LunchMenuItem {
	node, err := lmiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lmiuo *LunchMenuItemUpdateOne) Exec(ctx context.Context) error {
	_, err := lmiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lmiuo *LunchMenuItemUpdateOne) ExecX(ctx context.Context) {
	if err := lmiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lmiuo *LunchMenuItemUpdateOne) check() error {
	if v, ok := lmiuo.mutation.Language(); ok {
		if err := lunchmenuitem.LanguageValidator(v); err != nil {
			return &ValidationError{Name: "language", err: fmt.Errorf(`ent: validator failed for field "LunchMenuItem.language": %w`, err)}
		}
	}
	if _, ok := lmiuo.mutation.MenuID(); lmiuo.mutation.MenuCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "LunchMenuItem.menu"`)
	}
	return nil
}

func (lmiuo *LunchMenuItemUpdateOne) sqlSave(ctx context.Context) (_node *LunchMenuItem, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lunchmenuitem.Table,
			Columns: lunchmenuitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lunchmenuitem.FieldID,
			},
		},
	}
	id, ok := lmiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LunchMenuItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lmiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lunchmenuitem.FieldID)
		for _, f := range fields {
			if !lunchmenuitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != lunchmenuitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lmiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lmiuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lunchmenuitem.FieldTitle,
		})
	}
	if lmiuo.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: lunchmenuitem.FieldTitle,
		})
	}
	if value, ok := lmiuo.mutation.Body(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lunchmenuitem.FieldBody,
		})
	}
	if value, ok := lmiuo.mutation.Language(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: lunchmenuitem.FieldLanguage,
		})
	}
	if lmiuo.mutation.LanguageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: lunchmenuitem.FieldLanguage,
		})
	}
	if value, ok := lmiuo.mutation.Preformatted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: lunchmenuitem.FieldPreformatted,
		})
	}
	if value, ok := lmiuo.mutation.Allergens(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: lunchmenuitem.FieldAllergens,
		})
	}
	if lmiuo.mutation.AllergensCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: lunchmenuitem.FieldAllergens,
		})
	}
	if value, ok := lmiuo.mutation.Emission(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: lunchmenuitem.FieldEmission,
		})
	}
	if value, ok := lmiuo.mutation.AddedEmission(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: lunchmenuitem.FieldEmission,
		})
	}
	if lmiuo.mutation.EmissionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: lunchmenuitem.FieldEmission,
		})
	}
	if value, ok := lmiuo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lunchmenuitem.FieldPrice,
		})
	}
	if lmiuo.mutation.PriceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: lunchmenuitem.FieldPrice,
		})
	}
	if lmiuo.mutation.MenuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lunchmenuitem.MenuTable,
			Columns: []string{lunchmenuitem.MenuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lunchmenu.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lmiuo.mutation.MenuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lunchmenuitem.MenuTable,
			Columns: []string{lunchmenuitem.MenuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lunchmenu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &LunchMenuItem{config: lmiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lmiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lunchmenuitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
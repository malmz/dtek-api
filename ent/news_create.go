// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dtekcth/dtek-api/ent/news"
)

// NewsCreate is the builder for creating a News entity.
type NewsCreate struct {
	config
	mutation *NewsMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (nc *NewsCreate) SetCreateTime(t time.Time) *NewsCreate {
	nc.mutation.SetCreateTime(t)
	return nc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (nc *NewsCreate) SetNillableCreateTime(t *time.Time) *NewsCreate {
	if t != nil {
		nc.SetCreateTime(*t)
	}
	return nc
}

// SetUpdateTime sets the "update_time" field.
func (nc *NewsCreate) SetUpdateTime(t time.Time) *NewsCreate {
	nc.mutation.SetUpdateTime(t)
	return nc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (nc *NewsCreate) SetNillableUpdateTime(t *time.Time) *NewsCreate {
	if t != nil {
		nc.SetUpdateTime(*t)
	}
	return nc
}

// SetTitle sets the "title" field.
func (nc *NewsCreate) SetTitle(s string) *NewsCreate {
	nc.mutation.SetTitle(s)
	return nc
}

// SetContent sets the "content" field.
func (nc *NewsCreate) SetContent(s string) *NewsCreate {
	nc.mutation.SetContent(s)
	return nc
}

// Mutation returns the NewsMutation object of the builder.
func (nc *NewsCreate) Mutation() *NewsMutation {
	return nc.mutation
}

// Save creates the News in the database.
func (nc *NewsCreate) Save(ctx context.Context) (*News, error) {
	var (
		err  error
		node *News
	)
	nc.defaults()
	if len(nc.hooks) == 0 {
		if err = nc.check(); err != nil {
			return nil, err
		}
		node, err = nc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NewsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nc.check(); err != nil {
				return nil, err
			}
			nc.mutation = mutation
			if node, err = nc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(nc.hooks) - 1; i >= 0; i-- {
			if nc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, nc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*News)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from NewsMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NewsCreate) SaveX(ctx context.Context) *News {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NewsCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NewsCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nc *NewsCreate) defaults() {
	if _, ok := nc.mutation.CreateTime(); !ok {
		v := news.DefaultCreateTime()
		nc.mutation.SetCreateTime(v)
	}
	if _, ok := nc.mutation.UpdateTime(); !ok {
		v := news.DefaultUpdateTime()
		nc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NewsCreate) check() error {
	if _, ok := nc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "News.create_time"`)}
	}
	if _, ok := nc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "News.update_time"`)}
	}
	if _, ok := nc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "News.title"`)}
	}
	if _, ok := nc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "News.content"`)}
	}
	return nil
}

func (nc *NewsCreate) sqlSave(ctx context.Context) (*News, error) {
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (nc *NewsCreate) createSpec() (*News, *sqlgraph.CreateSpec) {
	var (
		_node = &News{config: nc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: news.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: news.FieldID,
			},
		}
	)
	if value, ok := nc.mutation.CreateTime(); ok {
		_spec.SetField(news.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := nc.mutation.UpdateTime(); ok {
		_spec.SetField(news.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := nc.mutation.Title(); ok {
		_spec.SetField(news.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := nc.mutation.Content(); ok {
		_spec.SetField(news.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	return _node, _spec
}

// NewsCreateBulk is the builder for creating many News entities in bulk.
type NewsCreateBulk struct {
	config
	builders []*NewsCreate
}

// Save creates the News entities in the database.
func (ncb *NewsCreateBulk) Save(ctx context.Context) ([]*News, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*News, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NewsMutation)
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
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NewsCreateBulk) SaveX(ctx context.Context) []*News {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NewsCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NewsCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}

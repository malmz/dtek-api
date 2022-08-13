package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// News holds the schema definition for the News entity.
type News struct {
	ent.Schema
}

// Fields of the News.
func (News) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("content"),
	}
}

// Edges of the News.
func (News) Edges() []ent.Edge {
	return nil
}

func (News) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

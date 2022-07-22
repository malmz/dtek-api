package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// LunchMenu holds the schema definition for the LunchMenu entity.
type LunchMenu struct {
	ent.Schema
}

// Fields of the LunchMenu.
func (LunchMenu) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date"),
	}
}

// Edges of the LunchMenu.
func (LunchMenu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("items", LunchMenuItem.Type),
		edge.From("resturant", Resturant.Type).
			Ref("menu").
			Unique().
			Required(),
	}
}

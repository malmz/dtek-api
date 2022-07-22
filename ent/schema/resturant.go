package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Resturant holds the schema definition for the Resturant entity.
type Resturant struct {
	ent.Schema
}

// Fields of the Resturant.
func (Resturant) Fields() []ent.Field {
	return []ent.Field{
		field.String("slug").Unique(),
		field.String("name"),
		field.Enum("campus").Values("lindholmen", "johanneberg"),
	}
}

// Edges of the Resturant.
func (Resturant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("menu", LunchMenu.Type),
	}
}

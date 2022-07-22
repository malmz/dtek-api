package schema

import (
	"github.com/dtekcth/dtek-api/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// LunchMenuItem holds the schema definition for the LunchMenuItem entity.
type LunchMenuItem struct {
	ent.Schema
}

// Fields of the LunchMenuItem.
func (LunchMenuItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Optional(),
		field.String("body"),
		field.Enum("language").Values("se", "en").Optional(),
		field.Bool("preformatted").Default(false),
		field.JSON("allergens", []model.Allergen{}).Optional(),
		field.Float("emission").Optional(),
		field.String("price").Optional(),
	}
}

// Edges of the LunchMenuItem.
func (LunchMenuItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("menu", LunchMenu.Type).
			Ref("items").
			Unique().
			Required(),
	}
}

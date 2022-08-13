package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/dtekcth/dtek-api/model"
)

// LunchMenu holds the schema definition for the LunchMenu entity.
type LunchMenu struct {
	ent.Schema
}

// Fields of the LunchMenu.
func (LunchMenu) Fields() []ent.Field {
	return []ent.Field{
		field.String("resturant"),
		field.Time("date"),
		field.Enum("language").Values("se", "en").Optional(),
		field.String("name"),
		field.JSON("menu", []model.LunchMenuItem{}),
	}
}

// Edges of the LunchMenu.
func (LunchMenu) Edges() []ent.Edge {
	return nil
}

func (LunchMenu) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("resturant", "date", "language").Unique(),
	}
}

func (LunchMenu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.UpdateTime{},
	}
}

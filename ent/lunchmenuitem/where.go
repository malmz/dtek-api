// Code generated by ent, DO NOT EDIT.

package lunchmenuitem

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/dtekcth/dtek-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Body applies equality check predicate on the "body" field. It's identical to BodyEQ.
func Body(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBody), v))
	})
}

// Preformatted applies equality check predicate on the "preformatted" field. It's identical to PreformattedEQ.
func Preformatted(v bool) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPreformatted), v))
	})
}

// Emission applies equality check predicate on the "emission" field. It's identical to EmissionEQ.
func Emission(v float64) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmission), v))
	})
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.LunchMenuItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.LunchMenuItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleIsNil applies the IsNil predicate on the "title" field.
func TitleIsNil() predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTitle)))
	})
}

// TitleNotNil applies the NotNil predicate on the "title" field.
func TitleNotNil() predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTitle)))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// BodyEQ applies the EQ predicate on the "body" field.
func BodyEQ(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBody), v))
	})
}

// BodyNEQ applies the NEQ predicate on the "body" field.
func BodyNEQ(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBody), v))
	})
}

// BodyIn applies the In predicate on the "body" field.
func BodyIn(vs ...string) predicate.LunchMenuItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBody), v...))
	})
}

// BodyNotIn applies the NotIn predicate on the "body" field.
func BodyNotIn(vs ...string) predicate.LunchMenuItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBody), v...))
	})
}

// BodyGT applies the GT predicate on the "body" field.
func BodyGT(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBody), v))
	})
}

// BodyGTE applies the GTE predicate on the "body" field.
func BodyGTE(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBody), v))
	})
}

// BodyLT applies the LT predicate on the "body" field.
func BodyLT(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBody), v))
	})
}

// BodyLTE applies the LTE predicate on the "body" field.
func BodyLTE(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBody), v))
	})
}

// BodyContains applies the Contains predicate on the "body" field.
func BodyContains(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBody), v))
	})
}

// BodyHasPrefix applies the HasPrefix predicate on the "body" field.
func BodyHasPrefix(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBody), v))
	})
}

// BodyHasSuffix applies the HasSuffix predicate on the "body" field.
func BodyHasSuffix(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBody), v))
	})
}

// BodyEqualFold applies the EqualFold predicate on the "body" field.
func BodyEqualFold(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBody), v))
	})
}

// BodyContainsFold applies the ContainsFold predicate on the "body" field.
func BodyContainsFold(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBody), v))
	})
}

// LanguageEQ applies the EQ predicate on the "language" field.
func LanguageEQ(v Language) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLanguage), v))
	})
}

// LanguageNEQ applies the NEQ predicate on the "language" field.
func LanguageNEQ(v Language) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLanguage), v))
	})
}

// LanguageIn applies the In predicate on the "language" field.
func LanguageIn(vs ...Language) predicate.LunchMenuItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLanguage), v...))
	})
}

// LanguageNotIn applies the NotIn predicate on the "language" field.
func LanguageNotIn(vs ...Language) predicate.LunchMenuItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLanguage), v...))
	})
}

// LanguageIsNil applies the IsNil predicate on the "language" field.
func LanguageIsNil() predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLanguage)))
	})
}

// LanguageNotNil applies the NotNil predicate on the "language" field.
func LanguageNotNil() predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLanguage)))
	})
}

// PreformattedEQ applies the EQ predicate on the "preformatted" field.
func PreformattedEQ(v bool) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPreformatted), v))
	})
}

// PreformattedNEQ applies the NEQ predicate on the "preformatted" field.
func PreformattedNEQ(v bool) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPreformatted), v))
	})
}

// AllergensIsNil applies the IsNil predicate on the "allergens" field.
func AllergensIsNil() predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAllergens)))
	})
}

// AllergensNotNil applies the NotNil predicate on the "allergens" field.
func AllergensNotNil() predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAllergens)))
	})
}

// EmissionEQ applies the EQ predicate on the "emission" field.
func EmissionEQ(v float64) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmission), v))
	})
}

// EmissionNEQ applies the NEQ predicate on the "emission" field.
func EmissionNEQ(v float64) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmission), v))
	})
}

// EmissionIn applies the In predicate on the "emission" field.
func EmissionIn(vs ...float64) predicate.LunchMenuItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmission), v...))
	})
}

// EmissionNotIn applies the NotIn predicate on the "emission" field.
func EmissionNotIn(vs ...float64) predicate.LunchMenuItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmission), v...))
	})
}

// EmissionGT applies the GT predicate on the "emission" field.
func EmissionGT(v float64) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmission), v))
	})
}

// EmissionGTE applies the GTE predicate on the "emission" field.
func EmissionGTE(v float64) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmission), v))
	})
}

// EmissionLT applies the LT predicate on the "emission" field.
func EmissionLT(v float64) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmission), v))
	})
}

// EmissionLTE applies the LTE predicate on the "emission" field.
func EmissionLTE(v float64) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmission), v))
	})
}

// EmissionIsNil applies the IsNil predicate on the "emission" field.
func EmissionIsNil() predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEmission)))
	})
}

// EmissionNotNil applies the NotNil predicate on the "emission" field.
func EmissionNotNil() predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEmission)))
	})
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrice), v))
	})
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...string) predicate.LunchMenuItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrice), v...))
	})
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...string) predicate.LunchMenuItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrice), v...))
	})
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrice), v))
	})
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrice), v))
	})
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrice), v))
	})
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrice), v))
	})
}

// PriceContains applies the Contains predicate on the "price" field.
func PriceContains(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPrice), v))
	})
}

// PriceHasPrefix applies the HasPrefix predicate on the "price" field.
func PriceHasPrefix(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPrice), v))
	})
}

// PriceHasSuffix applies the HasSuffix predicate on the "price" field.
func PriceHasSuffix(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPrice), v))
	})
}

// PriceIsNil applies the IsNil predicate on the "price" field.
func PriceIsNil() predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPrice)))
	})
}

// PriceNotNil applies the NotNil predicate on the "price" field.
func PriceNotNil() predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPrice)))
	})
}

// PriceEqualFold applies the EqualFold predicate on the "price" field.
func PriceEqualFold(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPrice), v))
	})
}

// PriceContainsFold applies the ContainsFold predicate on the "price" field.
func PriceContainsFold(v string) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPrice), v))
	})
}

// HasMenu applies the HasEdge predicate on the "menu" edge.
func HasMenu() predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MenuTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MenuTable, MenuColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMenuWith applies the HasEdge predicate on the "menu" edge with a given conditions (other predicates).
func HasMenuWith(preds ...predicate.LunchMenu) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MenuInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MenuTable, MenuColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LunchMenuItem) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LunchMenuItem) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.LunchMenuItem) predicate.LunchMenuItem {
	return predicate.LunchMenuItem(func(s *sql.Selector) {
		p(s.Not())
	})
}
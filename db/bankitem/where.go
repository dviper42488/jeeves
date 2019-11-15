// Code generated by entc, DO NOT EDIT.

package bankitem

import (
	"github.com/AlecAivazis/jeeves/db/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldID), id))
		},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(ids) == 0 {
				s.Where(sql.False())
				return
			}
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i] = ids[i]
			}
			s.Where(sql.In(s.C(FieldID), v...))
		},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(ids) == 0 {
				s.Where(sql.False())
				return
			}
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i] = ids[i]
			}
			s.Where(sql.NotIn(s.C(FieldID), v...))
		},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldID), id))
		},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldID), id))
		},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldID), id))
		},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldID), id))
		},
	)
}

// ItemID applies equality check predicate on the "itemID" field. It's identical to ItemIDEQ.
func ItemID(v string) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldItemID), v))
		},
	)
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v int) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldQuantity), v))
		},
	)
}

// ItemIDEQ applies the EQ predicate on the "itemID" field.
func ItemIDEQ(v string) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldItemID), v))
		},
	)
}

// ItemIDNEQ applies the NEQ predicate on the "itemID" field.
func ItemIDNEQ(v string) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldItemID), v))
		},
	)
}

// ItemIDIn applies the In predicate on the "itemID" field.
func ItemIDIn(vs ...string) predicate.BankItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BankItem(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldItemID), v...))
		},
	)
}

// ItemIDNotIn applies the NotIn predicate on the "itemID" field.
func ItemIDNotIn(vs ...string) predicate.BankItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BankItem(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldItemID), v...))
		},
	)
}

// ItemIDGT applies the GT predicate on the "itemID" field.
func ItemIDGT(v string) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldItemID), v))
		},
	)
}

// ItemIDGTE applies the GTE predicate on the "itemID" field.
func ItemIDGTE(v string) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldItemID), v))
		},
	)
}

// ItemIDLT applies the LT predicate on the "itemID" field.
func ItemIDLT(v string) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldItemID), v))
		},
	)
}

// ItemIDLTE applies the LTE predicate on the "itemID" field.
func ItemIDLTE(v string) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldItemID), v))
		},
	)
}

// ItemIDContains applies the Contains predicate on the "itemID" field.
func ItemIDContains(v string) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldItemID), v))
		},
	)
}

// ItemIDHasPrefix applies the HasPrefix predicate on the "itemID" field.
func ItemIDHasPrefix(v string) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldItemID), v))
		},
	)
}

// ItemIDHasSuffix applies the HasSuffix predicate on the "itemID" field.
func ItemIDHasSuffix(v string) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldItemID), v))
		},
	)
}

// ItemIDEqualFold applies the EqualFold predicate on the "itemID" field.
func ItemIDEqualFold(v string) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.EqualFold(s.C(FieldItemID), v))
		},
	)
}

// ItemIDContainsFold applies the ContainsFold predicate on the "itemID" field.
func ItemIDContainsFold(v string) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.ContainsFold(s.C(FieldItemID), v))
		},
	)
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v int) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldQuantity), v))
		},
	)
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v int) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldQuantity), v))
		},
	)
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...int) predicate.BankItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BankItem(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldQuantity), v...))
		},
	)
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...int) predicate.BankItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BankItem(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldQuantity), v...))
		},
	)
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v int) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldQuantity), v))
		},
	)
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v int) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldQuantity), v))
		},
	)
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v int) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldQuantity), v))
		},
	)
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v int) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldQuantity), v))
		},
	)
}

// HasBank applies the HasEdge predicate on the "bank" edge.
func HasBank() predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(sql.NotNull(t1.C(BankColumn)))
		},
	)
}

// HasBankWith applies the HasEdge predicate on the "bank" edge with a given conditions (other predicates).
func HasBankWith(preds ...predicate.GuildBank) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			builder := sql.Dialect(s.Dialect())
			t1 := s.Table()
			t2 := builder.Select(FieldID).From(builder.Table(BankInverseTable))
			for _, p := range preds {
				p(t2)
			}
			s.Where(sql.In(t1.C(BankColumn), t2))
		},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.BankItem) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			for _, p := range predicates {
				p(s)
			}
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.BankItem) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			for i, p := range predicates {
				if i > 0 {
					s.Or()
				}
				p(s)
			}
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.BankItem) predicate.BankItem {
	return predicate.BankItem(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}

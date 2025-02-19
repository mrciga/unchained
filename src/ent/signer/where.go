// Code generated by ent, DO NOT EDIT.

package signer

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/KenshiTech/unchained/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Signer {
	return predicate.Signer(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Signer {
	return predicate.Signer(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Signer {
	return predicate.Signer(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Signer {
	return predicate.Signer(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Signer {
	return predicate.Signer(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Signer {
	return predicate.Signer(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Signer {
	return predicate.Signer(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Signer {
	return predicate.Signer(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Signer {
	return predicate.Signer(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Signer {
	return predicate.Signer(sql.FieldEQ(FieldName, v))
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v []byte) predicate.Signer {
	return predicate.Signer(sql.FieldEQ(FieldKey, v))
}

// Points applies equality check predicate on the "points" field. It's identical to PointsEQ.
func Points(v int64) predicate.Signer {
	return predicate.Signer(sql.FieldEQ(FieldPoints, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Signer {
	return predicate.Signer(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Signer {
	return predicate.Signer(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Signer {
	return predicate.Signer(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Signer {
	return predicate.Signer(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Signer {
	return predicate.Signer(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Signer {
	return predicate.Signer(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Signer {
	return predicate.Signer(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Signer {
	return predicate.Signer(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Signer {
	return predicate.Signer(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Signer {
	return predicate.Signer(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Signer {
	return predicate.Signer(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Signer {
	return predicate.Signer(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Signer {
	return predicate.Signer(sql.FieldContainsFold(FieldName, v))
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v []byte) predicate.Signer {
	return predicate.Signer(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v []byte) predicate.Signer {
	return predicate.Signer(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...[]byte) predicate.Signer {
	return predicate.Signer(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...[]byte) predicate.Signer {
	return predicate.Signer(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v []byte) predicate.Signer {
	return predicate.Signer(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v []byte) predicate.Signer {
	return predicate.Signer(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v []byte) predicate.Signer {
	return predicate.Signer(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v []byte) predicate.Signer {
	return predicate.Signer(sql.FieldLTE(FieldKey, v))
}

// PointsEQ applies the EQ predicate on the "points" field.
func PointsEQ(v int64) predicate.Signer {
	return predicate.Signer(sql.FieldEQ(FieldPoints, v))
}

// PointsNEQ applies the NEQ predicate on the "points" field.
func PointsNEQ(v int64) predicate.Signer {
	return predicate.Signer(sql.FieldNEQ(FieldPoints, v))
}

// PointsIn applies the In predicate on the "points" field.
func PointsIn(vs ...int64) predicate.Signer {
	return predicate.Signer(sql.FieldIn(FieldPoints, vs...))
}

// PointsNotIn applies the NotIn predicate on the "points" field.
func PointsNotIn(vs ...int64) predicate.Signer {
	return predicate.Signer(sql.FieldNotIn(FieldPoints, vs...))
}

// PointsGT applies the GT predicate on the "points" field.
func PointsGT(v int64) predicate.Signer {
	return predicate.Signer(sql.FieldGT(FieldPoints, v))
}

// PointsGTE applies the GTE predicate on the "points" field.
func PointsGTE(v int64) predicate.Signer {
	return predicate.Signer(sql.FieldGTE(FieldPoints, v))
}

// PointsLT applies the LT predicate on the "points" field.
func PointsLT(v int64) predicate.Signer {
	return predicate.Signer(sql.FieldLT(FieldPoints, v))
}

// PointsLTE applies the LTE predicate on the "points" field.
func PointsLTE(v int64) predicate.Signer {
	return predicate.Signer(sql.FieldLTE(FieldPoints, v))
}

// HasAssetPrice applies the HasEdge predicate on the "assetPrice" edge.
func HasAssetPrice() predicate.Signer {
	return predicate.Signer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, AssetPriceTable, AssetPricePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAssetPriceWith applies the HasEdge predicate on the "assetPrice" edge with a given conditions (other predicates).
func HasAssetPriceWith(preds ...predicate.AssetPrice) predicate.Signer {
	return predicate.Signer(func(s *sql.Selector) {
		step := newAssetPriceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Signer) predicate.Signer {
	return predicate.Signer(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Signer) predicate.Signer {
	return predicate.Signer(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Signer) predicate.Signer {
	return predicate.Signer(sql.NotPredicates(p))
}

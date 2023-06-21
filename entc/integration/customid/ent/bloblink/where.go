// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package bloblink

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"github.com/google/uuid"
)

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.BlobLink {
	return predicate.BlobLink(sql.FieldEQ(FieldCreatedAt, v))
}

// BlobID applies equality check predicate on the "blob_id" field. It's identical to BlobIDEQ.
func BlobID(v uuid.UUID) predicate.BlobLink {
	return predicate.BlobLink(sql.FieldEQ(FieldBlobID, v))
}

// LinkID applies equality check predicate on the "link_id" field. It's identical to LinkIDEQ.
func LinkID(v uuid.UUID) predicate.BlobLink {
	return predicate.BlobLink(sql.FieldEQ(FieldLinkID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.BlobLink {
	return predicate.BlobLink(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.BlobLink {
	return predicate.BlobLink(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.BlobLink {
	return predicate.BlobLink(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.BlobLink {
	return predicate.BlobLink(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.BlobLink {
	return predicate.BlobLink(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.BlobLink {
	return predicate.BlobLink(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.BlobLink {
	return predicate.BlobLink(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.BlobLink {
	return predicate.BlobLink(sql.FieldLTE(FieldCreatedAt, v))
}

// BlobIDEQ applies the EQ predicate on the "blob_id" field.
func BlobIDEQ(v uuid.UUID) predicate.BlobLink {
	return predicate.BlobLink(sql.FieldEQ(FieldBlobID, v))
}

// BlobIDNEQ applies the NEQ predicate on the "blob_id" field.
func BlobIDNEQ(v uuid.UUID) predicate.BlobLink {
	return predicate.BlobLink(sql.FieldNEQ(FieldBlobID, v))
}

// BlobIDIn applies the In predicate on the "blob_id" field.
func BlobIDIn(vs ...uuid.UUID) predicate.BlobLink {
	return predicate.BlobLink(sql.FieldIn(FieldBlobID, vs...))
}

// BlobIDNotIn applies the NotIn predicate on the "blob_id" field.
func BlobIDNotIn(vs ...uuid.UUID) predicate.BlobLink {
	return predicate.BlobLink(sql.FieldNotIn(FieldBlobID, vs...))
}

// LinkIDEQ applies the EQ predicate on the "link_id" field.
func LinkIDEQ(v uuid.UUID) predicate.BlobLink {
	return predicate.BlobLink(sql.FieldEQ(FieldLinkID, v))
}

// LinkIDNEQ applies the NEQ predicate on the "link_id" field.
func LinkIDNEQ(v uuid.UUID) predicate.BlobLink {
	return predicate.BlobLink(sql.FieldNEQ(FieldLinkID, v))
}

// LinkIDIn applies the In predicate on the "link_id" field.
func LinkIDIn(vs ...uuid.UUID) predicate.BlobLink {
	return predicate.BlobLink(sql.FieldIn(FieldLinkID, vs...))
}

// LinkIDNotIn applies the NotIn predicate on the "link_id" field.
func LinkIDNotIn(vs ...uuid.UUID) predicate.BlobLink {
	return predicate.BlobLink(sql.FieldNotIn(FieldLinkID, vs...))
}

// HasBlob applies the HasEdge predicate on the "blob" edge.
func HasBlob() predicate.BlobLink {
	return predicate.BlobLink(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, BlobColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, BlobTable, BlobColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBlobWith applies the HasEdge predicate on the "blob" edge with a given conditions (other predicates).
func HasBlobWith(preds ...predicate.Blob) predicate.BlobLink {
	return predicate.BlobLink(func(s *sql.Selector) {
		step := newBlobStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLink applies the HasEdge predicate on the "link" edge.
func HasLink() predicate.BlobLink {
	return predicate.BlobLink(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, LinkColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, LinkTable, LinkColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLinkWith applies the HasEdge predicate on the "link" edge with a given conditions (other predicates).
func HasLinkWith(preds ...predicate.Blob) predicate.BlobLink {
	return predicate.BlobLink(func(s *sql.Selector) {
		step := newLinkStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BlobLink) predicate.BlobLink {
	return predicate.BlobLink(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BlobLink) predicate.BlobLink {
	return predicate.BlobLink(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.BlobLink) predicate.BlobLink {
	return predicate.BlobLink(sql.NotPredicates(p))
}

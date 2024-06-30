// Code generated by ent, DO NOT EDIT.

package tags

import (
	"backend/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Tags {
	return predicate.Tags(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Tags {
	return predicate.Tags(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Tags {
	return predicate.Tags(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Tags {
	return predicate.Tags(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Tags {
	return predicate.Tags(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Tags {
	return predicate.Tags(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Tags {
	return predicate.Tags(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Tags {
	return predicate.Tags(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Tags {
	return predicate.Tags(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Tags {
	return predicate.Tags(sql.FieldEQ(FieldName, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uint64) predicate.Tags {
	return predicate.Tags(sql.FieldEQ(FieldUserID, v))
}

// ArticleID applies equality check predicate on the "article_id" field. It's identical to ArticleIDEQ.
func ArticleID(v uint64) predicate.Tags {
	return predicate.Tags(sql.FieldEQ(FieldArticleID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Tags {
	return predicate.Tags(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Tags {
	return predicate.Tags(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Tags {
	return predicate.Tags(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Tags {
	return predicate.Tags(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Tags {
	return predicate.Tags(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Tags {
	return predicate.Tags(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Tags {
	return predicate.Tags(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Tags {
	return predicate.Tags(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Tags {
	return predicate.Tags(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Tags {
	return predicate.Tags(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Tags {
	return predicate.Tags(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Tags {
	return predicate.Tags(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Tags {
	return predicate.Tags(sql.FieldContainsFold(FieldName, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uint64) predicate.Tags {
	return predicate.Tags(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uint64) predicate.Tags {
	return predicate.Tags(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uint64) predicate.Tags {
	return predicate.Tags(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uint64) predicate.Tags {
	return predicate.Tags(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uint64) predicate.Tags {
	return predicate.Tags(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uint64) predicate.Tags {
	return predicate.Tags(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uint64) predicate.Tags {
	return predicate.Tags(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uint64) predicate.Tags {
	return predicate.Tags(sql.FieldLTE(FieldUserID, v))
}

// ArticleIDEQ applies the EQ predicate on the "article_id" field.
func ArticleIDEQ(v uint64) predicate.Tags {
	return predicate.Tags(sql.FieldEQ(FieldArticleID, v))
}

// ArticleIDNEQ applies the NEQ predicate on the "article_id" field.
func ArticleIDNEQ(v uint64) predicate.Tags {
	return predicate.Tags(sql.FieldNEQ(FieldArticleID, v))
}

// ArticleIDIn applies the In predicate on the "article_id" field.
func ArticleIDIn(vs ...uint64) predicate.Tags {
	return predicate.Tags(sql.FieldIn(FieldArticleID, vs...))
}

// ArticleIDNotIn applies the NotIn predicate on the "article_id" field.
func ArticleIDNotIn(vs ...uint64) predicate.Tags {
	return predicate.Tags(sql.FieldNotIn(FieldArticleID, vs...))
}

// ArticleIDGT applies the GT predicate on the "article_id" field.
func ArticleIDGT(v uint64) predicate.Tags {
	return predicate.Tags(sql.FieldGT(FieldArticleID, v))
}

// ArticleIDGTE applies the GTE predicate on the "article_id" field.
func ArticleIDGTE(v uint64) predicate.Tags {
	return predicate.Tags(sql.FieldGTE(FieldArticleID, v))
}

// ArticleIDLT applies the LT predicate on the "article_id" field.
func ArticleIDLT(v uint64) predicate.Tags {
	return predicate.Tags(sql.FieldLT(FieldArticleID, v))
}

// ArticleIDLTE applies the LTE predicate on the "article_id" field.
func ArticleIDLTE(v uint64) predicate.Tags {
	return predicate.Tags(sql.FieldLTE(FieldArticleID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Tags) predicate.Tags {
	return predicate.Tags(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Tags) predicate.Tags {
	return predicate.Tags(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Tags) predicate.Tags {
	return predicate.Tags(sql.NotPredicates(p))
}
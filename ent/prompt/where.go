// Code generated by ent, DO NOT EDIT.

package prompt

import (
	"skillsdemo/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Prompt {
	return predicate.Prompt(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Prompt {
	return predicate.Prompt(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Prompt {
	return predicate.Prompt(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Prompt {
	return predicate.Prompt(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Prompt {
	return predicate.Prompt(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Prompt {
	return predicate.Prompt(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Prompt {
	return predicate.Prompt(sql.FieldLTE(FieldID, id))
}

// SortOrder applies equality check predicate on the "sort_order" field. It's identical to SortOrderEQ.
func SortOrder(v int) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldSortOrder, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldDescription, v))
}

// AdditionalFeedback applies equality check predicate on the "additional_feedback" field. It's identical to AdditionalFeedbackEQ.
func AdditionalFeedback(v bool) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldAdditionalFeedback, v))
}

// SortOrderEQ applies the EQ predicate on the "sort_order" field.
func SortOrderEQ(v int) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldSortOrder, v))
}

// SortOrderNEQ applies the NEQ predicate on the "sort_order" field.
func SortOrderNEQ(v int) predicate.Prompt {
	return predicate.Prompt(sql.FieldNEQ(FieldSortOrder, v))
}

// SortOrderIn applies the In predicate on the "sort_order" field.
func SortOrderIn(vs ...int) predicate.Prompt {
	return predicate.Prompt(sql.FieldIn(FieldSortOrder, vs...))
}

// SortOrderNotIn applies the NotIn predicate on the "sort_order" field.
func SortOrderNotIn(vs ...int) predicate.Prompt {
	return predicate.Prompt(sql.FieldNotIn(FieldSortOrder, vs...))
}

// SortOrderGT applies the GT predicate on the "sort_order" field.
func SortOrderGT(v int) predicate.Prompt {
	return predicate.Prompt(sql.FieldGT(FieldSortOrder, v))
}

// SortOrderGTE applies the GTE predicate on the "sort_order" field.
func SortOrderGTE(v int) predicate.Prompt {
	return predicate.Prompt(sql.FieldGTE(FieldSortOrder, v))
}

// SortOrderLT applies the LT predicate on the "sort_order" field.
func SortOrderLT(v int) predicate.Prompt {
	return predicate.Prompt(sql.FieldLT(FieldSortOrder, v))
}

// SortOrderLTE applies the LTE predicate on the "sort_order" field.
func SortOrderLTE(v int) predicate.Prompt {
	return predicate.Prompt(sql.FieldLTE(FieldSortOrder, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Prompt {
	return predicate.Prompt(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Prompt {
	return predicate.Prompt(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Prompt {
	return predicate.Prompt(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Prompt {
	return predicate.Prompt(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Prompt {
	return predicate.Prompt(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Prompt {
	return predicate.Prompt(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Prompt {
	return predicate.Prompt(sql.FieldContainsFold(FieldDescription, v))
}

// AdditionalFeedbackEQ applies the EQ predicate on the "additional_feedback" field.
func AdditionalFeedbackEQ(v bool) predicate.Prompt {
	return predicate.Prompt(sql.FieldEQ(FieldAdditionalFeedback, v))
}

// AdditionalFeedbackNEQ applies the NEQ predicate on the "additional_feedback" field.
func AdditionalFeedbackNEQ(v bool) predicate.Prompt {
	return predicate.Prompt(sql.FieldNEQ(FieldAdditionalFeedback, v))
}

// HasSurvey applies the HasEdge predicate on the "survey" edge.
func HasSurvey() predicate.Prompt {
	return predicate.Prompt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SurveyTable, SurveyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSurveyWith applies the HasEdge predicate on the "survey" edge with a given conditions (other predicates).
func HasSurveyWith(preds ...predicate.Survey) predicate.Prompt {
	return predicate.Prompt(func(s *sql.Selector) {
		step := newSurveyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Prompt) predicate.Prompt {
	return predicate.Prompt(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Prompt) predicate.Prompt {
	return predicate.Prompt(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Prompt) predicate.Prompt {
	return predicate.Prompt(sql.NotPredicates(p))
}

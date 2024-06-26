// Code generated by ent, DO NOT EDIT.

package prompt

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the prompt type in the database.
	Label = "prompt"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSortOrder holds the string denoting the sort_order field in the database.
	FieldSortOrder = "sort_order"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldResponseType holds the string denoting the response_type field in the database.
	FieldResponseType = "response_type"
	// FieldAdditionalFeedback holds the string denoting the additional_feedback field in the database.
	FieldAdditionalFeedback = "additional_feedback"
	// EdgeSurvey holds the string denoting the survey edge name in mutations.
	EdgeSurvey = "survey"
	// Table holds the table name of the prompt in the database.
	Table = "prompts"
	// SurveyTable is the table that holds the survey relation/edge.
	SurveyTable = "prompts"
	// SurveyInverseTable is the table name for the Survey entity.
	// It exists in this package in order to avoid circular dependency with the "survey" package.
	SurveyInverseTable = "surveys"
	// SurveyColumn is the table column denoting the survey relation/edge.
	SurveyColumn = "survey_prompts"
)

// Columns holds all SQL columns for prompt fields.
var Columns = []string{
	FieldID,
	FieldSortOrder,
	FieldTitle,
	FieldDescription,
	FieldResponseType,
	FieldAdditionalFeedback,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "prompts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"survey_prompts",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// SortOrderValidator is a validator for the "sort_order" field. It is called by the builders before save.
	SortOrderValidator func(int) error
	// DefaultAdditionalFeedback holds the default value on creation for the "additional_feedback" field.
	DefaultAdditionalFeedback bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Prompt queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySortOrder orders the results by the sort_order field.
func BySortOrder(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSortOrder, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByAdditionalFeedback orders the results by the additional_feedback field.
func ByAdditionalFeedback(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAdditionalFeedback, opts...).ToFunc()
}

// BySurveyField orders the results by survey field.
func BySurveyField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSurveyStep(), sql.OrderByField(field, opts...))
	}
}
func newSurveyStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SurveyInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SurveyTable, SurveyColumn),
	)
}

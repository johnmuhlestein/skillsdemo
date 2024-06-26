// Code generated by ent, DO NOT EDIT.

package promptresponse

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the promptresponse type in the database.
	Label = "prompt_response"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldParsedTemplate holds the string denoting the parsed_template field in the database.
	FieldParsedTemplate = "parsed_template"
	// FieldPromptIndex holds the string denoting the prompt_index field in the database.
	FieldPromptIndex = "prompt_index"
	// FieldRangeValue holds the string denoting the range_value field in the database.
	FieldRangeValue = "range_value"
	// FieldBoolValue holds the string denoting the bool_value field in the database.
	FieldBoolValue = "bool_value"
	// FieldEnumValue holds the string denoting the enum_value field in the database.
	FieldEnumValue = "enum_value"
	// FieldLabelValues holds the string denoting the label_values field in the database.
	FieldLabelValues = "label_values"
	// FieldFreeformValue holds the string denoting the freeform_value field in the database.
	FieldFreeformValue = "freeform_value"
	// FieldAnsweredTime holds the string denoting the answered_time field in the database.
	FieldAnsweredTime = "answered_time"
	// EdgeFeedback holds the string denoting the feedback edge name in mutations.
	EdgeFeedback = "feedback"
	// Table holds the table name of the promptresponse in the database.
	Table = "prompt_responses"
	// FeedbackTable is the table that holds the feedback relation/edge.
	FeedbackTable = "prompt_responses"
	// FeedbackInverseTable is the table name for the Feedback entity.
	// It exists in this package in order to avoid circular dependency with the "feedback" package.
	FeedbackInverseTable = "feedbacks"
	// FeedbackColumn is the table column denoting the feedback relation/edge.
	FeedbackColumn = "feedback_responses"
)

// Columns holds all SQL columns for promptresponse fields.
var Columns = []string{
	FieldID,
	FieldParsedTemplate,
	FieldPromptIndex,
	FieldRangeValue,
	FieldBoolValue,
	FieldEnumValue,
	FieldLabelValues,
	FieldFreeformValue,
	FieldAnsweredTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "prompt_responses"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"feedback_responses",
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
	// PromptIndexValidator is a validator for the "prompt_index" field. It is called by the builders before save.
	PromptIndexValidator func(int) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the PromptResponse queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByParsedTemplate orders the results by the parsed_template field.
func ByParsedTemplate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParsedTemplate, opts...).ToFunc()
}

// ByPromptIndex orders the results by the prompt_index field.
func ByPromptIndex(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPromptIndex, opts...).ToFunc()
}

// ByRangeValue orders the results by the range_value field.
func ByRangeValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRangeValue, opts...).ToFunc()
}

// ByBoolValue orders the results by the bool_value field.
func ByBoolValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBoolValue, opts...).ToFunc()
}

// ByFreeformValue orders the results by the freeform_value field.
func ByFreeformValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFreeformValue, opts...).ToFunc()
}

// ByAnsweredTime orders the results by the answered_time field.
func ByAnsweredTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAnsweredTime, opts...).ToFunc()
}

// ByFeedbackField orders the results by feedback field.
func ByFeedbackField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFeedbackStep(), sql.OrderByField(field, opts...))
	}
}
func newFeedbackStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FeedbackInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, FeedbackTable, FeedbackColumn),
	)
}

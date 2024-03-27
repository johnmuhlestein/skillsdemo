// Code generated by ent, DO NOT EDIT.

package survey

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the survey type in the database.
	Label = "survey"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldActiveTime holds the string denoting the active_time field in the database.
	FieldActiveTime = "active_time"
	// FieldArchiveTime holds the string denoting the archive_time field in the database.
	FieldArchiveTime = "archive_time"
	// EdgePrompts holds the string denoting the prompts edge name in mutations.
	EdgePrompts = "prompts"
	// Table holds the table name of the survey in the database.
	Table = "surveys"
	// PromptsTable is the table that holds the prompts relation/edge.
	PromptsTable = "prompts"
	// PromptsInverseTable is the table name for the Prompt entity.
	// It exists in this package in order to avoid circular dependency with the "prompt" package.
	PromptsInverseTable = "prompts"
	// PromptsColumn is the table column denoting the prompts relation/edge.
	PromptsColumn = "survey_prompts"
)

// Columns holds all SQL columns for survey fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldDescription,
	FieldStatus,
	FieldActiveTime,
	FieldArchiveTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// StatusUnpublished is the default value of the Status enum.
const DefaultStatus = StatusUnpublished

// Status values.
const (
	StatusUnpublished Status = "unpublished"
	StatusActive      Status = "active"
	StatusArchived    Status = "archived"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusUnpublished, StatusActive, StatusArchived:
		return nil
	default:
		return fmt.Errorf("survey: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Survey queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByActiveTime orders the results by the active_time field.
func ByActiveTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActiveTime, opts...).ToFunc()
}

// ByArchiveTime orders the results by the archive_time field.
func ByArchiveTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldArchiveTime, opts...).ToFunc()
}

// ByPromptsCount orders the results by prompts count.
func ByPromptsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPromptsStep(), opts...)
	}
}

// ByPrompts orders the results by prompts terms.
func ByPrompts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPromptsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPromptsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PromptsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PromptsTable, PromptsColumn),
	)
}
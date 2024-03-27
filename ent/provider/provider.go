// Code generated by ent, DO NOT EDIT.

package provider

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the provider type in the database.
	Label = "provider"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeAppointments holds the string denoting the appointments edge name in mutations.
	EdgeAppointments = "appointments"
	// Table holds the table name of the provider in the database.
	Table = "providers"
	// AppointmentsTable is the table that holds the appointments relation/edge.
	AppointmentsTable = "appointments"
	// AppointmentsInverseTable is the table name for the Appointment entity.
	// It exists in this package in order to avoid circular dependency with the "appointment" package.
	AppointmentsInverseTable = "appointments"
	// AppointmentsColumn is the table column denoting the appointments relation/edge.
	AppointmentsColumn = "provider_appointments"
)

// Columns holds all SQL columns for provider fields.
var Columns = []string{
	FieldID,
	FieldName,
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

// OrderOption defines the ordering options for the Provider queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAppointmentsCount orders the results by appointments count.
func ByAppointmentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAppointmentsStep(), opts...)
	}
}

// ByAppointments orders the results by appointments terms.
func ByAppointments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAppointmentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAppointmentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AppointmentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AppointmentsTable, AppointmentsColumn),
	)
}
// Code generated by ent, DO NOT EDIT.

package appointment

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the appointment type in the database.
	Label = "appointment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldPeriod holds the string denoting the period field in the database.
	FieldPeriod = "period"
	// EdgePatient holds the string denoting the patient edge name in mutations.
	EdgePatient = "patient"
	// EdgeProvider holds the string denoting the provider edge name in mutations.
	EdgeProvider = "provider"
	// EdgeDiagnoses holds the string denoting the diagnoses edge name in mutations.
	EdgeDiagnoses = "diagnoses"
	// EdgeSurvey holds the string denoting the survey edge name in mutations.
	EdgeSurvey = "survey"
	// Table holds the table name of the appointment in the database.
	Table = "appointments"
	// PatientTable is the table that holds the patient relation/edge.
	PatientTable = "appointments"
	// PatientInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientInverseTable = "patients"
	// PatientColumn is the table column denoting the patient relation/edge.
	PatientColumn = "patient_appointments"
	// ProviderTable is the table that holds the provider relation/edge.
	ProviderTable = "appointments"
	// ProviderInverseTable is the table name for the Provider entity.
	// It exists in this package in order to avoid circular dependency with the "provider" package.
	ProviderInverseTable = "providers"
	// ProviderColumn is the table column denoting the provider relation/edge.
	ProviderColumn = "provider_appointments"
	// DiagnosesTable is the table that holds the diagnoses relation/edge. The primary key declared below.
	DiagnosesTable = "appointment_diagnoses"
	// DiagnosesInverseTable is the table name for the Diagnosis entity.
	// It exists in this package in order to avoid circular dependency with the "diagnosis" package.
	DiagnosesInverseTable = "diagnoses"
	// SurveyTable is the table that holds the survey relation/edge.
	SurveyTable = "appointments"
	// SurveyInverseTable is the table name for the Survey entity.
	// It exists in this package in order to avoid circular dependency with the "survey" package.
	SurveyInverseTable = "surveys"
	// SurveyColumn is the table column denoting the survey relation/edge.
	SurveyColumn = "survey_appointments"
)

// Columns holds all SQL columns for appointment fields.
var Columns = []string{
	FieldID,
	FieldStatus,
	FieldPeriod,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "appointments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"patient_appointments",
	"provider_appointments",
	"survey_appointments",
}

var (
	// DiagnosesPrimaryKey and DiagnosesColumn2 are the table columns denoting the
	// primary key for the diagnoses relation (M2M).
	DiagnosesPrimaryKey = []string{"appointment_id", "diagnosis_id"}
)

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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Appointment queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByPatientField orders the results by patient field.
func ByPatientField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPatientStep(), sql.OrderByField(field, opts...))
	}
}

// ByProviderField orders the results by provider field.
func ByProviderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProviderStep(), sql.OrderByField(field, opts...))
	}
}

// ByDiagnosesCount orders the results by diagnoses count.
func ByDiagnosesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDiagnosesStep(), opts...)
	}
}

// ByDiagnoses orders the results by diagnoses terms.
func ByDiagnoses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDiagnosesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySurveyField orders the results by survey field.
func BySurveyField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSurveyStep(), sql.OrderByField(field, opts...))
	}
}
func newPatientStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PatientInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PatientTable, PatientColumn),
	)
}
func newProviderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProviderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProviderTable, ProviderColumn),
	)
}
func newDiagnosesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DiagnosesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, DiagnosesTable, DiagnosesPrimaryKey...),
	)
}
func newSurveyStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SurveyInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SurveyTable, SurveyColumn),
	)
}

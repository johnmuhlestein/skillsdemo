// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"skillsdemo/ent/patient"
	"skillsdemo/ent/schema"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Patient is the model entity for the Patient schema.
type Patient struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name schema.Name `json:"name,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender string `json:"gender,omitempty"`
	// Birthdate holds the value of the "birthdate" field.
	Birthdate time.Time `json:"birthdate,omitempty"`
	// Contact holds the value of the "contact" field.
	Contact []schema.Contact `json:"contact,omitempty"`
	// Address holds the value of the "address" field.
	Address []schema.Address `json:"address,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PatientQuery when eager-loading is set.
	Edges        PatientEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PatientEdges holds the relations/edges for other nodes in the graph.
type PatientEdges struct {
	// Appointments holds the value of the appointments edge.
	Appointments []*Appointment `json:"appointments,omitempty"`
	// Feedbacks holds the value of the feedbacks edge.
	Feedbacks []*Feedback `json:"feedbacks,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// AppointmentsOrErr returns the Appointments value or an error if the edge
// was not loaded in eager-loading.
func (e PatientEdges) AppointmentsOrErr() ([]*Appointment, error) {
	if e.loadedTypes[0] {
		return e.Appointments, nil
	}
	return nil, &NotLoadedError{edge: "appointments"}
}

// FeedbacksOrErr returns the Feedbacks value or an error if the edge
// was not loaded in eager-loading.
func (e PatientEdges) FeedbacksOrErr() ([]*Feedback, error) {
	if e.loadedTypes[1] {
		return e.Feedbacks, nil
	}
	return nil, &NotLoadedError{edge: "feedbacks"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Patient) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case patient.FieldName, patient.FieldContact, patient.FieldAddress:
			values[i] = new([]byte)
		case patient.FieldGender:
			values[i] = new(sql.NullString)
		case patient.FieldBirthdate:
			values[i] = new(sql.NullTime)
		case patient.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Patient fields.
func (pa *Patient) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case patient.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pa.ID = *value
			}
		case patient.FieldName:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pa.Name); err != nil {
					return fmt.Errorf("unmarshal field name: %w", err)
				}
			}
		case patient.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				pa.Gender = value.String
			}
		case patient.FieldBirthdate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field birthdate", values[i])
			} else if value.Valid {
				pa.Birthdate = value.Time
			}
		case patient.FieldContact:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field contact", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pa.Contact); err != nil {
					return fmt.Errorf("unmarshal field contact: %w", err)
				}
			}
		case patient.FieldAddress:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pa.Address); err != nil {
					return fmt.Errorf("unmarshal field address: %w", err)
				}
			}
		default:
			pa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Patient.
// This includes values selected through modifiers, order, etc.
func (pa *Patient) Value(name string) (ent.Value, error) {
	return pa.selectValues.Get(name)
}

// QueryAppointments queries the "appointments" edge of the Patient entity.
func (pa *Patient) QueryAppointments() *AppointmentQuery {
	return NewPatientClient(pa.config).QueryAppointments(pa)
}

// QueryFeedbacks queries the "feedbacks" edge of the Patient entity.
func (pa *Patient) QueryFeedbacks() *FeedbackQuery {
	return NewPatientClient(pa.config).QueryFeedbacks(pa)
}

// Update returns a builder for updating this Patient.
// Note that you need to call Patient.Unwrap() before calling this method if this Patient
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Patient) Update() *PatientUpdateOne {
	return NewPatientClient(pa.config).UpdateOne(pa)
}

// Unwrap unwraps the Patient entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Patient) Unwrap() *Patient {
	_tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Patient is not a transactional entity")
	}
	pa.config.driver = _tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Patient) String() string {
	var builder strings.Builder
	builder.WriteString("Patient(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pa.ID))
	builder.WriteString("name=")
	builder.WriteString(fmt.Sprintf("%v", pa.Name))
	builder.WriteString(", ")
	builder.WriteString("gender=")
	builder.WriteString(pa.Gender)
	builder.WriteString(", ")
	builder.WriteString("birthdate=")
	builder.WriteString(pa.Birthdate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("contact=")
	builder.WriteString(fmt.Sprintf("%v", pa.Contact))
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(fmt.Sprintf("%v", pa.Address))
	builder.WriteByte(')')
	return builder.String()
}

// Patients is a parsable slice of Patient.
type Patients []*Patient

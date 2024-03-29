// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"skillsdemo/ent/survey"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Survey is the model entity for the Survey schema.
type Survey struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Status holds the value of the "status" field.
	Status survey.Status `json:"status,omitempty"`
	// ActiveTime holds the value of the "active_time" field.
	ActiveTime time.Time `json:"active_time,omitempty"`
	// ArchiveTime holds the value of the "archive_time" field.
	ArchiveTime time.Time `json:"archive_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SurveyQuery when eager-loading is set.
	Edges        SurveyEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SurveyEdges holds the relations/edges for other nodes in the graph.
type SurveyEdges struct {
	// Prompts holds the value of the prompts edge.
	Prompts []*Prompt `json:"prompts,omitempty"`
	// Feedbacks holds the value of the feedbacks edge.
	Feedbacks []*Feedback `json:"feedbacks,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PromptsOrErr returns the Prompts value or an error if the edge
// was not loaded in eager-loading.
func (e SurveyEdges) PromptsOrErr() ([]*Prompt, error) {
	if e.loadedTypes[0] {
		return e.Prompts, nil
	}
	return nil, &NotLoadedError{edge: "prompts"}
}

// FeedbacksOrErr returns the Feedbacks value or an error if the edge
// was not loaded in eager-loading.
func (e SurveyEdges) FeedbacksOrErr() ([]*Feedback, error) {
	if e.loadedTypes[1] {
		return e.Feedbacks, nil
	}
	return nil, &NotLoadedError{edge: "feedbacks"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Survey) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case survey.FieldTitle, survey.FieldDescription, survey.FieldStatus:
			values[i] = new(sql.NullString)
		case survey.FieldActiveTime, survey.FieldArchiveTime:
			values[i] = new(sql.NullTime)
		case survey.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Survey fields.
func (s *Survey) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case survey.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case survey.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				s.Title = value.String
			}
		case survey.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				s.Description = value.String
			}
		case survey.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				s.Status = survey.Status(value.String)
			}
		case survey.FieldActiveTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field active_time", values[i])
			} else if value.Valid {
				s.ActiveTime = value.Time
			}
		case survey.FieldArchiveTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field archive_time", values[i])
			} else if value.Valid {
				s.ArchiveTime = value.Time
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Survey.
// This includes values selected through modifiers, order, etc.
func (s *Survey) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryPrompts queries the "prompts" edge of the Survey entity.
func (s *Survey) QueryPrompts() *PromptQuery {
	return NewSurveyClient(s.config).QueryPrompts(s)
}

// QueryFeedbacks queries the "feedbacks" edge of the Survey entity.
func (s *Survey) QueryFeedbacks() *FeedbackQuery {
	return NewSurveyClient(s.config).QueryFeedbacks(s)
}

// Update returns a builder for updating this Survey.
// Note that you need to call Survey.Unwrap() before calling this method if this Survey
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Survey) Update() *SurveyUpdateOne {
	return NewSurveyClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Survey entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Survey) Unwrap() *Survey {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Survey is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Survey) String() string {
	var builder strings.Builder
	builder.WriteString("Survey(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("title=")
	builder.WriteString(s.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(s.Description)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", s.Status))
	builder.WriteString(", ")
	builder.WriteString("active_time=")
	builder.WriteString(s.ActiveTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("archive_time=")
	builder.WriteString(s.ArchiveTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Surveys is a parsable slice of Survey.
type Surveys []*Survey

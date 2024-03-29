// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"skillsdemo/ent/feedback"
	"skillsdemo/ent/promptresponse"
	"skillsdemo/ent/schema"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// PromptResponse is the model entity for the PromptResponse schema.
type PromptResponse struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// ParsedTemplate holds the value of the "parsed_template" field.
	ParsedTemplate string `json:"parsed_template,omitempty"`
	// PromptIndex holds the value of the "prompt_index" field.
	PromptIndex int `json:"prompt_index,omitempty"`
	// RangeValue holds the value of the "range_value" field.
	RangeValue int `json:"range_value,omitempty"`
	// BoolValue holds the value of the "bool_value" field.
	BoolValue string `json:"bool_value,omitempty"`
	// EnumValue holds the value of the "enum_value" field.
	EnumValue schema.MeasureEnum `json:"enum_value,omitempty"`
	// LabelValues holds the value of the "label_values" field.
	LabelValues []string `json:"label_values,omitempty"`
	// FreeformValue holds the value of the "freeform_value" field.
	FreeformValue string `json:"freeform_value,omitempty"`
	// AnsweredTime holds the value of the "answered_time" field.
	AnsweredTime time.Time `json:"answered_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PromptResponseQuery when eager-loading is set.
	Edges              PromptResponseEdges `json:"edges"`
	feedback_responses *uuid.UUID
	selectValues       sql.SelectValues
}

// PromptResponseEdges holds the relations/edges for other nodes in the graph.
type PromptResponseEdges struct {
	// Feedback holds the value of the feedback edge.
	Feedback *Feedback `json:"feedback,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// FeedbackOrErr returns the Feedback value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PromptResponseEdges) FeedbackOrErr() (*Feedback, error) {
	if e.Feedback != nil {
		return e.Feedback, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: feedback.Label}
	}
	return nil, &NotLoadedError{edge: "feedback"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PromptResponse) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case promptresponse.FieldEnumValue, promptresponse.FieldLabelValues:
			values[i] = new([]byte)
		case promptresponse.FieldPromptIndex, promptresponse.FieldRangeValue:
			values[i] = new(sql.NullInt64)
		case promptresponse.FieldParsedTemplate, promptresponse.FieldBoolValue, promptresponse.FieldFreeformValue:
			values[i] = new(sql.NullString)
		case promptresponse.FieldAnsweredTime:
			values[i] = new(sql.NullTime)
		case promptresponse.FieldID:
			values[i] = new(uuid.UUID)
		case promptresponse.ForeignKeys[0]: // feedback_responses
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PromptResponse fields.
func (pr *PromptResponse) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case promptresponse.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case promptresponse.FieldParsedTemplate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parsed_template", values[i])
			} else if value.Valid {
				pr.ParsedTemplate = value.String
			}
		case promptresponse.FieldPromptIndex:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field prompt_index", values[i])
			} else if value.Valid {
				pr.PromptIndex = int(value.Int64)
			}
		case promptresponse.FieldRangeValue:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field range_value", values[i])
			} else if value.Valid {
				pr.RangeValue = int(value.Int64)
			}
		case promptresponse.FieldBoolValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bool_value", values[i])
			} else if value.Valid {
				pr.BoolValue = value.String
			}
		case promptresponse.FieldEnumValue:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field enum_value", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pr.EnumValue); err != nil {
					return fmt.Errorf("unmarshal field enum_value: %w", err)
				}
			}
		case promptresponse.FieldLabelValues:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field label_values", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pr.LabelValues); err != nil {
					return fmt.Errorf("unmarshal field label_values: %w", err)
				}
			}
		case promptresponse.FieldFreeformValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field freeform_value", values[i])
			} else if value.Valid {
				pr.FreeformValue = value.String
			}
		case promptresponse.FieldAnsweredTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field answered_time", values[i])
			} else if value.Valid {
				pr.AnsweredTime = value.Time
			}
		case promptresponse.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field feedback_responses", values[i])
			} else if value.Valid {
				pr.feedback_responses = new(uuid.UUID)
				*pr.feedback_responses = *value.S.(*uuid.UUID)
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PromptResponse.
// This includes values selected through modifiers, order, etc.
func (pr *PromptResponse) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryFeedback queries the "feedback" edge of the PromptResponse entity.
func (pr *PromptResponse) QueryFeedback() *FeedbackQuery {
	return NewPromptResponseClient(pr.config).QueryFeedback(pr)
}

// Update returns a builder for updating this PromptResponse.
// Note that you need to call PromptResponse.Unwrap() before calling this method if this PromptResponse
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *PromptResponse) Update() *PromptResponseUpdateOne {
	return NewPromptResponseClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the PromptResponse entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *PromptResponse) Unwrap() *PromptResponse {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: PromptResponse is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *PromptResponse) String() string {
	var builder strings.Builder
	builder.WriteString("PromptResponse(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("parsed_template=")
	builder.WriteString(pr.ParsedTemplate)
	builder.WriteString(", ")
	builder.WriteString("prompt_index=")
	builder.WriteString(fmt.Sprintf("%v", pr.PromptIndex))
	builder.WriteString(", ")
	builder.WriteString("range_value=")
	builder.WriteString(fmt.Sprintf("%v", pr.RangeValue))
	builder.WriteString(", ")
	builder.WriteString("bool_value=")
	builder.WriteString(pr.BoolValue)
	builder.WriteString(", ")
	builder.WriteString("enum_value=")
	builder.WriteString(fmt.Sprintf("%v", pr.EnumValue))
	builder.WriteString(", ")
	builder.WriteString("label_values=")
	builder.WriteString(fmt.Sprintf("%v", pr.LabelValues))
	builder.WriteString(", ")
	builder.WriteString("freeform_value=")
	builder.WriteString(pr.FreeformValue)
	builder.WriteString(", ")
	builder.WriteString("answered_time=")
	builder.WriteString(pr.AnsweredTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// PromptResponses is a parsable slice of PromptResponse.
type PromptResponses []*PromptResponse

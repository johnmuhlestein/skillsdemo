// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"skillsdemo/ent/prompt"
	"skillsdemo/ent/schema"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Prompt is the model entity for the Prompt schema.
type Prompt struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// SortOrder holds the value of the "sort_order" field.
	SortOrder int `json:"sort_order,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// ResponseType holds the value of the "response_type" field.
	ResponseType schema.Measure `json:"response_type,omitempty"`
	// AdditionalFeedback holds the value of the "additional_feedback" field.
	AdditionalFeedback bool `json:"additional_feedback,omitempty"`
	survey_prompts     *uuid.UUID
	selectValues       sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Prompt) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case prompt.FieldResponseType:
			values[i] = new([]byte)
		case prompt.FieldAdditionalFeedback:
			values[i] = new(sql.NullBool)
		case prompt.FieldSortOrder:
			values[i] = new(sql.NullInt64)
		case prompt.FieldTitle, prompt.FieldDescription:
			values[i] = new(sql.NullString)
		case prompt.FieldID:
			values[i] = new(uuid.UUID)
		case prompt.ForeignKeys[0]: // survey_prompts
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Prompt fields.
func (pr *Prompt) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case prompt.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case prompt.FieldSortOrder:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort_order", values[i])
			} else if value.Valid {
				pr.SortOrder = int(value.Int64)
			}
		case prompt.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				pr.Title = value.String
			}
		case prompt.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pr.Description = value.String
			}
		case prompt.FieldResponseType:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field response_type", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pr.ResponseType); err != nil {
					return fmt.Errorf("unmarshal field response_type: %w", err)
				}
			}
		case prompt.FieldAdditionalFeedback:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field additional_feedback", values[i])
			} else if value.Valid {
				pr.AdditionalFeedback = value.Bool
			}
		case prompt.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field survey_prompts", values[i])
			} else if value.Valid {
				pr.survey_prompts = new(uuid.UUID)
				*pr.survey_prompts = *value.S.(*uuid.UUID)
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Prompt.
// This includes values selected through modifiers, order, etc.
func (pr *Prompt) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// Update returns a builder for updating this Prompt.
// Note that you need to call Prompt.Unwrap() before calling this method if this Prompt
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Prompt) Update() *PromptUpdateOne {
	return NewPromptClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Prompt entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Prompt) Unwrap() *Prompt {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Prompt is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Prompt) String() string {
	var builder strings.Builder
	builder.WriteString("Prompt(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("sort_order=")
	builder.WriteString(fmt.Sprintf("%v", pr.SortOrder))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(pr.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pr.Description)
	builder.WriteString(", ")
	builder.WriteString("response_type=")
	builder.WriteString(fmt.Sprintf("%v", pr.ResponseType))
	builder.WriteString(", ")
	builder.WriteString("additional_feedback=")
	builder.WriteString(fmt.Sprintf("%v", pr.AdditionalFeedback))
	builder.WriteByte(')')
	return builder.String()
}

// Prompts is a parsable slice of Prompt.
type Prompts []*Prompt
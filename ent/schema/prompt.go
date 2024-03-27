package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Prompt holds the schema definition for the Prompt entity.
type Prompt struct {
	ent.Schema
}

type MeasureEnum struct {
	Value int `json:"value"`
	Label string `json:"label"`
}
type Measure struct {
	Type string `json:"type"`
	LowRange int `json:"lowRange"`
	HighRange int `json:"highRange"`
	BooleanTrue string `json:"booleanTrue"`
	BooleanFalse string `json:"booleanFalse"`
	Enumeration []MeasureEnum `json:"enumeration"`
	Flags []string `json:"flags"`
}

// Fields of the Prompt.
func (Prompt) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Int("sort_order").Positive(),
		field.String("title").Unique(),
		field.String("description").Optional(),
		field.JSON("response_type", Measure{}),
		field.Bool("additional_feedback").Default(false),
	}
}

// Edges of the Prompt.
func (Prompt) Edges() []ent.Edge {
	return nil
}

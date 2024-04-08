package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PromptResponse holds the schema definition for the PromptResponse entity.
type PromptResponse struct {
	ent.Schema
}

// Fields of the PromptResponse.
func (PromptResponse) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("parsed_template"),
		field.Int("prompt_index").Positive(),
		field.Int("range_value").Optional().Nillable(),
		field.String("bool_value").Optional().Nillable(),
		field.JSON("enum_value", MeasureEnum{}).Optional(),
		field.JSON("label_values", []string{}).Optional(),
		field.String("freeform_value").Optional().Nillable(),
		field.Time("answered_time").Optional().Nillable(),
	}
}

// Edges of the PromptResponse.
func (PromptResponse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("feedback", Feedback.Type).Ref("responses").Unique(),
	}
}

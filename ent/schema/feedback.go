package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Feedback holds the schema definition for the Feedback entity.
type Feedback struct {
	ent.Schema
}

// Fields of the Feedback.
func (Feedback) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("status"),
		field.Time("start_time"),
		field.Time("completion_time"),
	}
}

// Edges of the Feedback.
func (Feedback) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("responses", PromptResponse.Type),
		edge.From("patient", Patient.Type).Ref("feedbacks").Unique(),
		edge.From("survey", Survey.Type).Ref("feedbacks").Unique(),
	}
}

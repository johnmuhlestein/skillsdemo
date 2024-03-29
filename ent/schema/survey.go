package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Survey holds the schema definition for the Survey entity.
type Survey struct {
	ent.Schema
}

// Fields of the Survey.
func (Survey) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("title"),
		field.String("description"),
		field.Enum("status").Values("unpublished","active","archived").Default("unpublished"),
		field.Time("active_time").Optional(),
		field.Time("archive_time").Optional(),
	}
}

// Edges of the Survey.
func (Survey) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("prompts", Prompt.Type),
		edge.To("feedbacks", Feedback.Type),
	}
}

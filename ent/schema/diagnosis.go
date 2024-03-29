package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Diagnosis holds the schema definition for the Diagnosis entity.
type Diagnosis struct {
	ent.Schema
}

type Coding struct {
	System string `json:"system"`
	Code string `json:"code"`
	Name string `json:"name"`
}
type Code struct {
	Coding []Coding `json:"coding"`
}

// Fields of the Diagnosis.
func (Diagnosis) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("status"),
		field.Time("last_updated").Default(time.Now()),
		field.JSON("code", Code{}),
	}
}

// Edges of the Diagnosis.
func (Diagnosis) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("appointments", Appointment.Type).Ref("diagnoses"),
	}
}

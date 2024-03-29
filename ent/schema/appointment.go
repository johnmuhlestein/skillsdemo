package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Appointment holds the schema definition for the Appointment entity.
type Appointment struct {
	ent.Schema
}

type Period struct {
	Start time.Time `json:"starttime"`
	End time.Time `json:"endtime"`
}

// Fields of the Appointment.
func (Appointment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("status"),
		field.JSON("period", Period{}),
	}
}

// Edges of the Appointment.
func (Appointment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("patient", Patient.Type).Ref("appointments").Unique(),
		edge.From("provider", Provider.Type).Ref("appointments").Unique(),
		edge.To("diagnoses", Diagnosis.Type),
		edge.From("survey", Survey.Type).Ref("appointments").Unique(),
	}
}

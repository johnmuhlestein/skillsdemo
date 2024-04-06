package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Patient holds the schema definition for the Patient entity.
type Patient struct {
	ent.Schema
}

type Name struct {
	Text string `json:"text"`
	Family string `json:"family"`
	Given string `json:"given"`
}

type Contact struct {
	System string `json:"system"`
	Value string `json:"value"`
	Use string `json:"use"`
}

type Address struct {
	Use string `json:"use"`
	Line []string `json:"line"`
	State string `json:"state"`
	City string `json:"city"`
	PostalCode string `json:"postalcode"`
}

// Fields of the Patient.
func (Patient) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.JSON("name",Name{}),
		field.String("gender"),
		field.Time("birthdate").Optional().SchemaType(map[string]string{dialect.Postgres: "date"}),
		field.JSON("contact", []Contact{}).Optional(),
		field.JSON("address", []Address{}).Optional(),
	}
}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("appointments", Appointment.Type),
		edge.To("feedbacks", Feedback.Type),
	}
}

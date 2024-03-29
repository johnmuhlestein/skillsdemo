// Code generated by ent, DO NOT EDIT.

package diagnosis

import (
	"skillsdemo/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldLTE(FieldID, id))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldEQ(FieldStatus, v))
}

// LastUpdated applies equality check predicate on the "last_updated" field. It's identical to LastUpdatedEQ.
func LastUpdated(v time.Time) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldEQ(FieldLastUpdated, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldContainsFold(FieldStatus, v))
}

// LastUpdatedEQ applies the EQ predicate on the "last_updated" field.
func LastUpdatedEQ(v time.Time) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldEQ(FieldLastUpdated, v))
}

// LastUpdatedNEQ applies the NEQ predicate on the "last_updated" field.
func LastUpdatedNEQ(v time.Time) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldNEQ(FieldLastUpdated, v))
}

// LastUpdatedIn applies the In predicate on the "last_updated" field.
func LastUpdatedIn(vs ...time.Time) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldIn(FieldLastUpdated, vs...))
}

// LastUpdatedNotIn applies the NotIn predicate on the "last_updated" field.
func LastUpdatedNotIn(vs ...time.Time) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldNotIn(FieldLastUpdated, vs...))
}

// LastUpdatedGT applies the GT predicate on the "last_updated" field.
func LastUpdatedGT(v time.Time) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldGT(FieldLastUpdated, v))
}

// LastUpdatedGTE applies the GTE predicate on the "last_updated" field.
func LastUpdatedGTE(v time.Time) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldGTE(FieldLastUpdated, v))
}

// LastUpdatedLT applies the LT predicate on the "last_updated" field.
func LastUpdatedLT(v time.Time) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldLT(FieldLastUpdated, v))
}

// LastUpdatedLTE applies the LTE predicate on the "last_updated" field.
func LastUpdatedLTE(v time.Time) predicate.Diagnosis {
	return predicate.Diagnosis(sql.FieldLTE(FieldLastUpdated, v))
}

// HasAppointments applies the HasEdge predicate on the "appointments" edge.
func HasAppointments() predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, AppointmentsTable, AppointmentsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAppointmentsWith applies the HasEdge predicate on the "appointments" edge with a given conditions (other predicates).
func HasAppointmentsWith(preds ...predicate.Appointment) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		step := newAppointmentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Diagnosis) predicate.Diagnosis {
	return predicate.Diagnosis(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Diagnosis) predicate.Diagnosis {
	return predicate.Diagnosis(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Diagnosis) predicate.Diagnosis {
	return predicate.Diagnosis(sql.NotPredicates(p))
}

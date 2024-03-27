// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"skillsdemo/ent/appointment"
	"skillsdemo/ent/predicate"
	"skillsdemo/ent/provider"
	"skillsdemo/ent/schema"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ProviderUpdate is the builder for updating Provider entities.
type ProviderUpdate struct {
	config
	hooks    []Hook
	mutation *ProviderMutation
}

// Where appends a list predicates to the ProviderUpdate builder.
func (pu *ProviderUpdate) Where(ps ...predicate.Provider) *ProviderUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *ProviderUpdate) SetName(s schema.Name) *ProviderUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *ProviderUpdate) SetNillableName(s *schema.Name) *ProviderUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// AddAppointmentIDs adds the "appointments" edge to the Appointment entity by IDs.
func (pu *ProviderUpdate) AddAppointmentIDs(ids ...uuid.UUID) *ProviderUpdate {
	pu.mutation.AddAppointmentIDs(ids...)
	return pu
}

// AddAppointments adds the "appointments" edges to the Appointment entity.
func (pu *ProviderUpdate) AddAppointments(a ...*Appointment) *ProviderUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pu.AddAppointmentIDs(ids...)
}

// Mutation returns the ProviderMutation object of the builder.
func (pu *ProviderUpdate) Mutation() *ProviderMutation {
	return pu.mutation
}

// ClearAppointments clears all "appointments" edges to the Appointment entity.
func (pu *ProviderUpdate) ClearAppointments() *ProviderUpdate {
	pu.mutation.ClearAppointments()
	return pu
}

// RemoveAppointmentIDs removes the "appointments" edge to Appointment entities by IDs.
func (pu *ProviderUpdate) RemoveAppointmentIDs(ids ...uuid.UUID) *ProviderUpdate {
	pu.mutation.RemoveAppointmentIDs(ids...)
	return pu
}

// RemoveAppointments removes "appointments" edges to Appointment entities.
func (pu *ProviderUpdate) RemoveAppointments(a ...*Appointment) *ProviderUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pu.RemoveAppointmentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProviderUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProviderUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProviderUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProviderUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProviderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(provider.Table, provider.Columns, sqlgraph.NewFieldSpec(provider.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(provider.FieldName, field.TypeJSON, value)
	}
	if pu.mutation.AppointmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   provider.AppointmentsTable,
			Columns: []string{provider.AppointmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appointment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedAppointmentsIDs(); len(nodes) > 0 && !pu.mutation.AppointmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   provider.AppointmentsTable,
			Columns: []string{provider.AppointmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appointment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.AppointmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   provider.AppointmentsTable,
			Columns: []string{provider.AppointmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appointment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{provider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProviderUpdateOne is the builder for updating a single Provider entity.
type ProviderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProviderMutation
}

// SetName sets the "name" field.
func (puo *ProviderUpdateOne) SetName(s schema.Name) *ProviderUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *ProviderUpdateOne) SetNillableName(s *schema.Name) *ProviderUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// AddAppointmentIDs adds the "appointments" edge to the Appointment entity by IDs.
func (puo *ProviderUpdateOne) AddAppointmentIDs(ids ...uuid.UUID) *ProviderUpdateOne {
	puo.mutation.AddAppointmentIDs(ids...)
	return puo
}

// AddAppointments adds the "appointments" edges to the Appointment entity.
func (puo *ProviderUpdateOne) AddAppointments(a ...*Appointment) *ProviderUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return puo.AddAppointmentIDs(ids...)
}

// Mutation returns the ProviderMutation object of the builder.
func (puo *ProviderUpdateOne) Mutation() *ProviderMutation {
	return puo.mutation
}

// ClearAppointments clears all "appointments" edges to the Appointment entity.
func (puo *ProviderUpdateOne) ClearAppointments() *ProviderUpdateOne {
	puo.mutation.ClearAppointments()
	return puo
}

// RemoveAppointmentIDs removes the "appointments" edge to Appointment entities by IDs.
func (puo *ProviderUpdateOne) RemoveAppointmentIDs(ids ...uuid.UUID) *ProviderUpdateOne {
	puo.mutation.RemoveAppointmentIDs(ids...)
	return puo
}

// RemoveAppointments removes "appointments" edges to Appointment entities.
func (puo *ProviderUpdateOne) RemoveAppointments(a ...*Appointment) *ProviderUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return puo.RemoveAppointmentIDs(ids...)
}

// Where appends a list predicates to the ProviderUpdate builder.
func (puo *ProviderUpdateOne) Where(ps ...predicate.Provider) *ProviderUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProviderUpdateOne) Select(field string, fields ...string) *ProviderUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Provider entity.
func (puo *ProviderUpdateOne) Save(ctx context.Context) (*Provider, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProviderUpdateOne) SaveX(ctx context.Context) *Provider {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProviderUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProviderUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProviderUpdateOne) sqlSave(ctx context.Context) (_node *Provider, err error) {
	_spec := sqlgraph.NewUpdateSpec(provider.Table, provider.Columns, sqlgraph.NewFieldSpec(provider.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Provider.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, provider.FieldID)
		for _, f := range fields {
			if !provider.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != provider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(provider.FieldName, field.TypeJSON, value)
	}
	if puo.mutation.AppointmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   provider.AppointmentsTable,
			Columns: []string{provider.AppointmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appointment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedAppointmentsIDs(); len(nodes) > 0 && !puo.mutation.AppointmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   provider.AppointmentsTable,
			Columns: []string{provider.AppointmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appointment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.AppointmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   provider.AppointmentsTable,
			Columns: []string{provider.AppointmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appointment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Provider{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{provider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
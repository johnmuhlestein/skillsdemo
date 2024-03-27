// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"skillsdemo/ent/appointment"
	"skillsdemo/ent/diagnosis"
	"skillsdemo/ent/schema"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DiagnosisCreate is the builder for creating a Diagnosis entity.
type DiagnosisCreate struct {
	config
	mutation *DiagnosisMutation
	hooks    []Hook
}

// SetStatus sets the "status" field.
func (dc *DiagnosisCreate) SetStatus(s string) *DiagnosisCreate {
	dc.mutation.SetStatus(s)
	return dc
}

// SetLastUpdated sets the "lastUpdated" field.
func (dc *DiagnosisCreate) SetLastUpdated(t time.Time) *DiagnosisCreate {
	dc.mutation.SetLastUpdated(t)
	return dc
}

// SetNillableLastUpdated sets the "lastUpdated" field if the given value is not nil.
func (dc *DiagnosisCreate) SetNillableLastUpdated(t *time.Time) *DiagnosisCreate {
	if t != nil {
		dc.SetLastUpdated(*t)
	}
	return dc
}

// SetCode sets the "code" field.
func (dc *DiagnosisCreate) SetCode(s schema.Code) *DiagnosisCreate {
	dc.mutation.SetCode(s)
	return dc
}

// SetID sets the "id" field.
func (dc *DiagnosisCreate) SetID(u uuid.UUID) *DiagnosisCreate {
	dc.mutation.SetID(u)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DiagnosisCreate) SetNillableID(u *uuid.UUID) *DiagnosisCreate {
	if u != nil {
		dc.SetID(*u)
	}
	return dc
}

// AddAppointmentIDs adds the "appointment" edge to the Appointment entity by IDs.
func (dc *DiagnosisCreate) AddAppointmentIDs(ids ...uuid.UUID) *DiagnosisCreate {
	dc.mutation.AddAppointmentIDs(ids...)
	return dc
}

// AddAppointment adds the "appointment" edges to the Appointment entity.
func (dc *DiagnosisCreate) AddAppointment(a ...*Appointment) *DiagnosisCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return dc.AddAppointmentIDs(ids...)
}

// Mutation returns the DiagnosisMutation object of the builder.
func (dc *DiagnosisCreate) Mutation() *DiagnosisMutation {
	return dc.mutation
}

// Save creates the Diagnosis in the database.
func (dc *DiagnosisCreate) Save(ctx context.Context) (*Diagnosis, error) {
	dc.defaults()
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DiagnosisCreate) SaveX(ctx context.Context) *Diagnosis {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DiagnosisCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DiagnosisCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DiagnosisCreate) defaults() {
	if _, ok := dc.mutation.LastUpdated(); !ok {
		v := diagnosis.DefaultLastUpdated
		dc.mutation.SetLastUpdated(v)
	}
	if _, ok := dc.mutation.ID(); !ok {
		v := diagnosis.DefaultID()
		dc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DiagnosisCreate) check() error {
	if _, ok := dc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Diagnosis.status"`)}
	}
	if _, ok := dc.mutation.LastUpdated(); !ok {
		return &ValidationError{Name: "lastUpdated", err: errors.New(`ent: missing required field "Diagnosis.lastUpdated"`)}
	}
	if _, ok := dc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Diagnosis.code"`)}
	}
	return nil
}

func (dc *DiagnosisCreate) sqlSave(ctx context.Context) (*Diagnosis, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DiagnosisCreate) createSpec() (*Diagnosis, *sqlgraph.CreateSpec) {
	var (
		_node = &Diagnosis{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(diagnosis.Table, sqlgraph.NewFieldSpec(diagnosis.FieldID, field.TypeUUID))
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dc.mutation.Status(); ok {
		_spec.SetField(diagnosis.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := dc.mutation.LastUpdated(); ok {
		_spec.SetField(diagnosis.FieldLastUpdated, field.TypeTime, value)
		_node.LastUpdated = value
	}
	if value, ok := dc.mutation.Code(); ok {
		_spec.SetField(diagnosis.FieldCode, field.TypeJSON, value)
		_node.Code = value
	}
	if nodes := dc.mutation.AppointmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   diagnosis.AppointmentTable,
			Columns: diagnosis.AppointmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appointment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DiagnosisCreateBulk is the builder for creating many Diagnosis entities in bulk.
type DiagnosisCreateBulk struct {
	config
	err      error
	builders []*DiagnosisCreate
}

// Save creates the Diagnosis entities in the database.
func (dcb *DiagnosisCreateBulk) Save(ctx context.Context) ([]*Diagnosis, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Diagnosis, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DiagnosisMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DiagnosisCreateBulk) SaveX(ctx context.Context) []*Diagnosis {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DiagnosisCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DiagnosisCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
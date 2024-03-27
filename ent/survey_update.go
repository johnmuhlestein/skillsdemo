// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"skillsdemo/ent/predicate"
	"skillsdemo/ent/prompt"
	"skillsdemo/ent/survey"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SurveyUpdate is the builder for updating Survey entities.
type SurveyUpdate struct {
	config
	hooks    []Hook
	mutation *SurveyMutation
}

// Where appends a list predicates to the SurveyUpdate builder.
func (su *SurveyUpdate) Where(ps ...predicate.Survey) *SurveyUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetTitle sets the "title" field.
func (su *SurveyUpdate) SetTitle(s string) *SurveyUpdate {
	su.mutation.SetTitle(s)
	return su
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (su *SurveyUpdate) SetNillableTitle(s *string) *SurveyUpdate {
	if s != nil {
		su.SetTitle(*s)
	}
	return su
}

// SetDescription sets the "description" field.
func (su *SurveyUpdate) SetDescription(s string) *SurveyUpdate {
	su.mutation.SetDescription(s)
	return su
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (su *SurveyUpdate) SetNillableDescription(s *string) *SurveyUpdate {
	if s != nil {
		su.SetDescription(*s)
	}
	return su
}

// SetStatus sets the "status" field.
func (su *SurveyUpdate) SetStatus(s survey.Status) *SurveyUpdate {
	su.mutation.SetStatus(s)
	return su
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (su *SurveyUpdate) SetNillableStatus(s *survey.Status) *SurveyUpdate {
	if s != nil {
		su.SetStatus(*s)
	}
	return su
}

// SetActiveTime sets the "active_time" field.
func (su *SurveyUpdate) SetActiveTime(t time.Time) *SurveyUpdate {
	su.mutation.SetActiveTime(t)
	return su
}

// SetNillableActiveTime sets the "active_time" field if the given value is not nil.
func (su *SurveyUpdate) SetNillableActiveTime(t *time.Time) *SurveyUpdate {
	if t != nil {
		su.SetActiveTime(*t)
	}
	return su
}

// SetArchiveTime sets the "archive_time" field.
func (su *SurveyUpdate) SetArchiveTime(t time.Time) *SurveyUpdate {
	su.mutation.SetArchiveTime(t)
	return su
}

// SetNillableArchiveTime sets the "archive_time" field if the given value is not nil.
func (su *SurveyUpdate) SetNillableArchiveTime(t *time.Time) *SurveyUpdate {
	if t != nil {
		su.SetArchiveTime(*t)
	}
	return su
}

// AddPromptIDs adds the "prompts" edge to the Prompt entity by IDs.
func (su *SurveyUpdate) AddPromptIDs(ids ...uuid.UUID) *SurveyUpdate {
	su.mutation.AddPromptIDs(ids...)
	return su
}

// AddPrompts adds the "prompts" edges to the Prompt entity.
func (su *SurveyUpdate) AddPrompts(p ...*Prompt) *SurveyUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return su.AddPromptIDs(ids...)
}

// Mutation returns the SurveyMutation object of the builder.
func (su *SurveyUpdate) Mutation() *SurveyMutation {
	return su.mutation
}

// ClearPrompts clears all "prompts" edges to the Prompt entity.
func (su *SurveyUpdate) ClearPrompts() *SurveyUpdate {
	su.mutation.ClearPrompts()
	return su
}

// RemovePromptIDs removes the "prompts" edge to Prompt entities by IDs.
func (su *SurveyUpdate) RemovePromptIDs(ids ...uuid.UUID) *SurveyUpdate {
	su.mutation.RemovePromptIDs(ids...)
	return su
}

// RemovePrompts removes "prompts" edges to Prompt entities.
func (su *SurveyUpdate) RemovePrompts(p ...*Prompt) *SurveyUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return su.RemovePromptIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SurveyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SurveyUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SurveyUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SurveyUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SurveyUpdate) check() error {
	if v, ok := su.mutation.Status(); ok {
		if err := survey.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Survey.status": %w`, err)}
		}
	}
	return nil
}

func (su *SurveyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(survey.Table, survey.Columns, sqlgraph.NewFieldSpec(survey.FieldID, field.TypeUUID))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Title(); ok {
		_spec.SetField(survey.FieldTitle, field.TypeString, value)
	}
	if value, ok := su.mutation.Description(); ok {
		_spec.SetField(survey.FieldDescription, field.TypeString, value)
	}
	if value, ok := su.mutation.Status(); ok {
		_spec.SetField(survey.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := su.mutation.ActiveTime(); ok {
		_spec.SetField(survey.FieldActiveTime, field.TypeTime, value)
	}
	if value, ok := su.mutation.ArchiveTime(); ok {
		_spec.SetField(survey.FieldArchiveTime, field.TypeTime, value)
	}
	if su.mutation.PromptsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   survey.PromptsTable,
			Columns: []string{survey.PromptsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(prompt.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedPromptsIDs(); len(nodes) > 0 && !su.mutation.PromptsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   survey.PromptsTable,
			Columns: []string{survey.PromptsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(prompt.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.PromptsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   survey.PromptsTable,
			Columns: []string{survey.PromptsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(prompt.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{survey.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SurveyUpdateOne is the builder for updating a single Survey entity.
type SurveyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SurveyMutation
}

// SetTitle sets the "title" field.
func (suo *SurveyUpdateOne) SetTitle(s string) *SurveyUpdateOne {
	suo.mutation.SetTitle(s)
	return suo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (suo *SurveyUpdateOne) SetNillableTitle(s *string) *SurveyUpdateOne {
	if s != nil {
		suo.SetTitle(*s)
	}
	return suo
}

// SetDescription sets the "description" field.
func (suo *SurveyUpdateOne) SetDescription(s string) *SurveyUpdateOne {
	suo.mutation.SetDescription(s)
	return suo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (suo *SurveyUpdateOne) SetNillableDescription(s *string) *SurveyUpdateOne {
	if s != nil {
		suo.SetDescription(*s)
	}
	return suo
}

// SetStatus sets the "status" field.
func (suo *SurveyUpdateOne) SetStatus(s survey.Status) *SurveyUpdateOne {
	suo.mutation.SetStatus(s)
	return suo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (suo *SurveyUpdateOne) SetNillableStatus(s *survey.Status) *SurveyUpdateOne {
	if s != nil {
		suo.SetStatus(*s)
	}
	return suo
}

// SetActiveTime sets the "active_time" field.
func (suo *SurveyUpdateOne) SetActiveTime(t time.Time) *SurveyUpdateOne {
	suo.mutation.SetActiveTime(t)
	return suo
}

// SetNillableActiveTime sets the "active_time" field if the given value is not nil.
func (suo *SurveyUpdateOne) SetNillableActiveTime(t *time.Time) *SurveyUpdateOne {
	if t != nil {
		suo.SetActiveTime(*t)
	}
	return suo
}

// SetArchiveTime sets the "archive_time" field.
func (suo *SurveyUpdateOne) SetArchiveTime(t time.Time) *SurveyUpdateOne {
	suo.mutation.SetArchiveTime(t)
	return suo
}

// SetNillableArchiveTime sets the "archive_time" field if the given value is not nil.
func (suo *SurveyUpdateOne) SetNillableArchiveTime(t *time.Time) *SurveyUpdateOne {
	if t != nil {
		suo.SetArchiveTime(*t)
	}
	return suo
}

// AddPromptIDs adds the "prompts" edge to the Prompt entity by IDs.
func (suo *SurveyUpdateOne) AddPromptIDs(ids ...uuid.UUID) *SurveyUpdateOne {
	suo.mutation.AddPromptIDs(ids...)
	return suo
}

// AddPrompts adds the "prompts" edges to the Prompt entity.
func (suo *SurveyUpdateOne) AddPrompts(p ...*Prompt) *SurveyUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return suo.AddPromptIDs(ids...)
}

// Mutation returns the SurveyMutation object of the builder.
func (suo *SurveyUpdateOne) Mutation() *SurveyMutation {
	return suo.mutation
}

// ClearPrompts clears all "prompts" edges to the Prompt entity.
func (suo *SurveyUpdateOne) ClearPrompts() *SurveyUpdateOne {
	suo.mutation.ClearPrompts()
	return suo
}

// RemovePromptIDs removes the "prompts" edge to Prompt entities by IDs.
func (suo *SurveyUpdateOne) RemovePromptIDs(ids ...uuid.UUID) *SurveyUpdateOne {
	suo.mutation.RemovePromptIDs(ids...)
	return suo
}

// RemovePrompts removes "prompts" edges to Prompt entities.
func (suo *SurveyUpdateOne) RemovePrompts(p ...*Prompt) *SurveyUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return suo.RemovePromptIDs(ids...)
}

// Where appends a list predicates to the SurveyUpdate builder.
func (suo *SurveyUpdateOne) Where(ps ...predicate.Survey) *SurveyUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SurveyUpdateOne) Select(field string, fields ...string) *SurveyUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Survey entity.
func (suo *SurveyUpdateOne) Save(ctx context.Context) (*Survey, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SurveyUpdateOne) SaveX(ctx context.Context) *Survey {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SurveyUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SurveyUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SurveyUpdateOne) check() error {
	if v, ok := suo.mutation.Status(); ok {
		if err := survey.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Survey.status": %w`, err)}
		}
	}
	return nil
}

func (suo *SurveyUpdateOne) sqlSave(ctx context.Context) (_node *Survey, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(survey.Table, survey.Columns, sqlgraph.NewFieldSpec(survey.FieldID, field.TypeUUID))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Survey.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, survey.FieldID)
		for _, f := range fields {
			if !survey.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != survey.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Title(); ok {
		_spec.SetField(survey.FieldTitle, field.TypeString, value)
	}
	if value, ok := suo.mutation.Description(); ok {
		_spec.SetField(survey.FieldDescription, field.TypeString, value)
	}
	if value, ok := suo.mutation.Status(); ok {
		_spec.SetField(survey.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := suo.mutation.ActiveTime(); ok {
		_spec.SetField(survey.FieldActiveTime, field.TypeTime, value)
	}
	if value, ok := suo.mutation.ArchiveTime(); ok {
		_spec.SetField(survey.FieldArchiveTime, field.TypeTime, value)
	}
	if suo.mutation.PromptsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   survey.PromptsTable,
			Columns: []string{survey.PromptsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(prompt.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedPromptsIDs(); len(nodes) > 0 && !suo.mutation.PromptsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   survey.PromptsTable,
			Columns: []string{survey.PromptsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(prompt.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.PromptsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   survey.PromptsTable,
			Columns: []string{survey.PromptsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(prompt.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Survey{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{survey.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}

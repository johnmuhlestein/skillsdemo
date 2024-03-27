// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"skillsdemo/ent/appointment"
	"skillsdemo/ent/diagnosis"
	"skillsdemo/ent/patient"
	"skillsdemo/ent/predicate"
	"skillsdemo/ent/provider"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AppointmentQuery is the builder for querying Appointment entities.
type AppointmentQuery struct {
	config
	ctx           *QueryContext
	order         []appointment.OrderOption
	inters        []Interceptor
	predicates    []predicate.Appointment
	withPatient   *PatientQuery
	withProvider  *ProviderQuery
	withDiagnoses *DiagnosisQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppointmentQuery builder.
func (aq *AppointmentQuery) Where(ps ...predicate.Appointment) *AppointmentQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *AppointmentQuery) Limit(limit int) *AppointmentQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *AppointmentQuery) Offset(offset int) *AppointmentQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AppointmentQuery) Unique(unique bool) *AppointmentQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *AppointmentQuery) Order(o ...appointment.OrderOption) *AppointmentQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryPatient chains the current query on the "patient" edge.
func (aq *AppointmentQuery) QueryPatient() *PatientQuery {
	query := (&PatientClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(appointment.Table, appointment.FieldID, selector),
			sqlgraph.To(patient.Table, patient.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, appointment.PatientTable, appointment.PatientColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProvider chains the current query on the "provider" edge.
func (aq *AppointmentQuery) QueryProvider() *ProviderQuery {
	query := (&ProviderClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(appointment.Table, appointment.FieldID, selector),
			sqlgraph.To(provider.Table, provider.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, appointment.ProviderTable, appointment.ProviderColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDiagnoses chains the current query on the "diagnoses" edge.
func (aq *AppointmentQuery) QueryDiagnoses() *DiagnosisQuery {
	query := (&DiagnosisClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(appointment.Table, appointment.FieldID, selector),
			sqlgraph.To(diagnosis.Table, diagnosis.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, appointment.DiagnosesTable, appointment.DiagnosesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Appointment entity from the query.
// Returns a *NotFoundError when no Appointment was found.
func (aq *AppointmentQuery) First(ctx context.Context) (*Appointment, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appointment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AppointmentQuery) FirstX(ctx context.Context) *Appointment {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Appointment ID from the query.
// Returns a *NotFoundError when no Appointment ID was found.
func (aq *AppointmentQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appointment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AppointmentQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Appointment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Appointment entity is found.
// Returns a *NotFoundError when no Appointment entities are found.
func (aq *AppointmentQuery) Only(ctx context.Context) (*Appointment, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appointment.Label}
	default:
		return nil, &NotSingularError{appointment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AppointmentQuery) OnlyX(ctx context.Context) *Appointment {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Appointment ID in the query.
// Returns a *NotSingularError when more than one Appointment ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AppointmentQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appointment.Label}
	default:
		err = &NotSingularError{appointment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AppointmentQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Appointments.
func (aq *AppointmentQuery) All(ctx context.Context) ([]*Appointment, error) {
	ctx = setContextOp(ctx, aq.ctx, "All")
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Appointment, *AppointmentQuery]()
	return withInterceptors[[]*Appointment](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *AppointmentQuery) AllX(ctx context.Context) []*Appointment {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Appointment IDs.
func (aq *AppointmentQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, "IDs")
	if err = aq.Select(appointment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AppointmentQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AppointmentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, "Count")
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*AppointmentQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AppointmentQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AppointmentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, "Exist")
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AppointmentQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppointmentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AppointmentQuery) Clone() *AppointmentQuery {
	if aq == nil {
		return nil
	}
	return &AppointmentQuery{
		config:        aq.config,
		ctx:           aq.ctx.Clone(),
		order:         append([]appointment.OrderOption{}, aq.order...),
		inters:        append([]Interceptor{}, aq.inters...),
		predicates:    append([]predicate.Appointment{}, aq.predicates...),
		withPatient:   aq.withPatient.Clone(),
		withProvider:  aq.withProvider.Clone(),
		withDiagnoses: aq.withDiagnoses.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithPatient tells the query-builder to eager-load the nodes that are connected to
// the "patient" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AppointmentQuery) WithPatient(opts ...func(*PatientQuery)) *AppointmentQuery {
	query := (&PatientClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withPatient = query
	return aq
}

// WithProvider tells the query-builder to eager-load the nodes that are connected to
// the "provider" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AppointmentQuery) WithProvider(opts ...func(*ProviderQuery)) *AppointmentQuery {
	query := (&ProviderClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withProvider = query
	return aq
}

// WithDiagnoses tells the query-builder to eager-load the nodes that are connected to
// the "diagnoses" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AppointmentQuery) WithDiagnoses(opts ...func(*DiagnosisQuery)) *AppointmentQuery {
	query := (&DiagnosisClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withDiagnoses = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Status string `json:"status,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Appointment.Query().
//		GroupBy(appointment.FieldStatus).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AppointmentQuery) GroupBy(field string, fields ...string) *AppointmentGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AppointmentGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = appointment.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Status string `json:"status,omitempty"`
//	}
//
//	client.Appointment.Query().
//		Select(appointment.FieldStatus).
//		Scan(ctx, &v)
func (aq *AppointmentQuery) Select(fields ...string) *AppointmentSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &AppointmentSelect{AppointmentQuery: aq}
	sbuild.label = appointment.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AppointmentSelect configured with the given aggregations.
func (aq *AppointmentQuery) Aggregate(fns ...AggregateFunc) *AppointmentSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AppointmentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !appointment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AppointmentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Appointment, error) {
	var (
		nodes       = []*Appointment{}
		withFKs     = aq.withFKs
		_spec       = aq.querySpec()
		loadedTypes = [3]bool{
			aq.withPatient != nil,
			aq.withProvider != nil,
			aq.withDiagnoses != nil,
		}
	)
	if aq.withPatient != nil || aq.withProvider != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, appointment.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Appointment).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Appointment{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withPatient; query != nil {
		if err := aq.loadPatient(ctx, query, nodes, nil,
			func(n *Appointment, e *Patient) { n.Edges.Patient = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withProvider; query != nil {
		if err := aq.loadProvider(ctx, query, nodes, nil,
			func(n *Appointment, e *Provider) { n.Edges.Provider = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withDiagnoses; query != nil {
		if err := aq.loadDiagnoses(ctx, query, nodes,
			func(n *Appointment) { n.Edges.Diagnoses = []*Diagnosis{} },
			func(n *Appointment, e *Diagnosis) { n.Edges.Diagnoses = append(n.Edges.Diagnoses, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AppointmentQuery) loadPatient(ctx context.Context, query *PatientQuery, nodes []*Appointment, init func(*Appointment), assign func(*Appointment, *Patient)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Appointment)
	for i := range nodes {
		if nodes[i].patient_appointments == nil {
			continue
		}
		fk := *nodes[i].patient_appointments
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(patient.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "patient_appointments" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AppointmentQuery) loadProvider(ctx context.Context, query *ProviderQuery, nodes []*Appointment, init func(*Appointment), assign func(*Appointment, *Provider)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Appointment)
	for i := range nodes {
		if nodes[i].provider_appointments == nil {
			continue
		}
		fk := *nodes[i].provider_appointments
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(provider.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "provider_appointments" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AppointmentQuery) loadDiagnoses(ctx context.Context, query *DiagnosisQuery, nodes []*Appointment, init func(*Appointment), assign func(*Appointment, *Diagnosis)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Appointment)
	nids := make(map[uuid.UUID]map[*Appointment]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(appointment.DiagnosesTable)
		s.Join(joinT).On(s.C(diagnosis.FieldID), joinT.C(appointment.DiagnosesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(appointment.DiagnosesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(appointment.DiagnosesPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Appointment]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Diagnosis](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "diagnoses" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (aq *AppointmentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AppointmentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(appointment.Table, appointment.Columns, sqlgraph.NewFieldSpec(appointment.FieldID, field.TypeUUID))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appointment.FieldID)
		for i := range fields {
			if fields[i] != appointment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AppointmentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(appointment.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = appointment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AppointmentGroupBy is the group-by builder for Appointment entities.
type AppointmentGroupBy struct {
	selector
	build *AppointmentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AppointmentGroupBy) Aggregate(fns ...AggregateFunc) *AppointmentGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *AppointmentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, "GroupBy")
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppointmentQuery, *AppointmentGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *AppointmentGroupBy) sqlScan(ctx context.Context, root *AppointmentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AppointmentSelect is the builder for selecting fields of Appointment entities.
type AppointmentSelect struct {
	*AppointmentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AppointmentSelect) Aggregate(fns ...AggregateFunc) *AppointmentSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AppointmentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, "Select")
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppointmentQuery, *AppointmentSelect](ctx, as.AppointmentQuery, as, as.inters, v)
}

func (as *AppointmentSelect) sqlScan(ctx context.Context, root *AppointmentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
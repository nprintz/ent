// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/file"
	"entgo.io/ent/entc/integration/ent/group"
	"entgo.io/ent/entc/integration/ent/groupinfo"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/entc/integration/ent/user"
	"entgo.io/ent/schema/field"
)

// GroupQuery is the builder for querying Group entities.
type GroupQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Group
	// eager-loading edges.
	withFiles   *FileQuery
	withBlocked *UserQuery
	withUsers   *UserQuery
	withInfo    *GroupInfoQuery
	withFKs     bool
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GroupQuery builder.
func (gq *GroupQuery) Where(ps ...predicate.Group) *GroupQuery {
	gq.predicates = append(gq.predicates, ps...)
	return gq
}

// Limit adds a limit step to the query.
func (gq *GroupQuery) Limit(limit int) *GroupQuery {
	gq.limit = &limit
	return gq
}

// Offset adds an offset step to the query.
func (gq *GroupQuery) Offset(offset int) *GroupQuery {
	gq.offset = &offset
	return gq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gq *GroupQuery) Unique(unique bool) *GroupQuery {
	gq.unique = &unique
	return gq
}

// Order adds an order step to the query.
func (gq *GroupQuery) Order(o ...OrderFunc) *GroupQuery {
	gq.order = append(gq.order, o...)
	return gq
}

// QueryFiles chains the current query on the "files" edge.
func (gq *GroupQuery) QueryFiles() *FileQuery {
	query := &FileQuery{config: gq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, selector),
			sqlgraph.To(file.Table, file.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, group.FilesTable, group.FilesColumn),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBlocked chains the current query on the "blocked" edge.
func (gq *GroupQuery) QueryBlocked() *UserQuery {
	query := &UserQuery{config: gq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, group.BlockedTable, group.BlockedColumn),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUsers chains the current query on the "users" edge.
func (gq *GroupQuery) QueryUsers() *UserQuery {
	query := &UserQuery{config: gq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, group.UsersTable, group.UsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInfo chains the current query on the "info" edge.
func (gq *GroupQuery) QueryInfo() *GroupInfoQuery {
	query := &GroupInfoQuery{config: gq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, selector),
			sqlgraph.To(groupinfo.Table, groupinfo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, group.InfoTable, group.InfoColumn),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Group entity from the query.
// Returns a *NotFoundError when no Group was found.
func (gq *GroupQuery) First(ctx context.Context) (*Group, error) {
	nodes, err := gq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{group.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gq *GroupQuery) FirstX(ctx context.Context) *Group {
	node, err := gq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Group ID from the query.
// Returns a *NotFoundError when no Group ID was found.
func (gq *GroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{group.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gq *GroupQuery) FirstIDX(ctx context.Context) int {
	id, err := gq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Group entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Group entity is found.
// Returns a *NotFoundError when no Group entities are found.
func (gq *GroupQuery) Only(ctx context.Context) (*Group, error) {
	nodes, err := gq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{group.Label}
	default:
		return nil, &NotSingularError{group.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gq *GroupQuery) OnlyX(ctx context.Context) *Group {
	node, err := gq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Group ID in the query.
// Returns a *NotSingularError when more than one Group ID is found.
// Returns a *NotFoundError when no entities are found.
func (gq *GroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{group.Label}
	default:
		err = &NotSingularError{group.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gq *GroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := gq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Groups.
func (gq *GroupQuery) All(ctx context.Context) ([]*Group, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gq *GroupQuery) AllX(ctx context.Context) []*Group {
	nodes, err := gq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Group IDs.
func (gq *GroupQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := gq.Select(group.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gq *GroupQuery) IDsX(ctx context.Context) []int {
	ids, err := gq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gq *GroupQuery) Count(ctx context.Context) (int, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gq *GroupQuery) CountX(ctx context.Context) int {
	count, err := gq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gq *GroupQuery) Exist(ctx context.Context) (bool, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gq *GroupQuery) ExistX(ctx context.Context) bool {
	exist, err := gq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gq *GroupQuery) Clone() *GroupQuery {
	if gq == nil {
		return nil
	}
	return &GroupQuery{
		config:      gq.config,
		limit:       gq.limit,
		offset:      gq.offset,
		order:       append([]OrderFunc{}, gq.order...),
		predicates:  append([]predicate.Group{}, gq.predicates...),
		withFiles:   gq.withFiles.Clone(),
		withBlocked: gq.withBlocked.Clone(),
		withUsers:   gq.withUsers.Clone(),
		withInfo:    gq.withInfo.Clone(),
		// clone intermediate query.
		sql:    gq.sql.Clone(),
		path:   gq.path,
		unique: gq.unique,
	}
}

// WithFiles tells the query-builder to eager-load the nodes that are connected to
// the "files" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GroupQuery) WithFiles(opts ...func(*FileQuery)) *GroupQuery {
	query := &FileQuery{config: gq.config}
	for _, opt := range opts {
		opt(query)
	}
	gq.withFiles = query
	return gq
}

// WithBlocked tells the query-builder to eager-load the nodes that are connected to
// the "blocked" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GroupQuery) WithBlocked(opts ...func(*UserQuery)) *GroupQuery {
	query := &UserQuery{config: gq.config}
	for _, opt := range opts {
		opt(query)
	}
	gq.withBlocked = query
	return gq
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GroupQuery) WithUsers(opts ...func(*UserQuery)) *GroupQuery {
	query := &UserQuery{config: gq.config}
	for _, opt := range opts {
		opt(query)
	}
	gq.withUsers = query
	return gq
}

// WithInfo tells the query-builder to eager-load the nodes that are connected to
// the "info" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GroupQuery) WithInfo(opts ...func(*GroupInfoQuery)) *GroupQuery {
	query := &GroupInfoQuery{config: gq.config}
	for _, opt := range opts {
		opt(query)
	}
	gq.withInfo = query
	return gq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Active bool `json:"active,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Group.Query().
//		GroupBy(group.FieldActive).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (gq *GroupQuery) GroupBy(field string, fields ...string) *GroupGroupBy {
	grbuild := &GroupGroupBy{config: gq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gq.sqlQuery(ctx), nil
	}
	grbuild.label = group.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Active bool `json:"active,omitempty"`
//	}
//
//	client.Group.Query().
//		Select(group.FieldActive).
//		Scan(ctx, &v)
//
func (gq *GroupQuery) Select(fields ...string) *GroupSelect {
	gq.fields = append(gq.fields, fields...)
	selbuild := &GroupSelect{GroupQuery: gq}
	selbuild.label = group.Label
	selbuild.flds, selbuild.scan = &gq.fields, selbuild.Scan
	return selbuild
}

func (gq *GroupQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gq.fields {
		if !group.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gq.path != nil {
		prev, err := gq.path(ctx)
		if err != nil {
			return err
		}
		gq.sql = prev
	}
	return nil
}

func (gq *GroupQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Group, error) {
	var (
		nodes       = []*Group{}
		withFKs     = gq.withFKs
		_spec       = gq.querySpec()
		loadedTypes = [4]bool{
			gq.withFiles != nil,
			gq.withBlocked != nil,
			gq.withUsers != nil,
			gq.withInfo != nil,
		}
	)
	if gq.withInfo != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, group.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Group).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Group{config: gq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(gq.modifiers) > 0 {
		_spec.Modifiers = gq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := gq.withFiles; query != nil {
		if err := gq.loadFiles(ctx, query, nodes,
			func(n *Group) { n.Edges.Files = []*File{} },
			func(n *Group, e *File) { n.Edges.Files = append(n.Edges.Files, e) }); err != nil {
			return nil, err
		}
	}
	if query := gq.withBlocked; query != nil {
		if err := gq.loadBlocked(ctx, query, nodes,
			func(n *Group) { n.Edges.Blocked = []*User{} },
			func(n *Group, e *User) { n.Edges.Blocked = append(n.Edges.Blocked, e) }); err != nil {
			return nil, err
		}
	}
	if query := gq.withUsers; query != nil {
		if err := gq.loadUsers(ctx, query, nodes,
			func(n *Group) { n.Edges.Users = []*User{} },
			func(n *Group, e *User) { n.Edges.Users = append(n.Edges.Users, e) }); err != nil {
			return nil, err
		}
	}
	if query := gq.withInfo; query != nil {
		if err := gq.loadInfo(ctx, query, nodes, nil,
			func(n *Group, e *GroupInfo) { n.Edges.Info = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (gq *GroupQuery) loadFiles(ctx context.Context, query *FileQuery, nodes []*Group, init func(*Group), assign func(*Group, *File)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Group)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.File(func(s *sql.Selector) {
		s.Where(sql.InValues(group.FilesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.group_files
		if fk == nil {
			return fmt.Errorf(`foreign-key "group_files" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "group_files" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (gq *GroupQuery) loadBlocked(ctx context.Context, query *UserQuery, nodes []*Group, init func(*Group), assign func(*Group, *User)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Group)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.User(func(s *sql.Selector) {
		s.Where(sql.InValues(group.BlockedColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.group_blocked
		if fk == nil {
			return fmt.Errorf(`foreign-key "group_blocked" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "group_blocked" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (gq *GroupQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*Group, init func(*Group), assign func(*Group, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Group)
	nids := make(map[int]map[*Group]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(group.UsersTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(group.UsersPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(group.UsersPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(group.UsersPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]interface{}, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]interface{}{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []interface{}) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*Group]struct{}{byID[outValue]: struct{}{}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "users" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (gq *GroupQuery) loadInfo(ctx context.Context, query *GroupInfoQuery, nodes []*Group, init func(*Group), assign func(*Group, *GroupInfo)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Group)
	for i := range nodes {
		if nodes[i].group_info == nil {
			continue
		}
		fk := *nodes[i].group_info
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(groupinfo.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "group_info" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (gq *GroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gq.querySpec()
	if len(gq.modifiers) > 0 {
		_spec.Modifiers = gq.modifiers
	}
	_spec.Node.Columns = gq.fields
	if len(gq.fields) > 0 {
		_spec.Unique = gq.unique != nil && *gq.unique
	}
	return sqlgraph.CountNodes(ctx, gq.driver, _spec)
}

func (gq *GroupQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (gq *GroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   group.Table,
			Columns: group.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: group.FieldID,
			},
		},
		From:   gq.sql,
		Unique: true,
	}
	if unique := gq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, group.FieldID)
		for i := range fields {
			if fields[i] != group.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gq *GroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gq.driver.Dialect())
	t1 := builder.Table(group.Table)
	columns := gq.fields
	if len(columns) == 0 {
		columns = group.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gq.sql != nil {
		selector = gq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gq.unique != nil && *gq.unique {
		selector.Distinct()
	}
	for _, m := range gq.modifiers {
		m(selector)
	}
	for _, p := range gq.predicates {
		p(selector)
	}
	for _, p := range gq.order {
		p(selector)
	}
	if offset := gq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (gq *GroupQuery) ForUpdate(opts ...sql.LockOption) *GroupQuery {
	if gq.driver.Dialect() == dialect.Postgres {
		gq.Unique(false)
	}
	gq.modifiers = append(gq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return gq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (gq *GroupQuery) ForShare(opts ...sql.LockOption) *GroupQuery {
	if gq.driver.Dialect() == dialect.Postgres {
		gq.Unique(false)
	}
	gq.modifiers = append(gq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return gq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gq *GroupQuery) Modify(modifiers ...func(s *sql.Selector)) *GroupSelect {
	gq.modifiers = append(gq.modifiers, modifiers...)
	return gq.Select()
}

// GroupGroupBy is the group-by builder for Group entities.
type GroupGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ggb *GroupGroupBy) Aggregate(fns ...AggregateFunc) *GroupGroupBy {
	ggb.fns = append(ggb.fns, fns...)
	return ggb
}

// Scan applies the group-by query and scans the result into the given value.
func (ggb *GroupGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ggb.path(ctx)
	if err != nil {
		return err
	}
	ggb.sql = query
	return ggb.sqlScan(ctx, v)
}

func (ggb *GroupGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ggb.fields {
		if !group.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ggb *GroupGroupBy) sqlQuery() *sql.Selector {
	selector := ggb.sql.Select()
	aggregation := make([]string, 0, len(ggb.fns))
	for _, fn := range ggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ggb.fields)+len(ggb.fns))
		for _, f := range ggb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ggb.fields...)...)
}

// GroupSelect is the builder for selecting fields of Group entities.
type GroupSelect struct {
	*GroupQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gs *GroupSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gs.prepareQuery(ctx); err != nil {
		return err
	}
	gs.sql = gs.GroupQuery.sqlQuery(ctx)
	return gs.sqlScan(ctx, v)
}

func (gs *GroupSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gs.sql.Query()
	if err := gs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gs *GroupSelect) Modify(modifiers ...func(s *sql.Selector)) *GroupSelect {
	gs.modifiers = append(gs.modifiers, modifiers...)
	return gs
}

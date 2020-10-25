// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tew147258/app/ent/confirmation"
	"github.com/tew147258/app/ent/predicate"
	"github.com/tew147258/app/ent/stadium"
)

// StadiumUpdate is the builder for updating Stadium entities.
type StadiumUpdate struct {
	config
	hooks      []Hook
	mutation   *StadiumMutation
	predicates []predicate.Stadium
}

// Where adds a new predicate for the builder.
func (su *StadiumUpdate) Where(ps ...predicate.Stadium) *StadiumUpdate {
	su.predicates = append(su.predicates, ps...)
	return su
}

// SetNamestadium sets the namestadium field.
func (su *StadiumUpdate) SetNamestadium(s string) *StadiumUpdate {
	su.mutation.SetNamestadium(s)
	return su
}

// AddStadiumConfirmationIDs adds the StadiumConfirmation edge to Confirmation by ids.
func (su *StadiumUpdate) AddStadiumConfirmationIDs(ids ...int) *StadiumUpdate {
	su.mutation.AddStadiumConfirmationIDs(ids...)
	return su
}

// AddStadiumConfirmation adds the StadiumConfirmation edges to Confirmation.
func (su *StadiumUpdate) AddStadiumConfirmation(c ...*Confirmation) *StadiumUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.AddStadiumConfirmationIDs(ids...)
}

// Mutation returns the StadiumMutation object of the builder.
func (su *StadiumUpdate) Mutation() *StadiumMutation {
	return su.mutation
}

// RemoveStadiumConfirmationIDs removes the StadiumConfirmation edge to Confirmation by ids.
func (su *StadiumUpdate) RemoveStadiumConfirmationIDs(ids ...int) *StadiumUpdate {
	su.mutation.RemoveStadiumConfirmationIDs(ids...)
	return su
}

// RemoveStadiumConfirmation removes StadiumConfirmation edges to Confirmation.
func (su *StadiumUpdate) RemoveStadiumConfirmation(c ...*Confirmation) *StadiumUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.RemoveStadiumConfirmationIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (su *StadiumUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := su.mutation.Namestadium(); ok {
		if err := stadium.NamestadiumValidator(v); err != nil {
			return 0, &ValidationError{Name: "namestadium", err: fmt.Errorf("ent: validator failed for field \"namestadium\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StadiumMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *StadiumUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StadiumUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StadiumUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *StadiumUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   stadium.Table,
			Columns: stadium.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: stadium.FieldID,
			},
		},
	}
	if ps := su.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Namestadium(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: stadium.FieldNamestadium,
		})
	}
	if nodes := su.mutation.RemovedStadiumConfirmationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   stadium.StadiumConfirmationTable,
			Columns: []string{stadium.StadiumConfirmationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: confirmation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.StadiumConfirmationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   stadium.StadiumConfirmationTable,
			Columns: []string{stadium.StadiumConfirmationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: confirmation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{stadium.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// StadiumUpdateOne is the builder for updating a single Stadium entity.
type StadiumUpdateOne struct {
	config
	hooks    []Hook
	mutation *StadiumMutation
}

// SetNamestadium sets the namestadium field.
func (suo *StadiumUpdateOne) SetNamestadium(s string) *StadiumUpdateOne {
	suo.mutation.SetNamestadium(s)
	return suo
}

// AddStadiumConfirmationIDs adds the StadiumConfirmation edge to Confirmation by ids.
func (suo *StadiumUpdateOne) AddStadiumConfirmationIDs(ids ...int) *StadiumUpdateOne {
	suo.mutation.AddStadiumConfirmationIDs(ids...)
	return suo
}

// AddStadiumConfirmation adds the StadiumConfirmation edges to Confirmation.
func (suo *StadiumUpdateOne) AddStadiumConfirmation(c ...*Confirmation) *StadiumUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.AddStadiumConfirmationIDs(ids...)
}

// Mutation returns the StadiumMutation object of the builder.
func (suo *StadiumUpdateOne) Mutation() *StadiumMutation {
	return suo.mutation
}

// RemoveStadiumConfirmationIDs removes the StadiumConfirmation edge to Confirmation by ids.
func (suo *StadiumUpdateOne) RemoveStadiumConfirmationIDs(ids ...int) *StadiumUpdateOne {
	suo.mutation.RemoveStadiumConfirmationIDs(ids...)
	return suo
}

// RemoveStadiumConfirmation removes StadiumConfirmation edges to Confirmation.
func (suo *StadiumUpdateOne) RemoveStadiumConfirmation(c ...*Confirmation) *StadiumUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.RemoveStadiumConfirmationIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (suo *StadiumUpdateOne) Save(ctx context.Context) (*Stadium, error) {
	if v, ok := suo.mutation.Namestadium(); ok {
		if err := stadium.NamestadiumValidator(v); err != nil {
			return nil, &ValidationError{Name: "namestadium", err: fmt.Errorf("ent: validator failed for field \"namestadium\": %w", err)}
		}
	}

	var (
		err  error
		node *Stadium
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StadiumMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StadiumUpdateOne) SaveX(ctx context.Context) *Stadium {
	s, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// Exec executes the query on the entity.
func (suo *StadiumUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StadiumUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *StadiumUpdateOne) sqlSave(ctx context.Context) (s *Stadium, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   stadium.Table,
			Columns: stadium.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: stadium.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Stadium.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := suo.mutation.Namestadium(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: stadium.FieldNamestadium,
		})
	}
	if nodes := suo.mutation.RemovedStadiumConfirmationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   stadium.StadiumConfirmationTable,
			Columns: []string{stadium.StadiumConfirmationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: confirmation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.StadiumConfirmationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   stadium.StadiumConfirmationTable,
			Columns: []string{stadium.StadiumConfirmationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: confirmation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	s = &Stadium{config: suo.config}
	_spec.Assign = s.assignValues
	_spec.ScanValues = s.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{stadium.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}
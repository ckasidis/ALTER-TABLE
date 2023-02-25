// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mcbebu/ALTER-TABLE/ent/order"
	"github.com/mcbebu/ALTER-TABLE/ent/schema"
	"github.com/mcbebu/ALTER-TABLE/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetMobileNumber sets the "mobileNumber" field.
func (uc *UserCreate) SetMobileNumber(s string) *UserCreate {
	uc.mutation.SetMobileNumber(s)
	return uc
}

// SetAddresses sets the "addresses" field.
func (uc *UserCreate) SetAddresses(s []schema.Address) *UserCreate {
	uc.mutation.SetAddresses(s)
	return uc
}

// SetLeaveParcel sets the "leaveParcel" field.
func (uc *UserCreate) SetLeaveParcel(b bool) *UserCreate {
	uc.mutation.SetLeaveParcel(b)
	return uc
}

// SetNillableLeaveParcel sets the "leaveParcel" field if the given value is not nil.
func (uc *UserCreate) SetNillableLeaveParcel(b *bool) *UserCreate {
	if b != nil {
		uc.SetLeaveParcel(*b)
	}
	return uc
}

// SetInstructions sets the "instructions" field.
func (uc *UserCreate) SetInstructions(s []string) *UserCreate {
	uc.mutation.SetInstructions(s)
	return uc
}

// SetNotifications sets the "notifications" field.
func (uc *UserCreate) SetNotifications(b [4]bool) *UserCreate {
	uc.mutation.SetNotifications(b)
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(s string) *UserCreate {
	uc.mutation.SetID(s)
	return uc
}

// AddOrderIDs adds the "orders" edge to the Order entity by IDs.
func (uc *UserCreate) AddOrderIDs(ids ...int) *UserCreate {
	uc.mutation.AddOrderIDs(ids...)
	return uc
}

// AddOrders adds the "orders" edges to the Order entity.
func (uc *UserCreate) AddOrders(o ...*Order) *UserCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return uc.AddOrderIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks[*User, UserMutation](ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.Addresses(); !ok {
		v := user.DefaultAddresses
		uc.mutation.SetAddresses(v)
	}
	if _, ok := uc.mutation.LeaveParcel(); !ok {
		v := user.DefaultLeaveParcel
		uc.mutation.SetLeaveParcel(v)
	}
	if _, ok := uc.mutation.Instructions(); !ok {
		v := user.DefaultInstructions
		uc.mutation.SetInstructions(v)
	}
	if _, ok := uc.mutation.Notifications(); !ok {
		v := user.DefaultNotifications
		uc.mutation.SetNotifications(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.MobileNumber(); !ok {
		return &ValidationError{Name: "mobileNumber", err: errors.New(`ent: missing required field "User.mobileNumber"`)}
	}
	if v, ok := uc.mutation.MobileNumber(); ok {
		if err := user.MobileNumberValidator(v); err != nil {
			return &ValidationError{Name: "mobileNumber", err: fmt.Errorf(`ent: validator failed for field "User.mobileNumber": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Addresses(); !ok {
		return &ValidationError{Name: "addresses", err: errors.New(`ent: missing required field "User.addresses"`)}
	}
	if _, ok := uc.mutation.LeaveParcel(); !ok {
		return &ValidationError{Name: "leaveParcel", err: errors.New(`ent: missing required field "User.leaveParcel"`)}
	}
	if _, ok := uc.mutation.Instructions(); !ok {
		return &ValidationError{Name: "instructions", err: errors.New(`ent: missing required field "User.instructions"`)}
	}
	if _, ok := uc.mutation.Notifications(); !ok {
		return &ValidationError{Name: "notifications", err: errors.New(`ent: missing required field "User.notifications"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected User.ID type: %T", _spec.ID.Value)
		}
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.MobileNumber(); ok {
		_spec.SetField(user.FieldMobileNumber, field.TypeString, value)
		_node.MobileNumber = value
	}
	if value, ok := uc.mutation.Addresses(); ok {
		_spec.SetField(user.FieldAddresses, field.TypeJSON, value)
		_node.Addresses = value
	}
	if value, ok := uc.mutation.LeaveParcel(); ok {
		_spec.SetField(user.FieldLeaveParcel, field.TypeBool, value)
		_node.LeaveParcel = value
	}
	if value, ok := uc.mutation.Instructions(); ok {
		_spec.SetField(user.FieldInstructions, field.TypeJSON, value)
		_node.Instructions = value
	}
	if value, ok := uc.mutation.Notifications(); ok {
		_spec.SetField(user.FieldNotifications, field.TypeJSON, value)
		_node.Notifications = value
	}
	if nodes := uc.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.OrdersTable,
			Columns: []string{user.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/tew147258/app/ent/migrate"

	"github.com/tew147258/app/ent/borrow"
	"github.com/tew147258/app/ent/confirmation"
	"github.com/tew147258/app/ent/stadium"
	"github.com/tew147258/app/ent/user"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Borrow is the client for interacting with the Borrow builders.
	Borrow *BorrowClient
	// Confirmation is the client for interacting with the Confirmation builders.
	Confirmation *ConfirmationClient
	// Stadium is the client for interacting with the Stadium builders.
	Stadium *StadiumClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Borrow = NewBorrowClient(c.config)
	c.Confirmation = NewConfirmationClient(c.config)
	c.Stadium = NewStadiumClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Borrow:       NewBorrowClient(cfg),
		Confirmation: NewConfirmationClient(cfg),
		Stadium:      NewStadiumClient(cfg),
		User:         NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:       cfg,
		Borrow:       NewBorrowClient(cfg),
		Confirmation: NewConfirmationClient(cfg),
		Stadium:      NewStadiumClient(cfg),
		User:         NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Borrow.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Borrow.Use(hooks...)
	c.Confirmation.Use(hooks...)
	c.Stadium.Use(hooks...)
	c.User.Use(hooks...)
}

// BorrowClient is a client for the Borrow schema.
type BorrowClient struct {
	config
}

// NewBorrowClient returns a client for the Borrow from the given config.
func NewBorrowClient(c config) *BorrowClient {
	return &BorrowClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `borrow.Hooks(f(g(h())))`.
func (c *BorrowClient) Use(hooks ...Hook) {
	c.hooks.Borrow = append(c.hooks.Borrow, hooks...)
}

// Create returns a create builder for Borrow.
func (c *BorrowClient) Create() *BorrowCreate {
	mutation := newBorrowMutation(c.config, OpCreate)
	return &BorrowCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Borrow.
func (c *BorrowClient) Update() *BorrowUpdate {
	mutation := newBorrowMutation(c.config, OpUpdate)
	return &BorrowUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BorrowClient) UpdateOne(b *Borrow) *BorrowUpdateOne {
	mutation := newBorrowMutation(c.config, OpUpdateOne, withBorrow(b))
	return &BorrowUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BorrowClient) UpdateOneID(id int) *BorrowUpdateOne {
	mutation := newBorrowMutation(c.config, OpUpdateOne, withBorrowID(id))
	return &BorrowUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Borrow.
func (c *BorrowClient) Delete() *BorrowDelete {
	mutation := newBorrowMutation(c.config, OpDelete)
	return &BorrowDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BorrowClient) DeleteOne(b *Borrow) *BorrowDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BorrowClient) DeleteOneID(id int) *BorrowDeleteOne {
	builder := c.Delete().Where(borrow.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BorrowDeleteOne{builder}
}

// Create returns a query builder for Borrow.
func (c *BorrowClient) Query() *BorrowQuery {
	return &BorrowQuery{config: c.config}
}

// Get returns a Borrow entity by its id.
func (c *BorrowClient) Get(ctx context.Context, id int) (*Borrow, error) {
	return c.Query().Where(borrow.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BorrowClient) GetX(ctx context.Context, id int) *Borrow {
	b, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return b
}

// QueryBorrowConfirmation queries the BorrowConfirmation edge of a Borrow.
func (c *BorrowClient) QueryBorrowConfirmation(b *Borrow) *ConfirmationQuery {
	query := &ConfirmationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(borrow.Table, borrow.FieldID, id),
			sqlgraph.To(confirmation.Table, confirmation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, borrow.BorrowConfirmationTable, borrow.BorrowConfirmationColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BorrowClient) Hooks() []Hook {
	return c.hooks.Borrow
}

// ConfirmationClient is a client for the Confirmation schema.
type ConfirmationClient struct {
	config
}

// NewConfirmationClient returns a client for the Confirmation from the given config.
func NewConfirmationClient(c config) *ConfirmationClient {
	return &ConfirmationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `confirmation.Hooks(f(g(h())))`.
func (c *ConfirmationClient) Use(hooks ...Hook) {
	c.hooks.Confirmation = append(c.hooks.Confirmation, hooks...)
}

// Create returns a create builder for Confirmation.
func (c *ConfirmationClient) Create() *ConfirmationCreate {
	mutation := newConfirmationMutation(c.config, OpCreate)
	return &ConfirmationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Confirmation.
func (c *ConfirmationClient) Update() *ConfirmationUpdate {
	mutation := newConfirmationMutation(c.config, OpUpdate)
	return &ConfirmationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ConfirmationClient) UpdateOne(co *Confirmation) *ConfirmationUpdateOne {
	mutation := newConfirmationMutation(c.config, OpUpdateOne, withConfirmation(co))
	return &ConfirmationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ConfirmationClient) UpdateOneID(id int) *ConfirmationUpdateOne {
	mutation := newConfirmationMutation(c.config, OpUpdateOne, withConfirmationID(id))
	return &ConfirmationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Confirmation.
func (c *ConfirmationClient) Delete() *ConfirmationDelete {
	mutation := newConfirmationMutation(c.config, OpDelete)
	return &ConfirmationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ConfirmationClient) DeleteOne(co *Confirmation) *ConfirmationDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ConfirmationClient) DeleteOneID(id int) *ConfirmationDeleteOne {
	builder := c.Delete().Where(confirmation.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ConfirmationDeleteOne{builder}
}

// Create returns a query builder for Confirmation.
func (c *ConfirmationClient) Query() *ConfirmationQuery {
	return &ConfirmationQuery{config: c.config}
}

// Get returns a Confirmation entity by its id.
func (c *ConfirmationClient) Get(ctx context.Context, id int) (*Confirmation, error) {
	return c.Query().Where(confirmation.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ConfirmationClient) GetX(ctx context.Context, id int) *Confirmation {
	co, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return co
}

// QueryConfirmationUser queries the ConfirmationUser edge of a Confirmation.
func (c *ConfirmationClient) QueryConfirmationUser(co *Confirmation) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(confirmation.Table, confirmation.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, confirmation.ConfirmationUserTable, confirmation.ConfirmationUserColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryConfirmationStadium queries the ConfirmationStadium edge of a Confirmation.
func (c *ConfirmationClient) QueryConfirmationStadium(co *Confirmation) *StadiumQuery {
	query := &StadiumQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(confirmation.Table, confirmation.FieldID, id),
			sqlgraph.To(stadium.Table, stadium.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, confirmation.ConfirmationStadiumTable, confirmation.ConfirmationStadiumColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryConfirmationBorrow queries the ConfirmationBorrow edge of a Confirmation.
func (c *ConfirmationClient) QueryConfirmationBorrow(co *Confirmation) *BorrowQuery {
	query := &BorrowQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(confirmation.Table, confirmation.FieldID, id),
			sqlgraph.To(borrow.Table, borrow.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, confirmation.ConfirmationBorrowTable, confirmation.ConfirmationBorrowColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ConfirmationClient) Hooks() []Hook {
	return c.hooks.Confirmation
}

// StadiumClient is a client for the Stadium schema.
type StadiumClient struct {
	config
}

// NewStadiumClient returns a client for the Stadium from the given config.
func NewStadiumClient(c config) *StadiumClient {
	return &StadiumClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `stadium.Hooks(f(g(h())))`.
func (c *StadiumClient) Use(hooks ...Hook) {
	c.hooks.Stadium = append(c.hooks.Stadium, hooks...)
}

// Create returns a create builder for Stadium.
func (c *StadiumClient) Create() *StadiumCreate {
	mutation := newStadiumMutation(c.config, OpCreate)
	return &StadiumCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Stadium.
func (c *StadiumClient) Update() *StadiumUpdate {
	mutation := newStadiumMutation(c.config, OpUpdate)
	return &StadiumUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StadiumClient) UpdateOne(s *Stadium) *StadiumUpdateOne {
	mutation := newStadiumMutation(c.config, OpUpdateOne, withStadium(s))
	return &StadiumUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StadiumClient) UpdateOneID(id int) *StadiumUpdateOne {
	mutation := newStadiumMutation(c.config, OpUpdateOne, withStadiumID(id))
	return &StadiumUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Stadium.
func (c *StadiumClient) Delete() *StadiumDelete {
	mutation := newStadiumMutation(c.config, OpDelete)
	return &StadiumDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *StadiumClient) DeleteOne(s *Stadium) *StadiumDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *StadiumClient) DeleteOneID(id int) *StadiumDeleteOne {
	builder := c.Delete().Where(stadium.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StadiumDeleteOne{builder}
}

// Create returns a query builder for Stadium.
func (c *StadiumClient) Query() *StadiumQuery {
	return &StadiumQuery{config: c.config}
}

// Get returns a Stadium entity by its id.
func (c *StadiumClient) Get(ctx context.Context, id int) (*Stadium, error) {
	return c.Query().Where(stadium.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StadiumClient) GetX(ctx context.Context, id int) *Stadium {
	s, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return s
}

// QueryStadiumConfirmation queries the StadiumConfirmation edge of a Stadium.
func (c *StadiumClient) QueryStadiumConfirmation(s *Stadium) *ConfirmationQuery {
	query := &ConfirmationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(stadium.Table, stadium.FieldID, id),
			sqlgraph.To(confirmation.Table, confirmation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, stadium.StadiumConfirmationTable, stadium.StadiumConfirmationColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StadiumClient) Hooks() []Hook {
	return c.hooks.Stadium
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Create returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QueryUserConfirmation queries the UserConfirmation edge of a User.
func (c *UserClient) QueryUserConfirmation(u *User) *ConfirmationQuery {
	query := &ConfirmationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(confirmation.Table, confirmation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.UserConfirmationTable, user.UserConfirmationColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

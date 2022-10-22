// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"entrepro/ent/migrate"

	"entrepro/ent/domain"
	"entrepro/ent/marketplaceitem"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Domain is the client for interacting with the Domain builders.
	Domain *DomainClient
	// Marketplaceitem is the client for interacting with the Marketplaceitem builders.
	Marketplaceitem *MarketplaceitemClient
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
	c.Domain = NewDomainClient(c.config)
	c.Marketplaceitem = NewMarketplaceitemClient(c.config)
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
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:             ctx,
		config:          cfg,
		Domain:          NewDomainClient(cfg),
		Marketplaceitem: NewMarketplaceitemClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:             ctx,
		config:          cfg,
		Domain:          NewDomainClient(cfg),
		Marketplaceitem: NewMarketplaceitemClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Domain.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
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
	c.Domain.Use(hooks...)
	c.Marketplaceitem.Use(hooks...)
}

// DomainClient is a client for the Domain schema.
type DomainClient struct {
	config
}

// NewDomainClient returns a client for the Domain from the given config.
func NewDomainClient(c config) *DomainClient {
	return &DomainClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `domain.Hooks(f(g(h())))`.
func (c *DomainClient) Use(hooks ...Hook) {
	c.hooks.Domain = append(c.hooks.Domain, hooks...)
}

// Create returns a builder for creating a Domain entity.
func (c *DomainClient) Create() *DomainCreate {
	mutation := newDomainMutation(c.config, OpCreate)
	return &DomainCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Domain entities.
func (c *DomainClient) CreateBulk(builders ...*DomainCreate) *DomainCreateBulk {
	return &DomainCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Domain.
func (c *DomainClient) Update() *DomainUpdate {
	mutation := newDomainMutation(c.config, OpUpdate)
	return &DomainUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DomainClient) UpdateOne(d *Domain) *DomainUpdateOne {
	mutation := newDomainMutation(c.config, OpUpdateOne, withDomain(d))
	return &DomainUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DomainClient) UpdateOneID(id int) *DomainUpdateOne {
	mutation := newDomainMutation(c.config, OpUpdateOne, withDomainID(id))
	return &DomainUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Domain.
func (c *DomainClient) Delete() *DomainDelete {
	mutation := newDomainMutation(c.config, OpDelete)
	return &DomainDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DomainClient) DeleteOne(d *Domain) *DomainDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *DomainClient) DeleteOneID(id int) *DomainDeleteOne {
	builder := c.Delete().Where(domain.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DomainDeleteOne{builder}
}

// Query returns a query builder for Domain.
func (c *DomainClient) Query() *DomainQuery {
	return &DomainQuery{
		config: c.config,
	}
}

// Get returns a Domain entity by its id.
func (c *DomainClient) Get(ctx context.Context, id int) (*Domain, error) {
	return c.Query().Where(domain.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DomainClient) GetX(ctx context.Context, id int) *Domain {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *DomainClient) Hooks() []Hook {
	return c.hooks.Domain
}

// MarketplaceitemClient is a client for the Marketplaceitem schema.
type MarketplaceitemClient struct {
	config
}

// NewMarketplaceitemClient returns a client for the Marketplaceitem from the given config.
func NewMarketplaceitemClient(c config) *MarketplaceitemClient {
	return &MarketplaceitemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `marketplaceitem.Hooks(f(g(h())))`.
func (c *MarketplaceitemClient) Use(hooks ...Hook) {
	c.hooks.Marketplaceitem = append(c.hooks.Marketplaceitem, hooks...)
}

// Create returns a builder for creating a Marketplaceitem entity.
func (c *MarketplaceitemClient) Create() *MarketplaceitemCreate {
	mutation := newMarketplaceitemMutation(c.config, OpCreate)
	return &MarketplaceitemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Marketplaceitem entities.
func (c *MarketplaceitemClient) CreateBulk(builders ...*MarketplaceitemCreate) *MarketplaceitemCreateBulk {
	return &MarketplaceitemCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Marketplaceitem.
func (c *MarketplaceitemClient) Update() *MarketplaceitemUpdate {
	mutation := newMarketplaceitemMutation(c.config, OpUpdate)
	return &MarketplaceitemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MarketplaceitemClient) UpdateOne(m *Marketplaceitem) *MarketplaceitemUpdateOne {
	mutation := newMarketplaceitemMutation(c.config, OpUpdateOne, withMarketplaceitem(m))
	return &MarketplaceitemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MarketplaceitemClient) UpdateOneID(id int) *MarketplaceitemUpdateOne {
	mutation := newMarketplaceitemMutation(c.config, OpUpdateOne, withMarketplaceitemID(id))
	return &MarketplaceitemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Marketplaceitem.
func (c *MarketplaceitemClient) Delete() *MarketplaceitemDelete {
	mutation := newMarketplaceitemMutation(c.config, OpDelete)
	return &MarketplaceitemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MarketplaceitemClient) DeleteOne(m *Marketplaceitem) *MarketplaceitemDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *MarketplaceitemClient) DeleteOneID(id int) *MarketplaceitemDeleteOne {
	builder := c.Delete().Where(marketplaceitem.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MarketplaceitemDeleteOne{builder}
}

// Query returns a query builder for Marketplaceitem.
func (c *MarketplaceitemClient) Query() *MarketplaceitemQuery {
	return &MarketplaceitemQuery{
		config: c.config,
	}
}

// Get returns a Marketplaceitem entity by its id.
func (c *MarketplaceitemClient) Get(ctx context.Context, id int) (*Marketplaceitem, error) {
	return c.Query().Where(marketplaceitem.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MarketplaceitemClient) GetX(ctx context.Context, id int) *Marketplaceitem {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MarketplaceitemClient) Hooks() []Hook {
	return c.hooks.Marketplaceitem
}

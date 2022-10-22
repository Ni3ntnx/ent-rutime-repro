// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entrepro/ent/domain"
	"entrepro/ent/marketplaceitem"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 2)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   domain.Table,
			Columns: domain.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: domain.FieldID,
			},
		},
		Type: "Domain",
		Fields: map[string]*sqlgraph.FieldSpec{
			domain.FieldTenantID: {Type: field.TypeString, Column: domain.FieldTenantID},
			domain.FieldName:     {Type: field.TypeString, Column: domain.FieldName},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   marketplaceitem.Table,
			Columns: marketplaceitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: marketplaceitem.FieldID,
			},
		},
		Type: "Marketplaceitem",
		Fields: map[string]*sqlgraph.FieldSpec{
			marketplaceitem.FieldDescription: {Type: field.TypeString, Column: marketplaceitem.FieldDescription},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (dq *DomainQuery) addPredicate(pred func(s *sql.Selector)) {
	dq.predicates = append(dq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the DomainQuery builder.
func (dq *DomainQuery) Filter() *DomainFilter {
	return &DomainFilter{config: dq.config, predicateAdder: dq}
}

// addPredicate implements the predicateAdder interface.
func (m *DomainMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the DomainMutation builder.
func (m *DomainMutation) Filter() *DomainFilter {
	return &DomainFilter{config: m.config, predicateAdder: m}
}

// DomainFilter provides a generic filtering capability at runtime for DomainQuery.
type DomainFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *DomainFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *DomainFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(domain.FieldID))
}

// WhereTenantID applies the entql string predicate on the tenant_id field.
func (f *DomainFilter) WhereTenantID(p entql.StringP) {
	f.Where(p.Field(domain.FieldTenantID))
}

// WhereName applies the entql string predicate on the name field.
func (f *DomainFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(domain.FieldName))
}

// addPredicate implements the predicateAdder interface.
func (mq *MarketplaceitemQuery) addPredicate(pred func(s *sql.Selector)) {
	mq.predicates = append(mq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the MarketplaceitemQuery builder.
func (mq *MarketplaceitemQuery) Filter() *MarketplaceitemFilter {
	return &MarketplaceitemFilter{config: mq.config, predicateAdder: mq}
}

// addPredicate implements the predicateAdder interface.
func (m *MarketplaceitemMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the MarketplaceitemMutation builder.
func (m *MarketplaceitemMutation) Filter() *MarketplaceitemFilter {
	return &MarketplaceitemFilter{config: m.config, predicateAdder: m}
}

// MarketplaceitemFilter provides a generic filtering capability at runtime for MarketplaceitemQuery.
type MarketplaceitemFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *MarketplaceitemFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *MarketplaceitemFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(marketplaceitem.FieldID))
}

// WhereDescription applies the entql string predicate on the description field.
func (f *MarketplaceitemFilter) WhereDescription(p entql.StringP) {
	f.Where(p.Field(marketplaceitem.FieldDescription))
}

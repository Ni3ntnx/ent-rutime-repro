package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Domain holds the schema definition for the Domain entity.
type Domain struct {
	ent.Schema
}

func (Domain) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TenantMixin{},
	}
}

// Fields of the Domain.
func (Domain) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

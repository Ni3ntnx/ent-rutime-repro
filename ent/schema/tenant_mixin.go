package schema

import (
	"entgo.io/ent"
	//"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"entrepro/ent/privacy"
    "entrepro/ent/rule"
)

// TenantMixin for embedding the tenant info in different schemas.
type TenantMixin struct {
	mixin.Schema
}

// Fields for all schemas that embed TenantMixin.
func (TenantMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("tenant_id").
			Immutable(),
	}
}

// Policy of the Tenant.
func (TenantMixin) Policy() ent.Policy {
	//return FilterTenantRule()
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rule.FilterTenantRule(),
		},
		Mutation: privacy.MutationPolicy{
			rule.FilterTenantRule(),
		},
	}
}

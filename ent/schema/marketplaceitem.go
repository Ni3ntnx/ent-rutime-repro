package schema

import (
  "entgo.io/ent"
  "entgo.io/ent/schema/field"
)

// Marketplaceitem holds the schema definition for the marketplaceitem entity.
type Marketplaceitem struct {
  ent.Schema
}

// Fields of the marketplaceitem.
func (Marketplaceitem) Fields() []ent.Field {
  return []ent.Field{
    field.String("description"),
}
}

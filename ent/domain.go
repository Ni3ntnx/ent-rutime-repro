// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entrepro/ent/domain"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Domain is the model entity for the Domain schema.
type Domain struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TenantID holds the value of the "tenant_id" field.
	TenantID string `json:"tenant_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Domain) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case domain.FieldID:
			values[i] = new(sql.NullInt64)
		case domain.FieldTenantID, domain.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Domain", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Domain fields.
func (d *Domain) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case domain.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		case domain.FieldTenantID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value.Valid {
				d.TenantID = value.String
			}
		case domain.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Domain.
// Note that you need to call Domain.Unwrap() before calling this method if this Domain
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Domain) Update() *DomainUpdateOne {
	return (&DomainClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Domain entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Domain) Unwrap() *Domain {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Domain is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Domain) String() string {
	var builder strings.Builder
	builder.WriteString("Domain(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("tenant_id=")
	builder.WriteString(d.TenantID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(d.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Domains is a parsable slice of Domain.
type Domains []*Domain

func (d Domains) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}

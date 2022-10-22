// Code generated by ent, DO NOT EDIT.

package domain

import (
	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the domain type in the database.
	Label = "domain"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// Table holds the table name of the domain in the database.
	Table = "domains"
)

// Columns holds all SQL columns for domain fields.
var Columns = []string{
	FieldID,
	FieldTenantID,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "entrepro/ent/runtime"
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
)

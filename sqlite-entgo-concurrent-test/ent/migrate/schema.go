// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DummiesColumns holds the columns for the "dummies" table.
	DummiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "age", Type: field.TypeInt, Nullable: true},
	}
	// DummiesTable holds the schema information for the "dummies" table.
	DummiesTable = &schema.Table{
		Name:       "dummies",
		Columns:    DummiesColumns,
		PrimaryKey: []*schema.Column{DummiesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DummiesTable,
	}
)

func init() {
}
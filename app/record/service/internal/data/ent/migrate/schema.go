// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// RecordColumns holds the columns for the "record" table.
	RecordColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "sign_in_index", Type: field.TypeInt},
		{Name: "reward", Type: field.TypeFloat64},
		{Name: "sign_in_day", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// RecordTable holds the schema information for the "record" table.
	RecordTable = &schema.Table{
		Name:       "record",
		Columns:    RecordColumns,
		PrimaryKey: []*schema.Column{RecordColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "record_user_id_sign_in_day",
				Unique:  true,
				Columns: []*schema.Column{RecordColumns[1], RecordColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		RecordTable,
	}
)

func init() {
	RecordTable.Annotation = &entsql.Annotation{
		Table: "record",
	}
}

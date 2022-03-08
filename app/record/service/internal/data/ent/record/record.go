// Code generated by entc, DO NOT EDIT.

package record

import (
	"time"
)

const (
	// Label holds the string label denoting the record type in the database.
	Label = "record"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldSignInIndex holds the string denoting the sign_in_index field in the database.
	FieldSignInIndex = "sign_in_index"
	// FieldReward holds the string denoting the reward field in the database.
	FieldReward = "reward"
	// FieldSignInDay holds the string denoting the sign_in_day field in the database.
	FieldSignInDay = "sign_in_day"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the record in the database.
	Table = "record"
)

// Columns holds all SQL columns for record fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldSignInIndex,
	FieldReward,
	FieldSignInDay,
	FieldCreatedAt,
	FieldUpdatedAt,
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

var (
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)
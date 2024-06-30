// Code generated by ent, DO NOT EDIT.

package blog

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the blog type in the database.
	Label = "blog"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldCtime holds the string denoting the ctime field in the database.
	FieldCtime = "ctime"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// Table holds the table name of the blog in the database.
	Table = "blog"
)

// Columns holds all SQL columns for blog fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldAddress,
	FieldUserID,
	FieldCtime,
	FieldDesc,
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

// OrderOption defines the ordering options for the Blog queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByCtime orders the results by the ctime field.
func ByCtime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCtime, opts...).ToFunc()
}

// ByDesc orders the results by the desc field.
func ByDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDesc, opts...).ToFunc()
}
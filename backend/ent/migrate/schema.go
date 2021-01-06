// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// BorrowsColumns holds the columns for the "borrows" table.
	BorrowsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeString, Unique: true},
	}
	// BorrowsTable holds the schema information for the "borrows" table.
	BorrowsTable = &schema.Table{
		Name:        "borrows",
		Columns:     BorrowsColumns,
		PrimaryKey:  []*schema.Column{BorrowsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ConfirmationsColumns holds the columns for the "confirmations" table.
	ConfirmationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "bookingdate", Type: field.TypeTime},
		{Name: "bookingstart", Type: field.TypeTime},
		{Name: "bookingend", Type: field.TypeTime},
		{Name: "hourstime", Type: field.TypeInt},
		{Name: "borrow_borrow_confirmation", Type: field.TypeInt, Nullable: true},
		{Name: "stadium_stadium_confirmation", Type: field.TypeInt, Nullable: true},
		{Name: "user_user_confirmation", Type: field.TypeInt, Nullable: true},
	}
	// ConfirmationsTable holds the schema information for the "confirmations" table.
	ConfirmationsTable = &schema.Table{
		Name:       "confirmations",
		Columns:    ConfirmationsColumns,
		PrimaryKey: []*schema.Column{ConfirmationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "confirmations_borrows_BorrowConfirmation",
				Columns: []*schema.Column{ConfirmationsColumns[5]},

				RefColumns: []*schema.Column{BorrowsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "confirmations_stadia_StadiumConfirmation",
				Columns: []*schema.Column{ConfirmationsColumns[6]},

				RefColumns: []*schema.Column{StadiaColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "confirmations_users_UserConfirmation",
				Columns: []*schema.Column{ConfirmationsColumns[7]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// StadiaColumns holds the columns for the "stadia" table.
	StadiaColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "namestadium", Type: field.TypeString, Unique: true},
	}
	// StadiaTable holds the schema information for the "stadia" table.
	StadiaTable = &schema.Table{
		Name:        "stadia",
		Columns:     StadiaColumns,
		PrimaryKey:  []*schema.Column{StadiaColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "birthday", Type: field.TypeTime},
		{Name: "telephone", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BorrowsTable,
		ConfirmationsTable,
		StadiaTable,
		UsersTable,
	}
)

func init() {
	ConfirmationsTable.ForeignKeys[0].RefTable = BorrowsTable
	ConfirmationsTable.ForeignKeys[1].RefTable = StadiaTable
	ConfirmationsTable.ForeignKeys[2].RefTable = UsersTable
}

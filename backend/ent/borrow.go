// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/tew147258/app/ent/borrow"
)

// Borrow is the model entity for the Borrow schema.
type Borrow struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BorrowQuery when eager-loading is set.
	Edges BorrowEdges `json:"edges"`
}

// BorrowEdges holds the relations/edges for other nodes in the graph.
type BorrowEdges struct {
	// BorrowConfirmation holds the value of the BorrowConfirmation edge.
	BorrowConfirmation []*Confirmation
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BorrowConfirmationOrErr returns the BorrowConfirmation value or an error if the edge
// was not loaded in eager-loading.
func (e BorrowEdges) BorrowConfirmationOrErr() ([]*Confirmation, error) {
	if e.loadedTypes[0] {
		return e.BorrowConfirmation, nil
	}
	return nil, &NotLoadedError{edge: "BorrowConfirmation"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Borrow) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // type
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Borrow fields.
func (b *Borrow) assignValues(values ...interface{}) error {
	if m, n := len(values), len(borrow.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	b.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field type", values[0])
	} else if value.Valid {
		b.Type = value.String
	}
	return nil
}

// QueryBorrowConfirmation queries the BorrowConfirmation edge of the Borrow.
func (b *Borrow) QueryBorrowConfirmation() *ConfirmationQuery {
	return (&BorrowClient{config: b.config}).QueryBorrowConfirmation(b)
}

// Update returns a builder for updating this Borrow.
// Note that, you need to call Borrow.Unwrap() before calling this method, if this Borrow
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Borrow) Update() *BorrowUpdateOne {
	return (&BorrowClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (b *Borrow) Unwrap() *Borrow {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Borrow is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Borrow) String() string {
	var builder strings.Builder
	builder.WriteString("Borrow(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", type=")
	builder.WriteString(b.Type)
	builder.WriteByte(')')
	return builder.String()
}

// Borrows is a parsable slice of Borrow.
type Borrows []*Borrow

func (b Borrows) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}

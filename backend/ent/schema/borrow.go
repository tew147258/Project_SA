package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Borrow holds the schema definition for the Borrow entity.
type Borrow struct {
	ent.Schema
}

// Fields of the Borrow.
func (Borrow) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").NotEmpty().Unique(),
	}
}

// Edges of the Borrow.
func (Borrow) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("BorrowConfirmation", Confirmation.Type),
	}
}

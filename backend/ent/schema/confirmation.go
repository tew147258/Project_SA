package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Confirmation holds the schema definition for the Confirmation entity.
type Confirmation struct {
	ent.Schema
}

// Fields of the Confirmation.
func (Confirmation) Fields() []ent.Field {
	return []ent.Field{
		field.Time("bookingdate"),
		field.Time("bookingstart"),
		field.Time("bookingend"),
		field.Int("hourstime"),
	}
}

// Edges of the Confirmation.
func (Confirmation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ConfirmationUser", User.Type).Ref("UserConfirmation").Unique(),
		edge.From("ConfirmationStadium", Stadium.Type).Ref("StadiumConfirmation").Unique(),
		edge.From("ConfirmationBorrow", Borrow.Type).Ref("BorrowConfirmation").Unique(),
	}
}

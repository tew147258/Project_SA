package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Stadium holds the schema definition for the Stadium entity.
type Stadium struct {
	ent.Schema
}

// Fields of the Stadium.
func (Stadium) Fields() []ent.Field {
	return []ent.Field{
		field.String("namestadium").NotEmpty().Unique(),
	}
}

// Edges of the Stadium.
func (Stadium) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("StadiumConfirmation", Confirmation.Type),
	}
}

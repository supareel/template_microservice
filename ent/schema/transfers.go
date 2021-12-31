package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Transfers holds the schema definition for the Transfers entity.
type Transfers struct {
	ent.Schema
}

// Fields of the Transfers.
func (Transfers) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("from_accout_id").Positive(),
		field.Int64("to_account_id").Positive(),
		field.Int64("amount").Positive().Default(0),
		field.Time("created_at").Default(time.Now().UTC),
	}
}

// Edges of the Transfers.
func (Transfers) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account_of", Accounts.Type).
			Ref("from_account_id").Unique(),
		edge.From("account_to", Accounts.Type).
			Ref("to_account_id").Unique(),
	}
}

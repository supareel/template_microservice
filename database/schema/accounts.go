package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Accounts holds the schema definition for the Accounts entity.
type Accounts struct {
	ent.Schema
}

// Fields of the Accounts.
func (Accounts) Fields() []ent.Field {
	return []ent.Field{
		field.String("owner").NotEmpty(),
		field.Int64("balance").Positive().Default(0),
		field.String("currency").NotEmpty().Default("INR"),
		field.Time("created_at").Default(time.Now().UTC),
	}
}

// Edges of the Accounts.
func (Accounts) Edges() []ent.Edge {
	return nil
}

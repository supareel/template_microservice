package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Transfers holds the schema definition for the Transfers entity.
type User struct {
	ent.Schema
}

// Fields of the Transfers.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("from_account_id").Positive(),
		field.Int("to_account_id").Positive(),
		field.Int64("amount").Positive().Default(0),
		field.Time("created_at").Default(time.Now().UTC),
	}
}

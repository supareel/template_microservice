package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Unique(),
		field.Text("name").NotEmpty(),
		field.Bool("is_done").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return nil
}

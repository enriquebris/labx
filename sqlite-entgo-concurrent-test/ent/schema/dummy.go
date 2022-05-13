package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Dummy holds the schema definition for the Dummy entity.
type Dummy struct {
	ent.Schema
}

// Fields of the Dummy.
func (Dummy) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Optional(),
		field.Int("age").
			Optional(),
	}
}

// Edges of the Dummy.
func (Dummy) Edges() []ent.Edge {
	return nil
}

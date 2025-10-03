package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

func (Task) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SonyFlakeIDMixin{},
		TimeMixin{},
	}
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("title"),
		field.
			Bool("completed").
			Default(false),
		field.
			Time("completed_at").
			Nillable().
			Optional().
			StructTag(`json:"completedAt,omitempty"`),
		field.
			Uint64("owner_id").
			StructTag(`json:"ownerID,omitempty"`),
	}
}

// Edges of the Task.

func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("owner", User.Type).
			Ref("tasks").
			Field("owner_id").
			Unique().
			Required(),
	}
}

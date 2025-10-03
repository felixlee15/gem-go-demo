package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// -------------------------------------------------
// Mixin definition

// SonyFlakeIDMixin implements the ent.Mixin for sharing
// SonyFlake ID field with package schemas.
type SonyFlakeIDMixin struct {
	mixin.Schema
}

func (SonyFlakeIDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			DefaultFunc(NextID).
			Immutable().
			Unique(),
	}
}

// TimeMixin implements the ent.Mixin for sharing
// time fields with package schemas.
type TimeMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			StructTag(`json:"createdAt"`).
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			StructTag(`json:"updatedAt"`).
			Annotations(
				entgql.OrderField("UPDATED_AT"),
			),
	}
}

// DeleteTimeMixin implements the ent.Mixin for sharing
// deleted_at field with package schemas.
type DeleteTimeMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (DeleteTimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("deleted_at").
			Optional().
			Nillable().
			StructTag(`json:"deletedAt"`),
	}
}

// ShopMixin implements the ent.Mixin for sharing
// shop_id field with package schemas.
type ShopMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (ShopMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("shop_id").
			StructTag(`json:"shopId"`).
			Optional().
			Annotations(entgql.Type("ID")),
	}
}

package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Request holds the schema definition for the Request entity.
type Request struct {
	ent.Schema
}

// Fields of the Request.
func (Request) Fields() []ent.Field {
	return []ent.Field{
		field.Int("contract_id").
			Annotations(
				entgql.OrderField("CONTRACT_ID"),
			),

		field.Int("vendor_id").
			Annotations(
				entgql.OrderField("VENDOR_ID"),
			),

		field.Int("contractor_id").
			Annotations(
				entgql.OrderField("CONTRACTOR_ID"),
			),

		field.Int("amount").
			Annotations(
				entgql.OrderField("AMOUNT"),
			),

		field.Enum("status").
			NamedValues("created", "created", "readyToCollect", "readyToCollect",
				"arrangeCollectionDate", "arrangeCollectionDate", "collected", "collected"). // note
			Default("created").
			Annotations(
				entgql.OrderField("STATUS"),
			),

		field.Time("collection_date").
			Optional().
			Nillable().
			Annotations(
				entgql.OrderField("COLLECTION_DATE"),
			),

		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Immutable().
			Annotations(
				entgql.OrderField("UPDATED_AT"),
			),
	}
}

// Edges of the Request.
func (Request) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Request) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "requests"},
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

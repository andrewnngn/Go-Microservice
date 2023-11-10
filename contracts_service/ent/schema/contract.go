package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Contract holds the schema definition for the Contract entity.
type Contract struct {
	ent.Schema
}

// Fields of the Contract.
func (Contract) Fields() []ent.Field {
	return []ent.Field{

		field.Int("vendor_id").
			Annotations(
				entgql.OrderField("VENDOR_ID"),
			),

		field.Enum("status").
			NamedValues("ongoing", "ongoing", "ended", "ended", "negotiating", "negotiating"). // note
			Default("negotiating").
			Annotations(
				entgql.OrderField("STATUS"),
			),

		field.Time("start_date").
			Optional().
			Nillable().
			Annotations(
				entgql.OrderField("START_DATE"),
			),

		field.Time("end_date").
			Optional().
			Nillable().
			Annotations(
				entgql.OrderField("END_DATE"),
			),

		field.Int("total_amount").
			Default(0).
			Annotations(
				entgql.OrderField("TOTAL_AMOUNT"),
			),

		field.Int("remaining_amount").
			Default(0).
			Annotations(
				entgql.OrderField("REMAINING_AMOUNT"),
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

// Edges of the Contract.
func (Contract) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Contract) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "contracts"},
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

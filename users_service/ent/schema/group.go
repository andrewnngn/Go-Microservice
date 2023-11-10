package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").
			Unique().
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),

		field.Text("address").
			Unique().
			NotEmpty().
			Annotations(
				entgql.OrderField("ADDRESS"),
			),

		field.Text("email").
			Unique().
			NotEmpty().
			Annotations(
				entgql.OrderField("EMAIL"),
			),

		field.Enum("role").
			NamedValues("planner", "planner", "vendor", "vendor", "contractor", "contractor"). // note
			Annotations(
				entgql.OrderField("ROLE"),
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

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (Group) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "groups"},
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

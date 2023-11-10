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

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Text("first_name").
			Unique().
			NotEmpty().
			Annotations(
				entgql.OrderField("FIRST_NAME"),
			),

		field.Text("last_name").
			Unique().
			NotEmpty().
			Annotations(
				entgql.OrderField("LAST_NAME"),
			),

		field.Text("email").
			Unique().
			NotEmpty().
			Annotations(
				entgql.OrderField("EMAIL"),
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

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("groups", Group.Type).
			Ref("users").
			Unique(),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users"},
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

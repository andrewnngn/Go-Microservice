package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Auth holds the schema definition for the Auth entity.
type Auth struct {
	ent.Schema
}

// Fields of the Auth.
func (Auth) Fields() []ent.Field {
	return []ent.Field{
		field.Text("username").
			Unique().
			NotEmpty().
			Annotations(
				entgql.OrderField("USERNAME"),
			),

		field.Text("password").
			NotEmpty().
			Annotations(
				entgql.OrderField("PASSWORD"),
			),

		field.Text("access_token").
			Nillable().
			Optional().
			Annotations(
				entgql.OrderField("ACCESS_TOKEN"),
				entgql.Skip(
					entgql.SkipMutationCreateInput,
					entgql.SkipMutationUpdateInput,
				),
			),

		field.Text("refresh_token").
			Nillable().
			Optional().
			Annotations(
				entgql.OrderField("REFRESH_TOKEN"),
			),

		field.Int("user_id").
			Unique().
			Annotations(
				entgql.OrderField("USER_ID"),
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

// Edges of the Auth.
func (Auth) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Auth) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users"},
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

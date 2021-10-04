package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty().Annotations(
			entgql.OrderField("TITLE"),
		),
		field.Text("content").NotEmpty(),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).
			Ref("posts").
			Unique().
			Annotations(entgql.Bind()),
	}
}

// Mixin of the Post.
func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

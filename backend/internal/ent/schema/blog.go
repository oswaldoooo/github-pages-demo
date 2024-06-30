package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Blog holds the schema definition for the Blog entity.
type Blog struct {
	ent.Schema
}

// Fields of the Blog.
func (Blog) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("title"),
		field.String("address"),
		field.Uint64("user_id"),
		field.Int64("ctime"),
		field.String("desc"),
	}
}

// Edges of the Blog.
func (Blog) Edges() []ent.Edge {
	return nil
}

func (Blog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "blog"},
	}
}

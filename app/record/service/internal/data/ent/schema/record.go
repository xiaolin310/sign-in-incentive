package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// Record holds the schema definition for the Record entity.
type Record struct {
	ent.Schema
}

// Fields of the Record.
func (Record) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id"),
		field.Int("sign_in_index"),
		field.Float("reward"),
		field.String("sign_in_day"),
		field.Time("created_at"),
		field.Time("updated_at").Default(time.Now),
	}
}

func (Record) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "record"},
	}
}

func (Record) Indexes() []ent.Index {
	return []ent.Index{
		// 唯一约束索引
		index.Fields("user_id", "sign_in_day").Unique(),
	}
}

// Edges of the Record.
func (Record) Edges() []ent.Edge {
	return nil
}

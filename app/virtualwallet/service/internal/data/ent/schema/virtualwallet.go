package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"time"
)

// VirtualWallet holds the schema definition for the VirtualWallet entity.
type VirtualWallet struct {
	ent.Schema
}

// Fields of the VirtualWallet.
func (VirtualWallet) Fields() []ent.Field {
	return []ent.Field {
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Int64("user_id"),
		field.Float("balance").
			Default(0.0),
		field.Time("created_at"),
		field.Time("updated_at").Default(time.Now),
	}
}

func (VirtualWallet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "virtual_wallet"},
	}
}

func (VirtualWallet) Indexes() []ent.Index {
	return []ent.Index{
		// 唯一约束索引
		index.Fields("user_id").Unique(),
	}
}

// Edges of the VirtualWallet.
func (VirtualWallet) Edges() []ent.Edge {
	return nil
}

package schema

import (
	"database/sql/driver"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"fmt"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age"),
		field.String("name"),
		field.Bytes("user_bytes").
			Optional().
			Nillable().
			GoType(UserBytes{}),
	}
}

type UserBytes [12]byte

// Scan satisfies the Scanner interface.
func (s *UserBytes) Scan(v interface{}) error {
	switch v := v.(type) {
	case []byte:
		copy(s[:], v)
	case string:
		copy(s[:], v)
	default:
		return fmt.Errorf("unexpcted type: %T", v)
	}
	return nil
}

// Value satisfies the Valuer interface.
func (s UserBytes) Value() (driver.Value, error) {
	return string(s[:]), nil
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

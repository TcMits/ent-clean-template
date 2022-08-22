package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixins of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.String("jwt_token_key").
			DefaultFunc(uuid.NewString).
			Sensitive(),
		field.String("password").
			Optional().
			Sensitive().
			SchemaType(map[string]string{
				dialect.MySQL: "char(32)",
			}),
		field.String("username").
			Unique().
			MaxLen(128).
			Match(regexp.MustCompile("^[a-zA-Z0-9]{6,128}$")).
			Comment("Required. 128 characters or fewer. Letters, digits only."),
		field.String("first_name").
			Default("").MaxLen(128),
		field.String("last_name").
			Default("").MaxLen(128),
		field.String("email").
			Unique().
			Validate(func(s string) error {
				return validation.Validate(s, is.Email)
			}),
		field.Bool("is_staff"),
		field.Bool("is_active").Default(true),
		field.Time("join_time").
			Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "User"},
	}
}

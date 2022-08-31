package schema

import (
	"regexp"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
			Optional().
			DefaultFunc(uuid.NewString).
			Sensitive().
			Annotations(
				entgql.Skip(entgql.SkipAll),
			),
		field.String("password").
			Optional().
			Sensitive().
			SchemaType(map[string]string{
				dialect.MySQL: "char(32)",
			}).
			Annotations(
				entgql.Skip(entgql.SkipAll),
			),
		field.String("username").
			Unique().
			MaxLen(128).
			Match(regexp.MustCompile("^[a-zA-Z0-9]{6,128}$")).
			Comment("Required. 128 characters or fewer. Letters, digits only."),
		field.String("first_name").
			Optional().
			Default("").
			MaxLen(128),
		field.String("last_name").
			Optional().
			Default("").
			MaxLen(128),
		field.String("email").
			Unique().
			Validate(func(s string) error {
				return validation.Validate(s, is.Email)
			}),
		field.Bool("is_staff").
			Optional().
			Default(false),
		field.Bool("is_superuser").
			Optional().
			Default(false),
		field.Bool("is_active").
			Optional().
			Default(true),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username"),
		index.Fields("create_time"),
	}
}

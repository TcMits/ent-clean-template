// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/TcMits/ent-clean-template/ent/schema"
	"github.com/TcMits/ent-clean-template/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescJwtTokenKey is the schema descriptor for jwt_token_key field.
	userDescJwtTokenKey := userFields[1].Descriptor()
	// user.DefaultJwtTokenKey holds the default value on creation for the jwt_token_key field.
	user.DefaultJwtTokenKey = userDescJwtTokenKey.Default.(func() string)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[3].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = func() func(string) error {
		validators := userDescUsername.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(username string) error {
			for _, fn := range fns {
				if err := fn(username); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescFirstName is the schema descriptor for first_name field.
	userDescFirstName := userFields[4].Descriptor()
	// user.DefaultFirstName holds the default value on creation for the first_name field.
	user.DefaultFirstName = userDescFirstName.Default.(string)
	// user.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	user.FirstNameValidator = userDescFirstName.Validators[0].(func(string) error)
	// userDescLastName is the schema descriptor for last_name field.
	userDescLastName := userFields[5].Descriptor()
	// user.DefaultLastName holds the default value on creation for the last_name field.
	user.DefaultLastName = userDescLastName.Default.(string)
	// user.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	user.LastNameValidator = userDescLastName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[6].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescIsStaff is the schema descriptor for is_staff field.
	userDescIsStaff := userFields[7].Descriptor()
	// user.DefaultIsStaff holds the default value on creation for the is_staff field.
	user.DefaultIsStaff = userDescIsStaff.Default.(bool)
	// userDescIsSuperuser is the schema descriptor for is_superuser field.
	userDescIsSuperuser := userFields[8].Descriptor()
	// user.DefaultIsSuperuser holds the default value on creation for the is_superuser field.
	user.DefaultIsSuperuser = userDescIsSuperuser.Default.(bool)
	// userDescIsActive is the schema descriptor for is_active field.
	userDescIsActive := userFields[9].Descriptor()
	// user.DefaultIsActive holds the default value on creation for the is_active field.
	user.DefaultIsActive = userDescIsActive.Default.(bool)
	// userDescJoinTime is the schema descriptor for join_time field.
	userDescJoinTime := userFields[10].Descriptor()
	// user.DefaultJoinTime holds the default value on creation for the join_time field.
	user.DefaultJoinTime = userDescJoinTime.Default.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}

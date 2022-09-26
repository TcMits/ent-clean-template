package model

import (
	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/ent/predicate"
	"github.com/TcMits/ent-clean-template/ent/user"
)

// OrderableColumns holds all SQL columns for ordering.
var OrderableColumns = []string{
	user.FieldID,
	user.FieldCreateTime,
	user.FieldUsername,
}

// User is the model entity for the User schema.
type (
	User            = ent.User
	PredicateUser   = predicate.User
	UserQuery       = ent.UserQuery
	UserMutation    = ent.UserMutation
	UserCreate      = ent.UserCreate
	UserUpdate      = ent.UserUpdate
	UserUpdateOne   = ent.UserUpdateOne
	UserCreateInput = ent.UserCreateInput
	UserUpdateInput = ent.UserUpdateInput
	UserWhereInput  = ent.UserWhereInput
	UserOrderInput  = ent.UserOrderInput
	UserSerializer  = ent.UserSerializer
)

var (
	descCreateTimeOrderField, _ = ent.ParseOrderField(
		ent.OrderDirectionDescPrefix + user.FieldCreateTime,
	)
	DefaultUserOrderInput = &UserOrderInput{descCreateTimeOrderField}
)

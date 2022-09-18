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
type User = ent.User
type PredicateUser = predicate.User
type UserQuery = ent.UserQuery
type UserMutation = ent.UserMutation
type UserCreate = ent.UserCreate
type UserUpdate = ent.UserUpdate
type UserUpdateOne = ent.UserUpdateOne
type UserCreateInput = ent.UserCreateInput
type UserUpdateInput = ent.UserUpdateInput
type UserWhereInput = ent.UserWhereInput
type UserOrderInput = ent.UserOrderInput
type UserSerializer = ent.UserSerializer

var descCreateTimeOrderField, _ = ent.ParseOrderField(ent.OrderDirectionDescPrefix + user.FieldCreateTime)
var DefaultUserWhereInput = &UserWhereInput{}
var DefaultUserOrderInput = &UserOrderInput{descCreateTimeOrderField}

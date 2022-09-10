// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
)

type UserUpdateRepository struct {
	client              *Client
	preUpdateFunctions  []func(context.Context, *Client, *User, *UserUpdateInput) error
	postUpdateFunctions []func(context.Context, *Client, *User, *User) error
	isAtomic            bool
}

func NewUserUpdateRepository(
	client *Client,
	preUpdateFunctions []func(context.Context, *Client, *User, *UserUpdateInput) error,
	postUpdateFunctions []func(context.Context, *Client, *User, *User) error,
	isAtomic bool,
) *UserUpdateRepository {
	return &UserUpdateRepository{
		client:              client,
		preUpdateFunctions:  preUpdateFunctions,
		postUpdateFunctions: postUpdateFunctions,
		isAtomic:            isAtomic,
	}
}

func (r *UserUpdateRepository) runPreUpdate(
	ctx context.Context, client *Client, instance *User, i *UserUpdateInput,
) error {
	for _, function := range r.preUpdateFunctions {
		err := function(ctx, client, instance, i)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *UserUpdateRepository) runPostUpdate(
	ctx context.Context, client *Client, oldInstance *User, newInstance *User,
) error {
	for _, function := range r.postUpdateFunctions {
		err := function(ctx, client, oldInstance, newInstance)
		if err != nil {
			return err
		}
	}
	return nil
}

// using in Tx
func (r *UserUpdateRepository) UpdateWithClient(
	ctx context.Context, client *Client, instance *User, input *UserUpdateInput,
) (*User, error) {
	err := r.runPreUpdate(ctx, client, instance, input)
	if err != nil {
		return nil, err
	}
	newInstance, err := client.User.UpdateOne(instance).SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}
	err = r.runPostUpdate(ctx, client, instance, newInstance)
	if err != nil {
		return nil, err
	}
	return newInstance, nil
}

func (r *UserUpdateRepository) Update(
	ctx context.Context, instance *User, input *UserUpdateInput,
) (*User, error) {
	if !r.isAtomic {
		return r.UpdateWithClient(ctx, r.client, instance, input)
	}
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	instance, err = r.UpdateWithClient(ctx, tx.Client(), instance, input)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("rolling back transaction: %w", rerr)
		}
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("committing transaction: %w", err)
	}
	return instance, nil
}

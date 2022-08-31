package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/TcMits/ent-clean-template/internal/repository"
	"github.com/TcMits/ent-clean-template/internal/testutils"
	"github.com/TcMits/ent-clean-template/pkg/entity/factory"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_getModelUseCase_Get(t *testing.T) {
	type fields struct {
		repository           repository.GetModelRepository[*model.User, *model.UserWhereInput]
		toRepoWhereInputFunc ConverFunc[*model.UserWhereInput, *model.UserWhereInput]
		wrapGetErrorFunc     func(error) error
	}
	type args struct {
		ctx   context.Context
		input *model.UserWhereInput
	}
	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()
	u, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)

	repository := repository.NewUserRepository(client)
	newUUID := uuid.New()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				repository: repository,
				toRepoWhereInputFunc: func(uwi *model.UserWhereInput) (*model.UserWhereInput, error) {
					return uwi, nil
				},
				wrapGetErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx: ctx,
				input: &model.UserWhereInput{
					ID: &u.ID,
				},
			},
			want: u,
		},
		{
			name: "WhereInputFuncError",
			fields: fields{
				repository: repository,
				toRepoWhereInputFunc: func(uwi *model.UserWhereInput) (*model.UserWhereInput, error) {
					return nil, errors.New("test")
				},
				wrapGetErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx: ctx,
				input: &model.UserWhereInput{
					ID: &u.ID,
				},
			},
			wantErr: true,
		},
		{
			name: "GetError",
			fields: fields{
				repository: repository,
				toRepoWhereInputFunc: func(uwi *model.UserWhereInput) (*model.UserWhereInput, error) {
					return uwi, nil
				},
				wrapGetErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: &model.UserWhereInput{ID: &newUUID},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &getModelUseCase[*model.User, *model.UserWhereInput, *model.UserWhereInput]{
				repository:           tt.fields.repository,
				toRepoWhereInputFunc: tt.fields.toRepoWhereInputFunc,
				wrapGetErrorFunc:     tt.fields.wrapGetErrorFunc,
			}
			got, err := l.Get(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("getModelUseCase.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && tt.want != nil && got.ID != tt.want.ID {
				t.Errorf("getModelUseCase.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countModelUseCase_Count(t *testing.T) {
	type fields struct {
		repository           repository.CountModelRepository[*model.UserWhereInput]
		toRepoWhereInputFunc ConverFunc[*model.UserWhereInput, *model.UserWhereInput]
		wrapCountErrorFunc   func(error) error
	}
	type args struct {
		ctx   context.Context
		input *model.UserWhereInput
	}
	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()
	u, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)

	repository := repository.NewUserRepository(client)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				repository: repository,
				toRepoWhereInputFunc: func(uwi *model.UserWhereInput) (*model.UserWhereInput, error) {
					return uwi, nil
				},
				wrapCountErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx: ctx,
				input: &model.UserWhereInput{
					ID: &u.ID,
				},
			},
			want: 1,
		},
		{
			name: "WhereInputFuncError",
			fields: fields{
				repository: repository,
				toRepoWhereInputFunc: func(uwi *model.UserWhereInput) (*model.UserWhereInput, error) {
					return nil, errors.New("test")
				},
				wrapCountErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx: ctx,
				input: &model.UserWhereInput{
					ID: &u.ID,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &countModelUseCase[*model.UserWhereInput, *model.UserWhereInput]{
				repository:           tt.fields.repository,
				toRepoWhereInputFunc: tt.fields.toRepoWhereInputFunc,
				wrapCountErrorFunc:   tt.fields.wrapCountErrorFunc,
			}
			got, err := l.Count(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("countModelUseCase.Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("countModelUseCase.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listModelUseCase_List(t *testing.T) {
	type fields struct {
		repository           repository.ListModelRepository[*model.User, *model.UserOrderInput, *model.UserWhereInput]
		toRepoWhereInputFunc ConverFunc[*model.UserWhereInput, *model.UserWhereInput]
		toRepoOrderInputFunc ConverFunc[*model.UserOrderInput, *model.UserOrderInput]
		wrapListErrorFunc    func(error) error
	}
	type args struct {
		ctx        context.Context
		limit      int
		offset     int
		orderInput *model.UserOrderInput
		whereInput *model.UserWhereInput
	}
	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()
	u, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)

	repository := repository.NewUserRepository(client)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.User
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				repository: repository,
				toRepoWhereInputFunc: func(uwi *model.UserWhereInput) (*model.UserWhereInput, error) {
					return uwi, nil
				},
				toRepoOrderInputFunc: func(uoi *model.UserOrderInput) (*model.UserOrderInput, error) {
					return uoi, nil
				},
				wrapListErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:        ctx,
				limit:      1,
				offset:     0,
				orderInput: model.DefaultUserOrderInput,
				whereInput: model.DefaultUserWhereInput,
			},
			want: []*model.User{u},
		},
		{
			name: "WhereInputFuncError",
			fields: fields{
				repository: repository,
				toRepoWhereInputFunc: func(uwi *model.UserWhereInput) (*model.UserWhereInput, error) {
					return nil, errors.New("test")
				},
				toRepoOrderInputFunc: func(uoi *model.UserOrderInput) (*model.UserOrderInput, error) {
					return uoi, nil
				},
				wrapListErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:        ctx,
				limit:      1,
				offset:     0,
				orderInput: model.DefaultUserOrderInput,
				whereInput: model.DefaultUserWhereInput,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "OrderInputFuncError",
			fields: fields{
				repository: repository,
				toRepoWhereInputFunc: func(uwi *model.UserWhereInput) (*model.UserWhereInput, error) {
					return uwi, nil
				},
				toRepoOrderInputFunc: func(uoi *model.UserOrderInput) (*model.UserOrderInput, error) {
					return nil, errors.New("test")
				},
				wrapListErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:        ctx,
				limit:      1,
				offset:     0,
				orderInput: model.DefaultUserOrderInput,
				whereInput: model.DefaultUserWhereInput,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &listModelUseCase[*model.User, *model.UserOrderInput, *model.UserWhereInput, *model.UserOrderInput, *model.UserWhereInput]{
				repository:           tt.fields.repository,
				toRepoWhereInputFunc: tt.fields.toRepoWhereInputFunc,
				toRepoOrderInputFunc: tt.fields.toRepoOrderInputFunc,
				wrapListErrorFunc:    tt.fields.wrapListErrorFunc,
			}
			got, err := l.List(tt.args.ctx, tt.args.limit, tt.args.offset, tt.args.orderInput, tt.args.whereInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("listModelUseCase.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && tt.want != nil {
				for i, gotItem := range got {
					if gotItem.ID != tt.want[i].ID {
						t.Errorf("listModelUseCase.List() = %v, want %v", gotItem, tt.want[i])
					}
				}
			}
		})
	}
}

package usecase

import (
	"context"
	"errors"
	"net/url"
	"reflect"
	"testing"

	"github.com/TcMits/ent-clean-template/internal/repository"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/TcMits/ent-clean-template/pkg/tool/lazy"
	gomock "github.com/golang/mock/gomock"
)

func TestNewIsAuthenticatedPermissionChecker(t *testing.T) {
	type args struct {
		ctx      context.Context
		instance *model.User
	}

	ctx := context.Background()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "IsAuthenticated",
			args: args{
				ctx:      ctx,
				instance: &model.User{},
			},
			wantErr: false,
		},
		{
			name: "IsNotAuthenticated",
			args: args{
				ctx:      ctx,
				instance: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewIsAuthenticatedPermissionChecker()
			if gotResult := got.Check(tt.args.ctx, tt.args.instance); (gotResult != nil) != tt.wantErr {
				t.Errorf(
					"NewIsAuthenticatedPermissionChecker() = error %v, wantErr %v",
					got,
					tt.wantErr,
				)
			}
		})
	}
}

func TestNewIsSuperuserPermissionChecker(t *testing.T) {
	type args struct {
		ctx      context.Context
		instance *model.User
	}

	ctx := context.Background()

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "IsSuperuser",
			args: args{
				ctx:      ctx,
				instance: &model.User{IsSuperuser: true},
			},
			wantErr: false,
		},
		{
			name: "IsNotSuperuser",
			args: args{
				ctx:      ctx,
				instance: &model.User{IsSuperuser: false},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewIsSuperuserPermissionChecker()
			if gotResult := got.Check(tt.args.ctx, tt.args.instance); (gotResult != nil) != tt.wantErr {
				t.Errorf(
					"NewIsSuperuserPermissionChecker() = error %v, wantErr %v",
					got,
					tt.wantErr,
				)
			}
		})
	}
}

func Test_getPublicMeUseCase_Get(t *testing.T) {
	type args struct {
		ctx context.Context
		i   *struct{}
	}

	u := &model.User{}
	parentCtx := context.Background()
	ctx1 := context.WithValue(parentCtx, _currentUserKey, lazy.NewLazyObject(
		func() *model.User { return u },
	))

	tests := []struct {
		name    string
		u       *getPublicMeUseCase
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name: "Success",
			u:    new(getPublicMeUseCase),
			args: args{
				ctx: ctx1,
				i:   new(struct{}),
			},
			want: u,
		},
		{
			name: "WithoutUser",
			u:    new(getPublicMeUseCase),
			args: args{
				ctx: parentCtx,
				i:   new(struct{}),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &getPublicMeUseCase{}
			got, err := u.Get(tt.args.ctx, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPublicMeUseCase.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPublicMeUseCase.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateUpdateInputPublicMeUseCase_validateEmail(t *testing.T) {
	type fields struct {
		repository repository.GetModelRepository[*model.User, *model.UserWhereInput]
	}
	type args struct {
		ctx      context.Context
		instance *model.User
		i        *useCaseModel.PublicMeUseCaseUpdateInput
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	getRepo := repository.NewMockGetModelRepository[*model.User, *model.UserWhereInput](ctrl)
	email1 := "test1@gmail.com"
	email2 := "test2@gmail.com"

	getRepo.EXPECT().Get(
		gomock.Eq(ctx), gomock.Eq(&model.UserWhereInput{Email: &email1}),
	).Return(
		new(model.User), nil,
	).AnyTimes()

	getRepo.EXPECT().Get(
		gomock.Eq(ctx), gomock.Eq(&model.UserWhereInput{Email: &email2}),
	).Return(
		nil, errors.New(""),
	).AnyTimes()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "HaveUser",
			fields: fields{
				repository: getRepo,
			},
			args: args{
				ctx:      ctx,
				instance: &model.User{Email: "test@gmail.com"},
				i:        &useCaseModel.PublicMeUseCaseUpdateInput{Email: &email1},
			},
			wantErr: true,
		},
		{
			name: "HaveNotUser",
			fields: fields{
				repository: getRepo,
			},
			args: args{
				ctx:      ctx,
				instance: &model.User{Email: "test@gmail.com"},
				i:        &useCaseModel.PublicMeUseCaseUpdateInput{Email: &email2},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &validateUpdateInputPublicMeUseCase{
				repository: tt.fields.repository,
			}
			if err := u.validateEmail(tt.args.ctx, tt.args.instance, tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf(
					"validateUpdateInputPublicMeUseCase.validateEmail() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}

func Test_validateUpdateInputPublicMeUseCase_validateUsername(t *testing.T) {
	type fields struct {
		repository repository.GetModelRepository[*model.User, *model.UserWhereInput]
	}
	type args struct {
		ctx      context.Context
		instance *model.User
		i        *useCaseModel.PublicMeUseCaseUpdateInput
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	getRepo := repository.NewMockGetModelRepository[*model.User, *model.UserWhereInput](ctrl)
	username1 := "test1"
	username2 := "test2"

	getRepo.EXPECT().Get(
		gomock.Eq(ctx), gomock.Eq(&model.UserWhereInput{Username: &username1}),
	).Return(
		new(model.User), nil,
	).AnyTimes()

	getRepo.EXPECT().Get(
		gomock.Eq(ctx), gomock.Eq(&model.UserWhereInput{Username: &username2}),
	).Return(
		nil, errors.New(""),
	).AnyTimes()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "HaveUser",
			fields: fields{
				repository: getRepo,
			},
			args: args{
				ctx:      ctx,
				instance: &model.User{Username: "test"},
				i:        &useCaseModel.PublicMeUseCaseUpdateInput{Username: &username1},
			},
			wantErr: true,
		},
		{
			name: "HaveNotUser",
			fields: fields{
				repository: getRepo,
			},
			args: args{
				ctx:      ctx,
				instance: &model.User{Username: "test"},
				i:        &useCaseModel.PublicMeUseCaseUpdateInput{Username: &username2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &validateUpdateInputPublicMeUseCase{
				repository: tt.fields.repository,
			}
			if err := u.validateUsername(tt.args.ctx, tt.args.instance, tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf(
					"validateUpdateInputPublicMeUseCase.validateUsername() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}

func Test_validateUpdateInputPublicMeUseCase_Validate(t *testing.T) {
	type fields struct {
		repository repository.GetModelRepository[*model.User, *model.UserWhereInput]
	}
	type args struct {
		ctx      context.Context
		instance *model.User
		i        *useCaseModel.PublicMeUseCaseUpdateInput
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	getRepo := repository.NewMockGetModelRepository[*model.User, *model.UserWhereInput](ctrl)
	username := "test"
	email := "test@gmail.com"
	lastName := "last name"
	firstName := "first name"

	getRepo.EXPECT().Get(
		gomock.Eq(ctx), gomock.Any(),
	).Return(
		nil, errors.New(""),
	).AnyTimes()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.UserUpdateInput
		wantErr bool
	}{
		{
			name:   "Success",
			fields: fields{repository: getRepo},
			args: args{
				ctx:      ctx,
				instance: new(model.User),
				i: &useCaseModel.PublicMeUseCaseUpdateInput{
					Username:  &username,
					Email:     &email,
					LastName:  &lastName,
					FirstName: &firstName,
				},
			},
			want: &model.UserUpdateInput{
				Username:  &username,
				Email:     &email,
				LastName:  &lastName,
				FirstName: &firstName,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &validateUpdateInputPublicMeUseCase{
				repository: tt.fields.repository,
			}
			got, err := u.Validate(tt.args.ctx, tt.args.instance, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"validateUpdateInputPublicMeUseCase.Validate() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf(
					"validateUpdateInputPublicMeUseCase.Validate() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}

func TestNewPublicMeUseCase(t *testing.T) {
	type args struct {
		getRepository    repository.GetModelRepository[*model.User, *model.UserWhereInput]
		updateRepository repository.UpdateModelRepository[*model.User, *model.UserUpdateInput]
		getURL           func(*url.URL, url.Values, ...any) string
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	getRepo := repository.NewMockGetModelRepository[*model.User, *model.UserWhereInput](ctrl)
	updateRepo := repository.NewMockUpdateModelRepository[*model.User, *model.UserUpdateInput](ctrl)

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{
				getRepository:    getRepo,
				updateRepository: updateRepo,
				getURL:           func(u *url.URL, v url.Values, a ...any) string { return "" },
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = NewPublicMeUseCase(tt.args.getRepository, tt.args.updateRepository, tt.args.getURL)
		})
	}
}

package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	jwtKit "github.com/golang-jwt/jwt/v4"
	gomock "github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/TcMits/ent-clean-template/internal/repository"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/TcMits/ent-clean-template/pkg/tool/password"
)

func TestNewLoginUseCase(t *testing.T) {
	type args struct {
		getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput]
		secret        string
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	getRepository := repository.NewMockGetModelRepository[*model.User, *model.UserWhereInput](ctrl)

	want := loginUseCase{
		getRepository: getRepository,
		secret:        "secret",
	}

	tests := []struct {
		name string
		args args
		want LoginUseCase[*useCaseModel.LoginInput, *useCaseModel.JWTAuthenticatedPayload, *useCaseModel.RefreshTokenInput, *model.User]
	}{
		{
			name: "Success",
			args: args{
				getRepository: getRepository,
				secret:        "secret",
			},
			want: &want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLoginUseCase(tt.args.getRepository, tt.args.secret); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("NewLoginUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loginUseCase_getUserMapClaims(t *testing.T) {
	type fields struct {
		getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput]
		secret        string
	}
	type args struct {
		user *model.User
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	getRepository := repository.NewMockGetModelRepository[*model.User, *model.UserWhereInput](ctrl)
	u := &model.User{
		ID:          uuid.New(),
		Email:       "test@gmail.com",
		JwtTokenKey: uuid.NewString(),
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   jwtKit.MapClaims
	}{
		{
			name: "Success",
			fields: fields{
				getRepository: getRepository,
				secret:        "secret",
			},
			args: args{
				user: u,
			},
			want: jwtKit.MapClaims{
				_idFieldName:  u.ID.String(),
				"email":       u.Email,
				_keyFieldName: u.JwtTokenKey,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &loginUseCase{
				getRepository: tt.fields.getRepository,
				secret:        tt.fields.secret,
			}
			if got := l.getUserMapClaims(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loginUseCase.getUserMapClaims() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loginUseCase_getUserFromMapClaims(t *testing.T) {
	type fields struct {
		getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput]
		secret        string
	}
	type args struct {
		ctx          context.Context
		jwtMapClaims jwtKit.MapClaims
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	getRepository := repository.NewMockGetModelRepository[*model.User, *model.UserWhereInput](ctrl)
	u := &model.User{
		ID:          uuid.New(),
		Email:       "test@gmail.com",
		JwtTokenKey: uuid.NewString(),
	}
	isActive := true

	getRepository.EXPECT().Get(
		gomock.Eq(ctx), gomock.Eq(&model.UserWhereInput{ID: &u.ID, IsActive: &isActive}),
	).Return(
		u, nil,
	).AnyTimes()

	getRepository.EXPECT().Get(
		gomock.Eq(ctx), gomock.Any(),
	).Return(
		nil, errors.New(""),
	).AnyTimes()

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
				getRepository: getRepository,
				secret:        "secret",
			},
			args: args{
				ctx: ctx,
				jwtMapClaims: jwtKit.MapClaims{
					_idFieldName:  u.ID.String(),
					"email":       u.Email,
					_keyFieldName: u.JwtTokenKey,
				},
			},
			want: u,
		},
		{
			name: "MissingID",
			fields: fields{
				getRepository: getRepository,
				secret:        "secret",
			},
			args: args{
				ctx: ctx,
				jwtMapClaims: jwtKit.MapClaims{
					"email":       u.Email,
					_keyFieldName: u.JwtTokenKey,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "WrongUUID",
			fields: fields{
				getRepository: getRepository,
				secret:        "secret",
			},
			args: args{
				ctx: ctx,
				jwtMapClaims: jwtKit.MapClaims{
					_idFieldName:  "hello",
					"email":       u.Email,
					_keyFieldName: u.JwtTokenKey,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "WrongUUID",
			fields: fields{
				getRepository: getRepository,
				secret:        "secret",
			},
			args: args{
				ctx: ctx,
				jwtMapClaims: jwtKit.MapClaims{
					_idFieldName:  uuid.NewString(),
					"email":       u.Email,
					_keyFieldName: u.JwtTokenKey,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "WrongKey",
			fields: fields{
				getRepository: getRepository,
				secret:        "secret",
			},
			args: args{
				ctx: ctx,
				jwtMapClaims: jwtKit.MapClaims{
					_idFieldName:  u.ID,
					"email":       u.Email,
					_keyFieldName: "hello",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &loginUseCase{
				getRepository: tt.fields.getRepository,
				secret:        tt.fields.secret,
			}
			got, err := l.getUserFromMapClaims(tt.args.ctx, tt.args.jwtMapClaims)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"loginUseCase.getUserFromMapClaims() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}
			if got != nil && tt.want != nil && !reflect.DeepEqual(got.ID, tt.want.ID) {
				t.Errorf("loginUseCase.getUserFromMapClaims() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loginUseCase_createAccessToken(t *testing.T) {
	type fields struct {
		getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput]
		secret        string
	}
	type args struct {
		user *model.User
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	getRepository := repository.NewMockGetModelRepository[*model.User, *model.UserWhereInput](ctrl)
	u := &model.User{
		ID:          uuid.New(),
		Email:       "test@gmail.com",
		JwtTokenKey: uuid.NewString(),
	}
	isActive := true

	getRepository.EXPECT().Get(
		gomock.Eq(ctx), gomock.Eq(&model.UserWhereInput{ID: &u.ID, IsActive: &isActive}),
	).Return(
		u, nil,
	).AnyTimes()

	getRepository.EXPECT().Get(
		gomock.Eq(ctx), gomock.Any(),
	).Return(
		nil, errors.New(""),
	).AnyTimes()

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
				getRepository: getRepository,
				secret:        "Dummy",
			},
			args: args{
				u,
			},
			want: u,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &loginUseCase{
				getRepository: tt.fields.getRepository,
				secret:        tt.fields.secret,
			}
			got, err := l.createAccessToken(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("loginUseCase.createAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if newU, err := l.parseAccessToken(ctx, got); err != nil ||
				!reflect.DeepEqual(newU.ID, tt.want.ID) {
				t.Errorf("loginUseCase.createAccessToken() = %v, want %v", newU, tt.want)
			}
		})
	}
}

func Test_loginUseCase_createRefreshToken(t *testing.T) {
	type fields struct {
		getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput]
		secret        string
	}
	type args struct {
		user *model.User
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	getRepository := repository.NewMockGetModelRepository[*model.User, *model.UserWhereInput](ctrl)
	u := &model.User{
		ID:          uuid.New(),
		Email:       "test@gmail.com",
		JwtTokenKey: uuid.NewString(),
	}
	isActive := true

	getRepository.EXPECT().Get(
		gomock.Eq(ctx), gomock.Eq(&model.UserWhereInput{ID: &u.ID, IsActive: &isActive}),
	).Return(
		u, nil,
	).AnyTimes()

	getRepository.EXPECT().Get(
		gomock.Eq(ctx), gomock.Any(),
	).Return(
		nil, errors.New(""),
	).AnyTimes()

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
				getRepository: getRepository,
				secret:        "Dummy",
			},
			args: args{user: u},
			want: u,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &loginUseCase{
				getRepository: tt.fields.getRepository,
				secret:        tt.fields.secret,
			}
			got, err := l.createRefreshToken(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"loginUseCase.createRefreshToken() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}
			if newU, err := l.parseRefreshToken(ctx, got); err != nil ||
				!reflect.DeepEqual(newU.ID, tt.want.ID) {
				t.Errorf("loginUseCase.createRefreshToken() = %v, want %v", newU, tt.want)
			}
		})
	}
}

func Test_loginUseCase_parseAccessToken(t *testing.T) {
	type fields struct {
		getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput]
		secret        string
	}
	type args struct {
		ctx   context.Context
		token string
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	getRepository := repository.NewMockGetModelRepository[*model.User, *model.UserWhereInput](ctrl)
	u := &model.User{
		ID:          uuid.New(),
		Email:       "test@gmail.com",
		JwtTokenKey: uuid.NewString(),
	}
	isActive := true

	getRepository.EXPECT().Get(
		gomock.Eq(ctx), gomock.Eq(&model.UserWhereInput{ID: &u.ID, IsActive: &isActive}),
	).Return(
		u, nil,
	).AnyTimes()

	getRepository.EXPECT().Get(
		gomock.Eq(ctx), gomock.Any(),
	).Return(
		nil, errors.New(""),
	).AnyTimes()

	l := &loginUseCase{
		getRepository: getRepository,
		secret:        "Dummy",
	}
	token, err := l.createAccessToken(u)
	require.NoError(t, err)

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
				getRepository: getRepository,
				secret:        "Dummy",
			},
			args: args{
				ctx:   ctx,
				token: token,
			},
			want: u,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &loginUseCase{
				getRepository: tt.fields.getRepository,
				secret:        tt.fields.secret,
			}
			got, err := l.parseAccessToken(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("loginUseCase.parseAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.ID, tt.want.ID) {
				t.Errorf("loginUseCase.parseAccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loginUseCase_parseRefreshToken(t *testing.T) {
	type fields struct {
		getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput]
		secret        string
	}
	type args struct {
		ctx               context.Context
		refreshTokenInput *useCaseModel.RefreshTokenInput
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	getRepository := repository.NewMockGetModelRepository[*model.User, *model.UserWhereInput](ctrl)
	u := &model.User{
		ID:          uuid.New(),
		Email:       "test@gmail.com",
		JwtTokenKey: uuid.NewString(),
	}
	isActive := true

	getRepository.EXPECT().Get(
		gomock.Eq(ctx), gomock.Eq(&model.UserWhereInput{ID: &u.ID, IsActive: &isActive}),
	).Return(
		u, nil,
	).AnyTimes()

	getRepository.EXPECT().Get(
		gomock.Eq(ctx), gomock.Any(),
	).Return(
		nil, errors.New(""),
	).AnyTimes()

	l := &loginUseCase{
		getRepository: getRepository,
		secret:        "Dummy",
	}
	refreshTokenInput, err := l.createRefreshToken(u)
	require.NoError(t, err)

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
				getRepository: getRepository,
				secret:        "Dummy",
			},
			args: args{
				ctx:               ctx,
				refreshTokenInput: refreshTokenInput,
			},
			want: u,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &loginUseCase{
				getRepository: tt.fields.getRepository,
				secret:        tt.fields.secret,
			}
			got, err := l.parseRefreshToken(tt.args.ctx, tt.args.refreshTokenInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("loginUseCase.parseRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.ID, tt.want.ID) {
				t.Errorf("loginUseCase.parseRefreshToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loginUseCase_Login(t *testing.T) {
	type fields struct {
		getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput]
		secret        string
	}
	type args struct {
		ctx        context.Context
		loginInput *useCaseModel.LoginInput
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	getRepository := repository.NewMockGetModelRepository[*model.User, *model.UserWhereInput](ctrl)
	p, err := password.GetHashPassword("12345678")
	require.NoError(t, err)
	u := &model.User{
		Username:    "test",
		ID:          uuid.New(),
		Email:       "test@gmail.com",
		JwtTokenKey: uuid.NewString(),
		Password:    p,
	}
	isActive := true

	getRepository.EXPECT().Get(
		gomock.Eq(ctx), gomock.Eq(&model.UserWhereInput{ID: &u.ID, IsActive: &isActive}),
	).Return(
		u, nil,
	).AnyTimes()

	getRepository.EXPECT().Get(
		gomock.Eq(ctx),
		gomock.Eq(&model.UserWhereInput{Username: &u.Username, IsActive: &isActive}),
	).Return(
		u, nil,
	).AnyTimes()

	getRepository.EXPECT().Get(
		gomock.Eq(ctx), gomock.Any(),
	).Return(
		nil, errors.New(""),
	).AnyTimes()

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
				getRepository: getRepository,
				secret:        "Dummy",
			},
			args: args{
				ctx: ctx,
				loginInput: &useCaseModel.LoginInput{
					Username: u.Username,
					Password: "12345678",
				},
			},
			want: u,
		},
		{
			name: "WrongPassword",
			fields: fields{
				getRepository: getRepository,
				secret:        "Dummy",
			},
			args: args{
				ctx: ctx,
				loginInput: &useCaseModel.LoginInput{
					Username: u.Username,
					Password: "1234567",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &loginUseCase{
				getRepository: tt.fields.getRepository,
				secret:        tt.fields.secret,
			}
			got, err := l.Login(tt.args.ctx, tt.args.loginInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("loginUseCase.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				if newU, err := l.parseAccessToken(ctx, got.AccessToken); err != nil ||
					!reflect.DeepEqual(newU.ID, tt.want.ID) {
					t.Errorf("loginUseCase.Login() = %v, want %v", newU, tt.want)
				}
				if newU, err := l.parseRefreshToken(ctx, &useCaseModel.RefreshTokenInput{
					RefreshToken: got.RefreshToken,
					RefreshKey:   got.RefreshKey,
				}); err != nil || !reflect.DeepEqual(newU.ID, tt.want.ID) {
					t.Errorf("loginUseCase.Login() = %v, want %v", newU, tt.want)
				}
			}
		})
	}
}

func Test_loginUseCase_RefreshToken(t *testing.T) {
	type fields struct {
		getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput]
		secret        string
	}
	type args struct {
		ctx               context.Context
		refreshTokenInput *useCaseModel.RefreshTokenInput
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	getRepository := repository.NewMockGetModelRepository[*model.User, *model.UserWhereInput](ctrl)
	u := &model.User{
		Username:    "test",
		ID:          uuid.New(),
		Email:       "test@gmail.com",
		JwtTokenKey: uuid.NewString(),
	}
	isActive := true

	getRepository.EXPECT().Get(
		gomock.Eq(ctx), gomock.Eq(&model.UserWhereInput{ID: &u.ID, IsActive: &isActive}),
	).Return(
		u, nil,
	).AnyTimes()

	getRepository.EXPECT().Get(
		gomock.Eq(ctx), gomock.Any(),
	).Return(
		nil, errors.New(""),
	).AnyTimes()

	l := &loginUseCase{
		getRepository: getRepository,
		secret:        "Dummy",
	}
	refreshTokenInput, err := l.createRefreshToken(u)
	require.NoError(t, err)

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
				getRepository: getRepository,
				secret:        "Dummy",
			},
			args: args{
				ctx:               ctx,
				refreshTokenInput: refreshTokenInput,
			},
			want: u,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &loginUseCase{
				getRepository: tt.fields.getRepository,
				secret:        tt.fields.secret,
			}
			got, err := l.RefreshToken(tt.args.ctx, tt.args.refreshTokenInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("loginUseCase.RefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if newU, err := l.parseAccessToken(ctx, got); err != nil ||
				!reflect.DeepEqual(newU.ID, tt.want.ID) {
				t.Errorf("loginUseCase.RefreshToken() = %v, want %v", newU, tt.want)
			}
		})
	}
}

func Test_loginUseCase_VerifyToken(t *testing.T) {
	type fields struct {
		getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput]
		secret        string
	}
	type args struct {
		ctx   context.Context
		token string
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	getRepository := repository.NewMockGetModelRepository[*model.User, *model.UserWhereInput](ctrl)
	u := &model.User{
		Username:    "test",
		ID:          uuid.New(),
		Email:       "test@gmail.com",
		JwtTokenKey: uuid.NewString(),
	}
	isActive := true

	getRepository.EXPECT().Get(
		gomock.Eq(ctx), gomock.Eq(&model.UserWhereInput{ID: &u.ID, IsActive: &isActive}),
	).Return(
		u, nil,
	).AnyTimes()

	getRepository.EXPECT().Get(
		gomock.Eq(ctx), gomock.Any(),
	).Return(
		nil, errors.New(""),
	).AnyTimes()

	l := &loginUseCase{
		getRepository: getRepository,
		secret:        "Dummy",
	}
	token, err := l.createAccessToken(u)
	require.NoError(t, err)

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
				getRepository: getRepository,
				secret:        "Dummy",
			},
			args: args{
				ctx:   ctx,
				token: token,
			},
			want: u,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &loginUseCase{
				getRepository: tt.fields.getRepository,
				secret:        tt.fields.secret,
			}
			got, err := l.VerifyToken(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("loginUseCase.VerifyToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.ID, tt.want.ID) {
				t.Errorf("loginUseCase.VerifyToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

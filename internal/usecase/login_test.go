package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/TcMits/ent-clean-template/internal/repository"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	jwtKit "github.com/golang-jwt/jwt/v4"
)

func TestNewLoginUseCase(t *testing.T) {
	type args struct {
		repository repository.LoginRepository
		secret     string
	}
	tests := []struct {
		name string
		args args
		want LoginUseCase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := NewLoginUseCase(tt.args.repository, tt.args.secret); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewLoginUseCase() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_loginUseCase_getUserMapClaims(t *testing.T) {
	type fields struct {
		repository repository.LoginRepository
		secret     string
	}
	type args struct {
		user *model.User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   jwtKit.MapClaims
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		l := &loginUseCase{
			repository: tt.fields.repository,
			secret:     tt.fields.secret,
		}
		if got := l.getUserMapClaims(tt.args.user); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. loginUseCase.getUserMapClaims() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_loginUseCase_getUserFromMapClaims(t *testing.T) {
	type fields struct {
		repository repository.LoginRepository
		secret     string
	}
	type args struct {
		ctx          context.Context
		jwtMapClaims jwtKit.MapClaims
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		l := &loginUseCase{
			repository: tt.fields.repository,
			secret:     tt.fields.secret,
		}
		got, err := l.getUserFromMapClaims(tt.args.ctx, tt.args.jwtMapClaims)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. loginUseCase.getUserFromMapClaims() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. loginUseCase.getUserFromMapClaims() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_loginUseCase_createAccessToken(t *testing.T) {
	type fields struct {
		repository repository.LoginRepository
		secret     string
	}
	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		l := &loginUseCase{
			repository: tt.fields.repository,
			secret:     tt.fields.secret,
		}
		got, err := l.createAccessToken(tt.args.user)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. loginUseCase.createAccessToken() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. loginUseCase.createAccessToken() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_loginUseCase_createRefreshToken(t *testing.T) {
	type fields struct {
		repository repository.LoginRepository
		secret     string
	}
	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *useCaseModel.RefreshTokenInput
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		l := &loginUseCase{
			repository: tt.fields.repository,
			secret:     tt.fields.secret,
		}
		got, err := l.createRefreshToken(tt.args.user)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. loginUseCase.createRefreshToken() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. loginUseCase.createRefreshToken() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_loginUseCase_parseAccessToken(t *testing.T) {
	type fields struct {
		repository repository.LoginRepository
		secret     string
	}
	type args struct {
		ctx   context.Context
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		l := &loginUseCase{
			repository: tt.fields.repository,
			secret:     tt.fields.secret,
		}
		got, err := l.parseAccessToken(tt.args.ctx, tt.args.token)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. loginUseCase.parseAccessToken() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. loginUseCase.parseAccessToken() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_loginUseCase_parseRefreshToken(t *testing.T) {
	type fields struct {
		repository repository.LoginRepository
		secret     string
	}
	type args struct {
		ctx               context.Context
		refreshTokenInput *useCaseModel.RefreshTokenInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		l := &loginUseCase{
			repository: tt.fields.repository,
			secret:     tt.fields.secret,
		}
		got, err := l.parseRefreshToken(tt.args.ctx, tt.args.refreshTokenInput)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. loginUseCase.parseRefreshToken() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. loginUseCase.parseRefreshToken() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_loginUseCase_Login(t *testing.T) {
	type fields struct {
		repository repository.LoginRepository
		secret     string
	}
	type args struct {
		ctx        context.Context
		loginInput *useCaseModel.LoginInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *useCaseModel.JWTAuthenticatedPayload
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		l := &loginUseCase{
			repository: tt.fields.repository,
			secret:     tt.fields.secret,
		}
		got, err := l.Login(tt.args.ctx, tt.args.loginInput)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. loginUseCase.Login() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. loginUseCase.Login() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_loginUseCase_RefreshToken(t *testing.T) {
	type fields struct {
		repository repository.LoginRepository
		secret     string
	}
	type args struct {
		ctx               context.Context
		refreshTokenInput *useCaseModel.RefreshTokenInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		l := &loginUseCase{
			repository: tt.fields.repository,
			secret:     tt.fields.secret,
		}
		got, err := l.RefreshToken(tt.args.ctx, tt.args.refreshTokenInput)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. loginUseCase.RefreshToken() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. loginUseCase.RefreshToken() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_loginUseCase_VerifyToken(t *testing.T) {
	type fields struct {
		repository repository.LoginRepository
		secret     string
	}
	type args struct {
		ctx   context.Context
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		l := &loginUseCase{
			repository: tt.fields.repository,
			secret:     tt.fields.secret,
		}
		got, err := l.VerifyToken(tt.args.ctx, tt.args.token)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. loginUseCase.VerifyToken() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. loginUseCase.VerifyToken() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

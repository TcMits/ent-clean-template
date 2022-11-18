package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	jwtKit "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"github.com/TcMits/ent-clean-template/copygen"
	"github.com/TcMits/ent-clean-template/internal/repository"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/TcMits/ent-clean-template/pkg/tool/jwt"
	"github.com/TcMits/ent-clean-template/pkg/tool/password"
)

const (
	_defaultAccessTokenTimeOut  = time.Minute * 15
	_defaultRefreshTokenTimeOut = time.Hour * 24 * 7
)

const (
	_idFieldName         = "id"
	_keyFieldName        = "key"
	_refreshKeyFieldName = "refresh_key"
)

var (
	_wrapInvalidUsernameError = func(err error) error {
		return model.NewTranslatableError(
			fmt.Errorf("internal.usecase.LoginUseCase.Login: %w", err),
			_incorrectLoginInputMessage,
			AuthenticationError,
			nil,
		)
	}
	_wrapInvalidPasswordError = func(err error) error {
		return model.NewTranslatableError(
			fmt.Errorf("internal.usecase.LoginUseCase.Login: %w", err),
			_incorrectLoginInputMessage,
			AuthenticationError,
			nil,
		)
	}
	_wrapFailedAccessTokenCreation = func(err error) error {
		return model.NewTranslatableError(
			fmt.Errorf("internal.usecase.LoginUseCase.Login: %w", err),
			_canNotLoginNowMessage,
			InternalServerError,
			nil,
		)
	}
	_wrapFailedRefreshTokenCreation = func(err error) error {
		return model.NewTranslatableError(
			fmt.Errorf("internal.usecase.LoginUseCase.Login: %w", err),
			_canNotLoginNowMessage,
			InternalServerError,
			nil,
		)
	}
	_wrapInvalidRefreshToken = func(err error) error {
		return model.NewTranslatableError(
			fmt.Errorf("internal.usecase.LoginUseCase.RefreshToken: %w", err),
			_authenticationFailedMessage,
			AuthenticationError,
			nil,
		)
	}
	_wrapInvalidAccessToken = func(err error) error {
		return model.NewTranslatableError(
			fmt.Errorf("internal.usecase.LoginUseCase.VerifyToken: %w", err),
			_authenticationFailedMessage,
			AuthenticationError,
			nil,
		)
	}
)

type loginUseCase struct {
	getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput]
	secret        string
}

func NewLoginUseCase(
	getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput],
	secret string,
) LoginUseCase[*useCaseModel.LoginInput, *useCaseModel.JWTAuthenticatedPayload, *useCaseModel.RefreshTokenInput, *model.User] {
	if getRepository == nil {
		panic("getRepository is required")
	}
	return &loginUseCase{getRepository: getRepository, secret: secret}
}

func (*loginUseCase) getUserMapClaims(user *model.User) jwtKit.MapClaims {
	return jwtKit.MapClaims{
		_idFieldName:  user.ID.String(),
		"email":       user.Email,
		_keyFieldName: user.JwtTokenKey,
	}
}

func (l *loginUseCase) getUserFromMapClaims(
	ctx context.Context,
	jwtMapClaims jwtKit.MapClaims,
) (*model.User, error) {
	strId, ok := jwtMapClaims[_idFieldName].(string)
	if !ok {
		strId = ""
	}
	id, err := uuid.Parse(strId)
	if err != nil {
		return nil, err
	}
	isActive := true
	user, err := l.getRepository.Get(ctx, &model.UserWhereInput{
		ID: &id, IsActive: &isActive,
	})
	if err != nil {
		return nil, err
	}
	key, ok := jwtMapClaims[_keyFieldName].(string)
	if !ok || user.JwtTokenKey != key {
		return nil, errors.New("Invalid token key")
	}
	return user, nil
}

func (l *loginUseCase) createAccessToken(user *model.User) (string, error) {
	payload := l.getUserMapClaims(user)
	return jwt.NewToken(payload, l.secret, _defaultAccessTokenTimeOut)
}

func (l *loginUseCase) createRefreshToken(
	user *model.User,
) (*useCaseModel.RefreshTokenInput, error) {
	refreshKey, err := jwt.NewToken(
		jwtKit.MapClaims{}, l.secret, _defaultRefreshTokenTimeOut,
	)
	if err != nil {
		return nil, err
	}
	payload := l.getUserMapClaims(user)
	refreshToken, err := jwt.NewToken(payload, refreshKey+l.secret, _defaultRefreshTokenTimeOut)
	if err != nil {
		return nil, err
	}
	return &useCaseModel.RefreshTokenInput{
		RefreshToken: refreshToken,
		RefreshKey:   refreshKey,
	}, nil
}

func (l *loginUseCase) parseAccessToken(ctx context.Context, token string) (*model.User, error) {
	payload, err := jwt.ParseJWT(token, l.secret)
	if err != nil {
		return nil, err
	}

	return l.getUserFromMapClaims(ctx, payload)
}

func (l *loginUseCase) parseRefreshToken(
	ctx context.Context, refreshTokenInput *useCaseModel.RefreshTokenInput,
) (*model.User, error) {
	payload, err := jwt.ParseJWT(
		refreshTokenInput.RefreshToken,
		refreshTokenInput.RefreshKey+l.secret,
	)
	if err != nil {
		return nil, err
	}

	return l.getUserFromMapClaims(ctx, payload)
}

func (l *loginUseCase) Login(
	ctx context.Context,
	loginInput *useCaseModel.LoginInput,
) (*useCaseModel.JWTAuthenticatedPayload, error) {
	isActive := true
	userWhereInput := &model.UserWhereInput{IsActive: &isActive}
	copygen.LoginInputToUserWhereInput(userWhereInput, loginInput)
	user, err := l.getRepository.Get(ctx, userWhereInput)

	if err != nil {
		return nil, _wrapInvalidUsernameError(err)
	}

	if err := password.ValidatePassword(user.Password, loginInput.Password); err != nil {
		return nil, _wrapInvalidPasswordError(err)
	}

	accessToken, err := l.createAccessToken(user)
	if err != nil {
		return nil, _wrapFailedAccessTokenCreation(err)
	}

	refreshTokenInput, err := l.createRefreshToken(user)
	if err != nil {
		return nil, _wrapFailedRefreshTokenCreation(err)
	}

	return &useCaseModel.JWTAuthenticatedPayload{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenInput.RefreshToken,
		RefreshKey:   refreshTokenInput.RefreshKey,
	}, nil
}

func (l *loginUseCase) RefreshToken(
	ctx context.Context,
	refreshTokenInput *useCaseModel.RefreshTokenInput,
) (string, error) {

	user, err := l.parseRefreshToken(ctx, refreshTokenInput)
	if err != nil {
		return "", _wrapInvalidRefreshToken(err)
	}

	accessToken, err := l.createAccessToken(user)
	if err != nil {
		return "", _wrapFailedAccessTokenCreation(err)
	}

	return accessToken, nil
}

func (l *loginUseCase) VerifyToken(ctx context.Context, token string) (*model.User, error) {
	user, err := l.parseAccessToken(ctx, token)
	if err != nil {
		return nil, _wrapInvalidAccessToken(err)
	}
	return user, nil
}

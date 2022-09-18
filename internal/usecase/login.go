package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/TcMits/ent-clean-template/internal/repository"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/TcMits/ent-clean-template/pkg/tool/jwt"
	jwtKit "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const (
	_defaultAccessTokenTimeOut  = time.Minute * 15
	_defaultRefreshTokenTimeOut = time.Hour * 24 * 7
)

const (
	idFieldName         = "id"
	keyFieldName        = "key"
	refreshKeyFieldName = "refresh_key"
)

var (
	_wrapInvalidLoginInputError = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf("loginUseCase - Login - l.repository.Login: %w", err),
			"internal.usecase.login.loginUseCase.Login.InvalidLoginInput",
			"Your username or password is incorrect",
			AuthenticationError,
		)
	}
	_wrapFailedAccessTokenCreation = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf("loginUseCase - Login - l.createAccessToken: %w", err),
			"internal.usecase.login.loginUseCase.Login.FailedTokenCreation",
			"Can't login now",
			InternalServerError,
		)
	}
	_wrapFailedRefreshTokenCreation = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf("loginUseCase - Login - l.createAccessToken: %w", err),
			"internal.usecase.login.loginUseCase.Login.FailedTokenCreation",
			"Can't login now",
			InternalServerError,
		)
	}
	_wrapInvalidRefreshToken = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf("loginUseCase - RefreshToken - l.parseRefreshToken: %w", err),
			"internal.usecase.login.loginUseCase.Login.InvalidRefreshToken",
			"Authentication failed",
			AuthenticationError,
		)
	}
	_wrapInvalidAccessToken = func(err error) error {
		return useCaseModel.NewUseCaseError(
			fmt.Errorf("loginUseCase - RefreshToken - l.parseAccessToken: %w", err),
			"internal.usecase.login.loginUseCase.Login.InvalidAccessToken",
			"Authentication failed",
			AuthenticationError,
		)
	}
)

type loginUseCase struct {
	repository    repository.LoginRepository[*model.User, *model.UserWhereInput, *useCaseModel.LoginInput]
	getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput]
	secret        string
}

func NewLoginUseCase(
	repository repository.LoginRepository[*model.User, *model.UserWhereInput, *useCaseModel.LoginInput],
	getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput],
	secret string,
) LoginUseCase[*useCaseModel.LoginInput, *useCaseModel.JWTAuthenticatedPayload, *useCaseModel.RefreshTokenInput, *model.User] {
	if repository == nil {
		panic("repository is required")
	}
	if getRepository == nil {
		panic("getRepository is required")
	}
	return &loginUseCase{repository: repository, getRepository: getRepository, secret: secret}
}

func (*loginUseCase) getUserMapClaims(user *model.User) jwtKit.MapClaims {
	return jwtKit.MapClaims{
		idFieldName:  user.ID.String(),
		"email":      user.Email,
		keyFieldName: user.JwtTokenKey,
	}
}

func (l *loginUseCase) getUserFromMapClaims(ctx context.Context, jwtMapClaims jwtKit.MapClaims) (*model.User, error) {
	strId, ok := jwtMapClaims[idFieldName].(string)
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
	key, ok := jwtMapClaims[keyFieldName].(string)
	if !ok || user.JwtTokenKey != key {
		return nil, errors.New("loginUseCase - getUserFromMapClaims: Invalid token key")
	}
	return user, nil
}

func (l *loginUseCase) createAccessToken(user *model.User) (string, error) {
	return jwt.NewToken(l.getUserMapClaims(user), l.secret, _defaultAccessTokenTimeOut)
}

func (l *loginUseCase) createRefreshToken(user *model.User) (*useCaseModel.RefreshTokenInput, error) {
	refreshKey, err := jwt.NewToken(
		jwtKit.MapClaims{}, l.secret, _defaultRefreshTokenTimeOut,
	)
	if err != nil {
		return nil, err
	}
	userMapClaims := l.getUserMapClaims(user)
	userMapClaims[refreshKeyFieldName] = refreshKey
	refreshToken, err := jwt.NewToken(
		userMapClaims, l.secret, _defaultRefreshTokenTimeOut)
	if err != nil {
		return nil, err
	}
	return &useCaseModel.RefreshTokenInput{
		RefreshToken: refreshToken,
		RefreshKey:   refreshKey,
	}, nil
}

func (l *loginUseCase) parseAccessToken(ctx context.Context, token string) (*model.User, error) {
	jwtMapClaims, err := jwt.ParseJWT(token, l.secret)
	if err != nil {
		return nil, err
	}
	return l.getUserFromMapClaims(ctx, jwtMapClaims)
}

func (l *loginUseCase) parseRefreshToken(
	ctx context.Context, refreshTokenInput *useCaseModel.RefreshTokenInput) (*model.User, error) {
	jwtMapClaims, err := jwt.ParseJWT(
		refreshTokenInput.RefreshToken,
		l.secret,
	)
	if err != nil {
		return nil, err
	}
	key, ok := jwtMapClaims[refreshKeyFieldName].(string)
	if !ok || refreshTokenInput.RefreshKey != key {
		return nil, errors.New("loginUseCase - parseRefreshToken: Invalid token key")
	}
	return l.getUserFromMapClaims(ctx, jwtMapClaims)
}

func (l *loginUseCase) Login(
	ctx context.Context,
	loginInput *useCaseModel.LoginInput,
) (*useCaseModel.JWTAuthenticatedPayload, error) {
	user, err := l.repository.Login(ctx, loginInput)
	if err != nil {
		return nil, _wrapInvalidLoginInputError(err)
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

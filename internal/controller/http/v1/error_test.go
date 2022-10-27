package v1

import (
	"errors"
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/TcMits/ent-clean-template/internal/testutils"
	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
)

type InputStruct struct {
	Message string `validate:"min=2"`
}

type InputStructWithErrMessage struct {
	Message string `validate:"min=2"`
}

func (*InputStructWithErrMessage) GetErrorMessageFromStructField(_ error) *i18n.Message {
	return &i18n.Message{
		ID:    "Message",
		Other: "Message",
	}
}

type TestErrorWithCode struct{}

func (TestErrorWithCode) Error() string {
	return "test"
}

func (TestErrorWithCode) Code() string {
	return "test"
}

func Test_getCodeFromError(t *testing.T) {
	type args struct {
		err error
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "HaveCode",
			args: args{
				err: new(TestErrorWithCode),
			},
			want: "test",
		},
		{
			name: "UnknownCode",
			args: args{
				err: errors.New("test"),
			},
			want: _unknownError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCodeFromError(tt.args.err); got != tt.want {
				t.Errorf("getCodeFromError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getStatusCodeFromCode(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "PermissionDenied",
			args: args{code: usecase.PermissionDeniedError},
			want: iris.StatusForbidden,
		},
		{
			name: "AuthenticationError",
			args: args{code: usecase.AuthenticationError},
			want: iris.StatusUnauthorized,
		},
		{
			name: "InternalServerError",
			args: args{code: usecase.InternalServerError},
			want: iris.StatusInternalServerError,
		},
		{
			name: "InternalServerError",
			args: args{code: usecase.InternalServerError},
			want: iris.StatusInternalServerError,
		},
		{
			name: "_unknownError",
			args: args{code: _unknownError},
			want: iris.StatusInternalServerError,
		},
		{
			name: "ValidationError",
			args: args{code: usecase.ValidationError},
			want: iris.StatusBadRequest,
		},
		{
			name: "_useCaseInputValidationError",
			args: args{code: _useCaseInputValidationError},
			want: iris.StatusBadRequest,
		},
		{
			name: "NotFoundError",
			args: args{code: usecase.NotFoundError},
			want: iris.StatusNotFound,
		},
		{
			name: "DBError",
			args: args{code: usecase.DBError},
			want: iris.StatusNotAcceptable,
		},
		{
			name: "DefaultError",
			args: args{code: "test"},
			want: iris.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStatusCodeFromCode(tt.args.code); got != tt.want {
				t.Errorf("getStatusCodeFromCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_logError(t *testing.T) {
	type args struct {
		err  error
		code string
		l    logger.Interface
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "PermissionDenied",
			args: args{
				err:  errors.New("test"),
				l:    testutils.NullLogger{},
				code: usecase.PermissionDeniedError,
			},
		},
		{
			name: "AuthenticationError",
			args: args{
				err:  errors.New("test"),
				l:    testutils.NullLogger{},
				code: usecase.AuthenticationError,
			},
		},
		{
			name: "InternalServerError",
			args: args{
				err:  errors.New("test"),
				l:    testutils.NullLogger{},
				code: usecase.InternalServerError,
			},
		},
		{
			name: "InternalServerError",
			args: args{
				err:  errors.New("test"),
				l:    testutils.NullLogger{},
				code: usecase.InternalServerError,
			},
		},
		{
			name: "_unknownError",
			args: args{
				err:  errors.New("test"),
				l:    testutils.NullLogger{},
				code: _unknownError,
			},
		},
		{
			name: "ValidationError",
			args: args{
				err:  errors.New("test"),
				l:    testutils.NullLogger{},
				code: usecase.ValidationError,
			},
		},
		{
			name: "_useCaseInputValidationError",
			args: args{
				err:  errors.New("test"),
				l:    testutils.NullLogger{},
				code: _useCaseInputValidationError,
			},
		},
		{
			name: "NotFoundError",
			args: args{
				err:  errors.New("test"),
				l:    testutils.NullLogger{},
				code: usecase.NotFoundError,
			},
		},
		{
			name: "DBError",
			args: args{
				err:  errors.New("test"),
				l:    testutils.NullLogger{},
				code: usecase.DBError,
			},
		},
		{
			name: "DefaultError",
			args: args{
				err:  errors.New("test"),
				l:    testutils.NullLogger{},
				code: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logError(tt.args.err, tt.args.code, tt.args.l)
		})
	}
}

func Test_handleError(t *testing.T) {
	type args struct {
		err error
		l   logger.Interface
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "NormalError",
			args: args{
				err: errors.New("test"),
				l:   testutils.NullLogger{},
			},
		},
		{
			name: "UsecaseError",
			args: args{
				err: model.NewTranslatableError(
					errors.New(""), &i18n.Message{
						ID:    "test",
						Other: "test",
					}, usecase.AuthenticationError, nil,
				),
				l: testutils.NullLogger{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			routeHandler := func(irisCtx iris.Context) {
				handleError(irisCtx, tt.args.err, tt.args.l)
			}
			handler := NewHandler()
			handler.Get("/", routeHandler)
			httptest.New(t, handler).GET("/").Expect()
		})
	}
}

func Test_translatableErrorFromValidationErrors(t *testing.T) {
	type args struct {
		inputStructure any
		errs           validator.ValidationErrors
		tr             model.TranslateFunc
	}
	validate := validator.New()

	errs := validate.Struct(&InputStruct{Message: "1"}).(validator.ValidationErrors)
	errs_with_message := validate.Struct(&InputStructWithErrMessage{Message: "1"}).(validator.ValidationErrors)

	tests := []struct {
		name string
		args args
		want *model.TranslatableError
	}{
		{
			name: "Errs",
			args: args{
				inputStructure: new(InputStruct),
				errs:           errs,
			},
			want: model.NewTranslatableError(
				errs,
				_oneOrMoreFieldsFailedToBeValidatedMessage,
				_useCaseInputValidationError,
				nil,
			),
		},
		{
			name: "ErrsWithMessage",
			args: args{
				inputStructure: new(InputStructWithErrMessage),
				errs:           errs_with_message,
			},
			want: model.NewTranslatableError(
				errs_with_message[0],
				&i18n.Message{
					ID:    "Message",
					Other: "Message",
				},
				_useCaseInputValidationError,
				nil,
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := translatableErrorFromValidationErrors(
				tt.args.inputStructure,
				&tt.args.errs,
			)
			if !reflect.DeepEqual(got.Error(), tt.want.Error()) {
				t.Errorf(
					"translatableErrorFromValidationErrors() = %v, want %v",
					got.Error(),
					tt.want.Error(),
				)
			}
		})
	}
}

func Test_handleBindingError(t *testing.T) {
	type args struct {
		err                error
		l                  logger.Interface
		input              any
		wrapTranslateError func(error) error
	}

	validate := validator.New()
	errs := validate.Struct(&InputStruct{Message: "1"}).(validator.ValidationErrors)

	tests := []struct {
		name string
		args args
	}{
		{
			name: "WithValidationErrors",
			args: args{
				err:                errs,
				l:                  testutils.NullLogger{},
				input:              new(struct{}),
				wrapTranslateError: func(err error) error { return errors.New("") },
			},
		},
		{
			name: "WithNormalError",
			args: args{
				err:                errors.New(""),
				l:                  testutils.NullLogger{},
				input:              new(struct{}),
				wrapTranslateError: func(err error) error { return errors.New("") },
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			routeHandler := func(irisCtx iris.Context) {
				handleBindingError(
					irisCtx,
					tt.args.err,
					tt.args.l,
					tt.args.input,
					tt.args.wrapTranslateError,
				)
			}
			handler := NewHandler()
			handler.Get("/", routeHandler)
			httptest.New(t, handler).GET("/").Expect()
		})
	}
}

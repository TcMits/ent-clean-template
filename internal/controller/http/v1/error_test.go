package v1

import (
	"errors"
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"

	"github.com/TcMits/ent-clean-template/internal/testutils"
	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
)

type InputStruct struct {
	Message string `validate:"min=2"`
}

type InputStructWithErrMessage struct {
	Message string `validate:"min=2"`
}

func (*InputStructWithErrMessage) GetErrorMessageFromStructField(key string) (string, string) {
	return "Message", "Message"
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
			name: "_uscaseInputValidationError",
			args: args{code: _uscaseInputValidationError},
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
			name: "_uscaseInputValidationError",
			args: args{
				err:  errors.New("test"),
				l:    testutils.NullLogger{},
				code: _uscaseInputValidationError,
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
				err: useCaseModel.NewUseCaseError(
					errors.New(""), "test", "test", usecase.AuthenticationError,
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
	tr := func(string, ...any) string { return "" }

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
				tr:             tr,
			},
			want: model.NewTranslatableError(
				errs,
				_defaultInvalidErrorTranslateKey,
				tr,
				_defaultInvalidErrorMessage,
				_uscaseInputValidationError,
			),
		},
		{
			name: "ErrsWithMessage",
			args: args{
				inputStructure: new(InputStructWithErrMessage),
				errs:           errs_with_message,
				tr:             tr,
			},
			want: model.NewTranslatableError(
				errs_with_message[0],
				"Message",
				tr,
				"Message",
				_uscaseInputValidationError,
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := translatableErrorFromValidationErrors(
				tt.args.inputStructure,
				tt.args.errs,
				tt.args.tr,
			)
			if !reflect.DeepEqual(got.Error(), tt.want.Error()) {
				t.Errorf(
					"translatableErrorFromValidationErrors() = %v, want %v",
					got.Error(),
					tt.want.Error(),
				)
			}
			if !reflect.DeepEqual(got.Key(), tt.want.Key()) {
				t.Errorf(
					"translatableErrorFromValidationErrors() = %v, want %v",
					got.Key(),
					tt.want.Key(),
				)
			}
			if !reflect.DeepEqual(got.DefaultError(), tt.want.DefaultError()) {
				t.Errorf(
					"translatableErrorFromValidationErrors() = %v, want %v",
					got.DefaultError(),
					tt.want.DefaultError(),
				)
			}
			if !reflect.DeepEqual(got.Code(), tt.want.Code()) {
				t.Errorf(
					"translatableErrorFromValidationErrors() = %v, want %v",
					got.Code(),
					tt.want.Code(),
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
		wrapTranslateError func(model.TranslateFunc, error) error
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
				wrapTranslateError: func(tf model.TranslateFunc, err error) error { return errors.New("") },
			},
		},
		{
			name: "WithNormalError",
			args: args{
				err:                errors.New(""),
				l:                  testutils.NullLogger{},
				input:              new(struct{}),
				wrapTranslateError: func(tf model.TranslateFunc, err error) error { return errors.New("") },
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

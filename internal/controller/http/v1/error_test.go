package v1

import (
	"errors"
	"testing"

	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/kataras/iris/v12"
)

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
			want: UnknownError,
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
			name: "UnknownError",
			args: args{code: UnknownError},
			want: iris.StatusInternalServerError,
		},
		{
			name: "ValidationError",
			args: args{code: usecase.ValidationError},
			want: iris.StatusBadRequest,
		},
		{
			name: "UscaseInputValidationError",
			args: args{code: UscaseInputValidationError},
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

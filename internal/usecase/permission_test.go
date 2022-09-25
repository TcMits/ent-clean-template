package usecase

import (
	"context"
	"errors"
	"testing"
)

func Test_basePermissionCheckerUseCase_Check(t *testing.T) {
	type fields struct {
		checkFunc func(context.Context, any) error
	}
	type args struct {
		ctx context.Context
		u   any
	}
	ctx := context.Background()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "WantErrTrue",
			fields: fields{
				checkFunc: func(ctx context.Context, a any) error { return errors.New("") },
			},
			args:    args{ctx: ctx, u: nil},
			wantErr: true,
		},
		{
			name:    "WantErrFalse",
			fields:  fields{checkFunc: func(ctx context.Context, a any) error { return nil }},
			args:    args{ctx: ctx, u: nil},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &basePermissionCheckerUseCase[any]{
				checkFunc: tt.fields.checkFunc,
			}
			err := l.Check(tt.args.ctx, tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"basePermissionCheckerUseCase.Check() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}

func Test_basePermissionCheckerUseCase_And(t *testing.T) {
	type fields struct {
		checkFunc func(context.Context, any) error
	}
	type args struct {
		checker UserPermissionCheckerUseCase[any]
	}
	ctx := context.Background()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "WantErrTrue",
			fields: fields{
				checkFunc: func(ctx context.Context, a any) error { return errors.New("") },
			},
			args: args{checker: &basePermissionCheckerUseCase[any]{
				checkFunc: func(ctx context.Context, a any) error { return nil },
			}},
			wantErr: true,
		},
		{
			name:   "WantErrFalse",
			fields: fields{checkFunc: func(ctx context.Context, a any) error { return nil }},
			args: args{checker: &basePermissionCheckerUseCase[any]{
				checkFunc: func(ctx context.Context, a any) error { return nil },
			}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &basePermissionCheckerUseCase[any]{
				checkFunc: tt.fields.checkFunc,
			}
			got := l.And(tt.args.checker)
			if err := got.Check(ctx, nil); (err != nil) != tt.wantErr {
				t.Errorf(
					"basePermissionCheckerUseCase.Check() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}

func Test_basePermissionCheckerUseCase_Or(t *testing.T) {
	type fields struct {
		checkFunc func(context.Context, any) error
	}
	type args struct {
		checker UserPermissionCheckerUseCase[any]
	}
	ctx := context.Background()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "WantErrFalse",
			fields: fields{
				checkFunc: func(ctx context.Context, a any) error { return errors.New("") },
			},
			args: args{checker: &basePermissionCheckerUseCase[any]{
				checkFunc: func(ctx context.Context, a any) error { return nil },
			}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &basePermissionCheckerUseCase[any]{
				checkFunc: tt.fields.checkFunc,
			}
			got := l.Or(tt.args.checker)
			if err := got.Check(ctx, nil); (err != nil) != tt.wantErr {
				t.Errorf(
					"basePermissionCheckerUseCase.Check() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}

func Test_NewAllowAnyPermissionChecker(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Success",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAllowAnyPermissionChecker[any]()
			if err := got.Check(ctx, nil); (err != nil) != tt.wantErr {
				t.Errorf(
					"basePermissionCheckerUseCase.Check() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}

func Test_NewDisallowAnyPermissionChecker(t *testing.T) {
	type args struct {
		err error
	}

	ctx := context.Background()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Success",
			args:    args{err: errors.New("")},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDisallowAnyPermissionChecker[any](tt.args.err)
			if err := got.Check(ctx, nil); (err != nil) != tt.wantErr {
				t.Errorf(
					"basePermissionCheckerUseCase.Check() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}

func Test_NewDisallowZeroPermissionChecker(t *testing.T) {
	type args struct {
		err error
	}

	ctx := context.Background()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Success",
			args:    args{err: errors.New("")},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDisallowZeroPermissionChecker[*struct{}](tt.args.err)
			if err := got.Check(ctx, nil); (err != nil) != tt.wantErr {
				t.Errorf(
					"basePermissionCheckerUseCase.Check() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}

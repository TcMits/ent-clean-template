package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/TcMits/ent-clean-template/internal/repository"
)

func Test_getModelUseCase_Get(t *testing.T) {
	type fields struct {
		repository           repository.GetModelRepository[*struct{}, *struct{}]
		toRepoWhereInputFunc ConvertFunc[*struct{}, *struct{}]
		wrapGetErrorFunc     func(error) error
	}
	type args struct {
		ctx   context.Context
		input *struct{}
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := repository.NewMockGetModelRepository[*struct{}, *struct{}](ctrl)

	repo.EXPECT().Get(
		gomock.Eq(ctx), gomock.Eq(new(struct{})),
	).Return(
		new(struct{}), nil,
	).AnyTimes()

	repo.EXPECT().Get(
		gomock.Eq(ctx), gomock.Nil(),
	).Return(
		nil, errors.New(""),
	).AnyTimes()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *struct{}
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				repository: repo,
				toRepoWhereInputFunc: func(c context.Context, uwi *struct{}) (*struct{}, error) {
					return uwi, nil
				},
				wrapGetErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: new(struct{}),
			},
			want: new(struct{}),
		},
		{
			name: "WhereInputFuncError",
			fields: fields{
				repository: repo,
				toRepoWhereInputFunc: func(c context.Context, uwi *struct{}) (*struct{}, error) {
					return nil, errors.New("test")
				},
				wrapGetErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: new(struct{}),
			},
			wantErr: true,
		},
		{
			name: "GetError",
			fields: fields{
				repository: repo,
				toRepoWhereInputFunc: func(c context.Context, uwi *struct{}) (*struct{}, error) {
					return uwi, nil
				},
				wrapGetErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &getModelUseCase[*struct{}, *struct{}, *struct{}]{
				repository:           tt.fields.repository,
				toRepoWhereInputFunc: tt.fields.toRepoWhereInputFunc,
				wrapGetErrorFunc:     tt.fields.wrapGetErrorFunc,
			}
			got, err := l.Get(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("getModelUseCase.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getModelUseCase.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countModelUseCase_Count(t *testing.T) {
	type fields struct {
		repository           repository.CountModelRepository[*struct{}]
		toRepoWhereInputFunc ConvertFunc[*struct{}, *struct{}]
		wrapCountErrorFunc   func(error) error
	}
	type args struct {
		ctx   context.Context
		input *struct{}
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := repository.NewMockCountModelRepository[*struct{}](ctrl)

	repo.EXPECT().Count(
		gomock.Eq(ctx), gomock.Eq(new(struct{})),
	).Return(
		1, nil,
	).AnyTimes()

	repo.EXPECT().Count(
		gomock.Eq(ctx), gomock.Nil(),
	).Return(
		0, errors.New(""),
	).AnyTimes()

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
				repository: repo,
				toRepoWhereInputFunc: func(c context.Context, uwi *struct{}) (*struct{}, error) {
					return uwi, nil
				},
				wrapCountErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: new(struct{}),
			},
			want: 1,
		},
		{
			name: "WhereInputFuncError",
			fields: fields{
				repository: repo,
				toRepoWhereInputFunc: func(c context.Context, uwi *struct{}) (*struct{}, error) {
					return nil, errors.New("test")
				},
				wrapCountErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: new(struct{}),
			},
			wantErr: true,
		},
		{
			name: "CountError",
			fields: fields{
				repository: repo,
				toRepoWhereInputFunc: func(c context.Context, uwi *struct{}) (*struct{}, error) {
					return uwi, nil
				},
				wrapCountErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &countModelUseCase[*struct{}, *struct{}]{
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
		repository           repository.ListModelRepository[*struct{}, *struct{}, *struct{}]
		toRepoWhereInputFunc ConvertFunc[*struct{}, *struct{}]
		toRepoOrderInputFunc ConvertFunc[*struct{}, *struct{}]
		wrapListErrorFunc    func(error) error
	}
	type args struct {
		ctx        context.Context
		limit      *int
		offset     *int
		orderInput *struct{}
		whereInput *struct{}
	}
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := repository.NewMockListModelRepository[*struct{}, *struct{}, *struct{}](ctrl)
	limit := 10
	offset := 0

	repo.EXPECT().List(
		gomock.Eq(ctx), gomock.Eq(&limit), gomock.Eq(&offset), gomock.Eq(new(struct{})), gomock.Eq(new(struct{})),
	).Return(
		make([]*struct{}, 10), nil,
	).AnyTimes()

	repo.EXPECT().List(
		gomock.Eq(ctx), gomock.Eq(&limit), gomock.Eq(&offset), gomock.Eq(new(struct{})), gomock.Nil(),
	).Return(
		nil, errors.New(""),
	).AnyTimes()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*struct{}
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				repository: repo,
				toRepoWhereInputFunc: func(c context.Context, uwi *struct{}) (*struct{}, error) {
					return uwi, nil
				},
				toRepoOrderInputFunc: func(c context.Context, uoi *struct{}) (*struct{}, error) {
					return uoi, nil
				},
				wrapListErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:        ctx,
				limit:      &limit,
				offset:     &offset,
				orderInput: new(struct{}),
				whereInput: new(struct{}),
			},
			want: make([]*struct{}, 10),
		},
		{
			name: "WhereInputFuncError",
			fields: fields{
				repository: repo,
				toRepoWhereInputFunc: func(c context.Context, uwi *struct{}) (*struct{}, error) {
					return nil, errors.New("test")
				},
				toRepoOrderInputFunc: func(c context.Context, uoi *struct{}) (*struct{}, error) {
					return uoi, nil
				},
				wrapListErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:        ctx,
				limit:      &limit,
				offset:     &offset,
				orderInput: new(struct{}),
				whereInput: new(struct{}),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "OrderInputFuncError",
			fields: fields{
				repository: repo,
				toRepoWhereInputFunc: func(c context.Context, uwi *struct{}) (*struct{}, error) {
					return uwi, nil
				},
				toRepoOrderInputFunc: func(c context.Context, uoi *struct{}) (*struct{}, error) {
					return nil, errors.New("test")
				},
				wrapListErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:        ctx,
				limit:      &limit,
				offset:     &offset,
				orderInput: new(struct{}),
				whereInput: new(struct{}),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ListError",
			fields: fields{
				repository: repo,
				toRepoWhereInputFunc: func(c context.Context, uwi *struct{}) (*struct{}, error) {
					return uwi, nil
				},
				toRepoOrderInputFunc: func(c context.Context, uoi *struct{}) (*struct{}, error) {
					return uoi, nil
				},
				wrapListErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:        ctx,
				limit:      &limit,
				offset:     &offset,
				orderInput: new(struct{}),
				whereInput: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &listModelUseCase[*struct{}, *struct{}, *struct{}, *struct{}, *struct{}]{
				repository:           tt.fields.repository,
				toRepoWhereInputFunc: tt.fields.toRepoWhereInputFunc,
				toRepoOrderInputFunc: tt.fields.toRepoOrderInputFunc,
				wrapListErrorFunc:    tt.fields.wrapListErrorFunc,
			}
			got, err := l.List(
				tt.args.ctx,
				tt.args.limit,
				tt.args.offset,
				tt.args.orderInput,
				tt.args.whereInput,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf("listModelUseCase.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listModelUseCase.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

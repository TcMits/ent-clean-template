package usecase

import (
	"context"
	"net/url"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_modelSerializerField_Serialize(t *testing.T) {
	type args struct {
		ctx      context.Context
		instance *struct{}
	}
	ctx := context.Background()

	tests := []struct {
		name string
		args args
		sf   modelSerializerField[*struct{}, *struct{}]
		want *struct{}
	}{
		{
			name: "NormalSF",
			args: args{
				ctx:      ctx,
				instance: new(struct{}),
			},
			sf:   func(ctx context.Context, s *struct{}) *struct{} { return s },
			want: new(struct{}),
		},
		{
			name: "WithoutURL",
			args: args{
				ctx:      ctx,
				instance: new(struct{}),
			},
			sf:   nil,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sf.Serialize(tt.args.ctx, tt.args.instance); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("GetIrisReverseFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newHALIDModelSerializerField(t *testing.T) {
	type args struct {
		getURL   func(*url.URL, url.Values, ...any) string
		queries  map[string][]SerializeModelUseCase[*struct{}, string]
		params   []SerializeModelUseCase[*struct{}, any]
		ctx      context.Context
		instance *struct{}
	}
	u, err := url.Parse("https://sampleserver.com")
	require.NoError(t, err)
	ctx := context.WithValue(context.Background(), _currentURLKey, u)

	getURLFunc := func(u *url.URL, v url.Values, a ...any) string {
		return u.String()
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Success",
			args: args{
				ctx:      ctx,
				instance: new(struct{}),
				getURL:   getURLFunc,
				queries:  nil,
				params:   nil,
			},
			want: u.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newHALIDModelSerializerField(tt.args.getURL, tt.args.queries, tt.args.params...)
			if gotResult := got.Serialize(tt.args.ctx, tt.args.instance); gotResult != tt.want {
				t.Errorf("newHALIDModelSerializerField() = %v, want %v", got, tt.want)
			}
		})
	}
}

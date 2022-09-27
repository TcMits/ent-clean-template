package v1

import (
	"testing"

	_ "github.com/TcMits/ent-clean-template/docs"
	"github.com/TcMits/ent-clean-template/internal/testutils"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

func TestRegisterDocsController(t *testing.T) {
	type args struct {
		handler iris.Party
		l       logger.Interface
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{
				handler: iris.New(),
				l:       testutils.NullLogger{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterDocsController(tt.args.handler, tt.args.l)
		})
	}
}

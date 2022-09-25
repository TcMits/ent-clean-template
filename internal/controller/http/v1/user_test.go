package v1

import (
	"testing"

	"github.com/TcMits/ent-clean-template/internal/testutils"
	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/golang/mock/gomock"
	"github.com/kataras/iris/v12"
)

func TestRegisterPublicMeController(t *testing.T) {
	type args struct {
		handler             iris.Party
		getUseCase          usecase.GetModelUseCase[*model.User, *struct{}]
		getAndUpdateUseCase usecase.GetAndUpdateModelUseCase[*model.User, *struct{}, *useCaseModel.PublicMeUseCaseUpdateInput]
		serializeUseCase    usecase.SerializeModelUseCase[*model.User, map[string]any]
		l                   logger.Interface
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	getUseCase := usecase.NewMockGetModelUseCase[*model.User, *struct{}](ctrl)
	getAndUpdateUseCase := usecase.NewMockGetAndUpdateModelUseCase[*model.User, *struct{}, *useCaseModel.PublicMeUseCaseUpdateInput](
		ctrl,
	)

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{
				handler:             iris.New(),
				getUseCase:          getUseCase,
				getAndUpdateUseCase: getAndUpdateUseCase,
				l:                   testutils.NullLogger{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterPublicMeController(
				tt.args.handler,
				tt.args.getUseCase,
				tt.args.getAndUpdateUseCase,
				tt.args.serializeUseCase,
				tt.args.l,
			)
		})
	}
}

/* Specify the name of the generated file's package. */ package copygen

import (
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	"github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
)

type Copygen interface {
	LoginInputToUserWhereInput(*usecase.LoginInput) *model.UserWhereInput
	PublicMeUseCaseUpdateInputToUserUpdateInput(
		*useCaseModel.PublicMeUseCaseUpdateInput,
	) *model.UserUpdateInput
}

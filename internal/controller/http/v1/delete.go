package v1

import (
	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

func getDeleteHandler[
	ModelType any,
	PWhereUserInputType interface{ *WhereUserInputType },
	SerializedType,
	WhereUserInputType any,
](
	createUseCase usecase.GetAndDeleteModelUseCase[ModelType, PWhereUserInputType],
	serializeUseCase usecase.SerializeModelUseCase[ModelType, SerializedType],
	l logger.Interface,

	wrapReadParamsError func(model.TranslateFunc, error) error,
	wrapReadQueryError func(model.TranslateFunc, error) error,
) iris.Handler {
	return func(ctx iris.Context) {
		whereInput := PWhereUserInputType(new(WhereUserInputType))
		if err := ctx.ReadParams(whereInput); err != nil {
			handleBindingError(ctx, err, l, whereInput, wrapReadParamsError)
			return
		}
		if err := ctx.ReadQuery(whereInput); err != nil {
			handleBindingError(ctx, err, l, whereInput, wrapReadQueryError)
			return
		}
		context := ctx.Request().Context()
		err := createUseCase.GetAndDelete(context, whereInput)
		if err != nil {
			handleError(ctx, err, l)
			return
		}
		ctx.StatusCode(iris.StatusNoContent)
		ctx.JSON(iris.Map{})
	}
}

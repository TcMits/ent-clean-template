package v1

import (
	"github.com/kataras/iris/v12"

	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
)

func getUpdateHandler[
	ModelType any,
	PWhereUserInputType interface{ *WhereUserInputType },
	PUpdateInputType interface{ *UpdateInputType },
	SerializedType,
	WhereUserInputType,
	UpdateInputType any,
](
	getAndUpdateUseCase usecase.GetAndUpdateModelUseCase[ModelType, PWhereUserInputType, PUpdateInputType],
	serializeUseCase usecase.SerializeModelUseCase[ModelType, SerializedType],
	l logger.Interface,
	wrapReadParamsError func(error) error,
	wrapReadQueryError func(error) error,
	wrapReadBodyError func(error) error,
) iris.Handler {
	return func(ctx iris.Context) {
		whereInput := PWhereUserInputType(new(WhereUserInputType))
		updateInput := PUpdateInputType(new(UpdateInputType))
		if err := ctx.ReadParams(whereInput); err != nil {
			handleBindingError(ctx, err, l, whereInput, wrapReadParamsError)
			return
		}
		if err := ctx.ReadQuery(whereInput); err != nil {
			handleBindingError(ctx, err, l, whereInput, wrapReadQueryError)
			return
		}
		if err := ctx.ReadBody(updateInput); err != nil {
			handleBindingError(ctx, err, l, updateInput, wrapReadBodyError)
			return
		}
		context := ctx.Request().Context()
		instance, err := getAndUpdateUseCase.GetAndUpdate(context, whereInput, updateInput)
		if err != nil {
			handleError(ctx, err, l)
			return
		}
		payload := serializeUseCase.Serialize(context, instance)
		ctx.JSON(payload)
	}
}

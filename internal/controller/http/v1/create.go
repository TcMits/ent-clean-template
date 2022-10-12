package v1

import (
	"github.com/kataras/iris/v12"

	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
)

func getCreateHandler[
	ModelType any,
	PCreateInputType interface{ *CreateInputType },
	SerializedType,
	CreateInputType any,
](
	createUseCase usecase.CreateModelUseCase[ModelType, PCreateInputType],
	serializeUseCase usecase.SerializeModelUseCase[ModelType, SerializedType],
	l logger.Interface,

	wrapReadBodyError func(error) error,
) iris.Handler {
	return func(ctx iris.Context) {
		createInput := PCreateInputType(new(CreateInputType))
		if err := ctx.ReadBody(createInput); err != nil {
			handleBindingError(ctx, err, l, createInput, wrapReadBodyError)
			return
		}
		context := ctx.Request().Context()
		instance, err := createUseCase.Create(context, createInput)
		if err != nil {
			handleError(ctx, err, l)
			return
		}
		payload := serializeUseCase.Serialize(context, instance)
		ctx.StatusCode(iris.StatusCreated)
		ctx.JSON(payload)
	}
}

package v1

import (
	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
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

	wrapReadBodyError func(model.TranslateFunc, error) error,
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

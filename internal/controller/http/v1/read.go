package v1

import (
	"strconv"

	"github.com/kataras/iris/v12"

	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/TcMits/ent-clean-template/pkg/tool/url"
)

const (
	_defaultLimit = 10
)

var _defaultLimitOffsetInput = &limitOffsetQueryInput{
	Limit:     _defaultLimit,
	Offset:    0,
	WithCount: false,
}

type limitOffsetQueryInput struct {
	Limit     int  `url:"limit"      validate:"min=0,max=2000"`
	Offset    int  `url:"offset"     validate:"min=0"`
	WithCount bool `url:"with_count"`
}

func paginate[ModelType, OrderInputType, WhereInputType any](
	ctx iris.Context,
	listUseCase usecase.ListModelUseCase[ModelType, OrderInputType, WhereInputType],
	countUseCase usecase.CountModelUseCase[WhereInputType],
	orderInput OrderInputType,
	whereInput WhereInputType,
	limitOffsetInput *limitOffsetQueryInput,
) ([]ModelType, map[string]any, error) {
	req := ctx.Request()
	context := req.Context()
	paginateMeta := map[string]any{"next": nil, "previous": nil}
	limit := limitOffsetInput.Limit + 1
	instances, err := listUseCase.List(
		context,
		&limit,
		&limitOffsetInput.Offset,
		orderInput,
		whereInput,
	)
	if err != nil {
		return nil, nil, err
	}
	clonedURL := url.CloneURL(req.URL)

	// has next
	if len(instances) > limitOffsetInput.Limit {
		values := clonedURL.Query()
		values.Set("offset", strconv.Itoa(
			limitOffsetInput.Offset+limitOffsetInput.Limit,
		))
		clonedURL.RawQuery = values.Encode()
		paginateMeta["next"] = clonedURL.String()
	}

	// has previous
	if limitOffsetInput.Offset != 0 {
		previousOffset := limitOffsetInput.Offset - limitOffsetInput.Limit
		if previousOffset < 0 {
			previousOffset = 0
		}
		values := clonedURL.Query()
		values.Set("offset", strconv.Itoa(previousOffset))
		clonedURL.RawQuery = values.Encode()
		paginateMeta["previous"] = clonedURL.String()
	}

	// count
	if limitOffsetInput.WithCount && countUseCase != nil {
		count, err := countUseCase.Count(context, whereInput)
		if err != nil {
			return nil, nil, err
		}
		paginateMeta["count"] = count
	}
	return instances, paginateMeta, nil
}

func getDetailHandler[
	ModelType any,
	PWhereInputType interface{ *WhereInputType },
	SerializedType,
	WhereInputType any,
](
	getUseCase usecase.GetModelUseCase[ModelType, PWhereInputType],
	serializeUseCase usecase.SerializeModelUseCase[ModelType, SerializedType],
	l logger.Interface,
	wrapReadParamsError func(error) error,
	wrapReadQueryError func(error) error,
) iris.Handler {
	return func(ctx iris.Context) {
		whereInput := PWhereInputType(new(WhereInputType))
		if err := ctx.ReadParams(whereInput); err != nil {
			handleBindingError(ctx, err, l, whereInput, wrapReadParamsError)
			return
		}
		if err := ctx.ReadQuery(whereInput); err != nil {
			handleBindingError(ctx, err, l, whereInput, wrapReadQueryError)
			return
		}
		context := ctx.Request().Context()
		instance, err := getUseCase.Get(context, whereInput)
		if err != nil {
			handleError(ctx, err, l)
			return
		}
		payload := serializeUseCase.Serialize(context, instance)
		ctx.JSON(payload)
	}
}

func getListHandler[
	ModelType any,
	POrderInputType interface{ *OrderInputType },
	PWhereInputType interface{ *WhereInputType },
	SerializedType,
	OrderInputType,
	WhereInputType any,
](
	listUseCase usecase.ListModelUseCase[ModelType, POrderInputType, PWhereInputType],
	countUseCase usecase.CountModelUseCase[PWhereInputType],
	serializeUseCase usecase.SerializeModelUseCase[ModelType, SerializedType],
	l logger.Interface,

	wrapReadParamsError func(error) error,
	wrapReadQueryError func(error) error,
) iris.Handler {
	return func(ctx iris.Context) {
		whereInput := PWhereInputType(new(WhereInputType))
		orderInput := POrderInputType(new(OrderInputType))
		limitOffsetInput := *_defaultLimitOffsetInput
		if err := ctx.ReadParams(whereInput); err != nil {
			handleBindingError(ctx, err, l, whereInput, wrapReadParamsError)
			return
		}
		if err := ctx.ReadQuery(whereInput); err != nil {
			handleBindingError(ctx, err, l, whereInput, wrapReadQueryError)
			return
		}
		if err := ctx.ReadQuery(orderInput); err != nil {
			handleBindingError(ctx, err, l, orderInput, wrapReadQueryError)
			return
		}
		if err := ctx.ReadQuery(&limitOffsetInput); err != nil {
			handleBindingError(ctx, err, l, orderInput, wrapReadQueryError)
			return
		}

		instances, paginateMeta, err := paginate(
			ctx, listUseCase, countUseCase, orderInput, whereInput, &limitOffsetInput,
		)
		if err != nil {
			handleError(ctx, err, l)
			return
		}
		context := ctx.Request().Context()
		size := len(instances)
		payload := make([]SerializedType, 0, size)
		for i := 0; i < size; i++ {
			payload = append(payload, serializeUseCase.Serialize(context, instances[0]))
			instances = instances[1:]
		}
		ctx.JSON(iris.Map{"meta": paginateMeta, "results": payload})
	}
}

package usecase

import (
	"context"
	"net/url"

	"github.com/TcMits/ent-clean-template/pkg/tool/generic"
)

const (
	_currentURLKey = "URL"
)

type (
	modelSerializerField[ModelType, SerializedType any] func(context.Context, ModelType) SerializedType
)

func (s modelSerializerField[ModelType, SerializedType]) Serialize(
	ctx context.Context,
	instance ModelType,
) SerializedType {
	if s == nil {
		return generic.Zero[SerializedType]()
	}

	return s(ctx, instance)
}

func newHALIDModelSerializerField[ModelType any](
	getURL func(*url.URL, url.Values, ...any) string,
	queries map[string][]SerializeModelUseCase[ModelType, string],
	params ...SerializeModelUseCase[ModelType, any],
) SerializeModelUseCase[ModelType, string] {
	var result modelSerializerField[ModelType, string] = func(ctx context.Context, mt ModelType) string {
		paramValues := make([]any, 0, len(params))
		queryValues := url.Values{}
		currentURL, ok := ctx.Value(_currentURLKey).(*url.URL)

		if !ok {
			currentURL = nil
		}

		for _, paramFunc := range params {
			paramValues = append(paramValues, paramFunc.Serialize(ctx, mt))
		}

		for k, queryFuncs := range queries {
			queryValues[k] = make([]string, 0, len(queryFuncs))
			for _, queryFunc := range queryFuncs {
				queryValues[k] = append(queryValues[k], queryFunc.Serialize(ctx, mt))
			}
		}

		return getURL(currentURL, queryValues, paramValues...)
	}

	return result
}

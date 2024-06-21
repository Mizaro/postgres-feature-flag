package postgrestfeatureflags

import (
	"context"
	"fmt"
	"github.com/open-feature/go-sdk/openfeature"
	"github.com/supabase-community/postgrest-go"
	"strconv"
)

// PostgrestProvider implements the FeatureProvider interface and provides functions for evaluating flags
type PostgrestProvider struct {
	client *postgrest.Client
}

func NewPostgrestProvider(rawURL string) *PostgrestProvider {
	client := postgrest.NewClient(rawURL, "", nil)
	if client.ClientError != nil {
		panic(client.ClientError)
	}
	return &PostgrestProvider{
		client: client,
	}
}

func (p PostgrestProvider) Metadata() openfeature.Metadata {
	return openfeature.Metadata{
		Name: "PostgrestProvider",
	}
}

type result struct {
	Value string `json:"value"`
}

func evaluation[V any](p PostgrestProvider, flag string, conv func(string) (V, error)) (*V, *openfeature.ResolutionError) {
	var res []result

	_, err := p.client.From("feature").
		Select("value", "", false).
		Eq("name", flag).
		Limit(1, "").
		ExecuteTo(&res)
	if err != nil {
		x := openfeature.NewGeneralResolutionError(err.Error())
		return nil, &x
	}
	if len(res) < 1 {
		x := openfeature.NewFlagNotFoundResolutionError(fmt.Sprintf("flag `%s` could not beyond found", flag))
		return nil, &x
	}

	casted, err := conv(res[0].Value)
	if err != nil {
		x := openfeature.NewTypeMismatchResolutionError(err.Error())
		return nil, &x
	}

	return &casted, nil
}

func (p PostgrestProvider) BooleanEvaluation(ctx context.Context, flag string, defaultValue bool, evalCtx openfeature.FlattenedContext) openfeature.BoolResolutionDetail {
	value, resErr := evaluation[bool](p, flag, strconv.ParseBool)
	if resErr != nil {
		return openfeature.BoolResolutionDetail{
			Value: defaultValue,
			ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
				ResolutionError: *resErr,
				FlagMetadata:    openfeature.FlagMetadata(evalCtx),
			},
		}
	}

	return openfeature.BoolResolutionDetail{
		Value: *value,
	}
}

func (p PostgrestProvider) StringEvaluation(ctx context.Context, flag string, defaultValue string, evalCtx openfeature.FlattenedContext) openfeature.StringResolutionDetail {
	//TODO implement me
	panic("implement me")
}

func (p PostgrestProvider) FloatEvaluation(ctx context.Context, flag string, defaultValue float64, evalCtx openfeature.FlattenedContext) openfeature.FloatResolutionDetail {
	//TODO implement me
	panic("implement me")
}

func (p PostgrestProvider) IntEvaluation(ctx context.Context, flag string, defaultValue int64, evalCtx openfeature.FlattenedContext) openfeature.IntResolutionDetail {
	//TODO implement me
	panic("implement me")
}

func (p PostgrestProvider) ObjectEvaluation(ctx context.Context, flag string, defaultValue interface{}, evalCtx openfeature.FlattenedContext) openfeature.InterfaceResolutionDetail {
	//TODO implement me
	panic("implement me")
}

func (p PostgrestProvider) Hooks() []openfeature.Hook {
	return []openfeature.Hook{}
}

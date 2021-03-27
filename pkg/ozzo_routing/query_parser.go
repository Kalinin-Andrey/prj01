package ozzo_routing

import (
	routing "github.com/go-ozzo/ozzo-routing/v2"

	"carizza/pkg/selection_condition"
)

func ParseQueryParams(ctx *routing.Context, st interface{}) (*selection_condition.SelectionCondition, error) {
	return selection_condition.ParseQueryParams(ctx.Request.URL.Query(), st)
}

func ParseQueryParamsIntoStruct(ctx *routing.Context, out interface{}) error {
	return selection_condition.ParseQueryParamsIntoStruct(ctx.Request.URL.Query(), out)
}

func ParseUintParam(ctx *routing.Context, paramName string) (uint, error) {
	return selection_condition.ParseUintParam(ctx.Param(paramName))
}

func ParseUintQueryParam(ctx *routing.Context, paramName string) (uint, error) {
	return selection_condition.ParseUintParam(ctx.Query(paramName))
}

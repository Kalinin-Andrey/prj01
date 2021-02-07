package controller

import (
	"carizza/internal/domain"
	"carizza/internal/pkg/apperror"
	"carizza/internal/pkg/log"
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"strconv"
)

type IService interface{}

type Controller struct {
	Logger log.ILogger
}

var matchedParams = []string{}

func (c Controller) parseUintParam(ctx *routing.Context, paramName string) (uint, error) {
	str := ctx.Param(paramName)
	if str == "" {
		return 0, apperror.ErrNotFound
	}

	paramVal, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		c.Logger.With(ctx.Request.Context()).Info(err)
		return 0, err
	}
	return uint(paramVal), nil
}

func (c Controller) parseUintQueryParam(ctx *routing.Context, paramName string) (uint, error) {
	str := ctx.Query(paramName)
	if str == "" {
		return 0, apperror.ErrNotFound
	}

	paramVal, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		c.Logger.With(ctx.Request.Context()).Info(err)
		return 0, err
	}
	return uint(paramVal), nil
}

func (c Controller) ExtractQueryFromRoutingContext(ctx *routing.Context) domain.DBQueryConditions {
	query := map[string]interface{}{}

	for _, paramName := range matchedParams {

		if paramVal := ctx.Param(paramName); paramVal != "" {
			query[paramName] = paramVal
		}
	}

	return domain.DBQueryConditions{
		Where: query,
	}
}

package controller

import (
	"carizza/internal/pkg/apperror"
	"carizza/internal/pkg/log"
	"strconv"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

type IService interface{}

type Controller struct {
	Logger log.ILogger
}

var matchedParams = []string{}

func (c Controller) parseQueryParams(ctx *routing.Context, out interface{}) error {
	var v map[string]string

	for key, vals := range ctx.Request.URL.Query() {
		if len(vals) > 0 {
			v[key] = vals[0]
		}
	}

}

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

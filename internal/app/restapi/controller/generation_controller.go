package controller

import (
	"errors"

	"github.com/go-ozzo/ozzo-routing/v2"

	"carizza/internal/pkg/apperror"
	"carizza/internal/pkg/errorshandler"
	"carizza/internal/pkg/log"

	"carizza/internal/domain"
	"carizza/internal/domain/generation"
)

type generationController struct {
	Controller
	Service generation.IService
}

// RegisterHandlers sets up the routing of the HTTP handlers.
//	GET /api/generations/ - список всех моделей
//	GET /api/generation/{ID} - детали модели
func RegisterGenerationHandlers(r *routing.RouteGroup, service generation.IService, logger log.ILogger, authHandler routing.Handler) {
	c := generationController{
		Controller: Controller{
			Logger: logger,
		},
		Service: service,
	}

	r.Get("/generations", c.list)
	r.Get(`/generation/<id>`, c.get)
	r.Get(`/model/<id>/generations`, c.list)
}

// get method is for getting a one entity by ID
func (c generationController) get(ctx *routing.Context) error {
	id, err := c.parseUintParam(ctx, "id")
	if err != nil {
		errorshandler.BadRequest("ID is required to be uint")
	}

	entity, err := c.Service.Get(ctx.Request.Context(), id)
	if err != nil {
		if err == apperror.ErrNotFound {
			c.Logger.With(ctx.Request.Context()).Info(err)
			return errorshandler.NotFound("not found")
		}
		c.Logger.With(ctx.Request.Context()).Error(err)
		return errorshandler.InternalServerError("")
	}

	ctx.Response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return ctx.Write(entity)
}

// list method is for a getting a list of all entities
func (c generationController) list(ctx *routing.Context) error {
	cond := domain.DBQueryConditions{
		SortOrder: map[string]string{
			"name": "asc",
		},
	}

	modelId, err := c.parseUintParam(ctx, "modelId")
	if errors.Is(err, apperror.ErrNotFound) {
		modelId, err = c.parseUintQueryParam(ctx, "modelId")
	}
	if err == nil && modelId > 0 {
		cond.Where = map[string]interface{}{
			"id_car_model": modelId,
		}
	}

	items, err := c.Service.Query(ctx.Request.Context(), cond)
	if err != nil {
		if err == apperror.ErrNotFound {
			c.Logger.With(ctx.Request.Context()).Info(err)
			return errorshandler.NotFound("")
		}
		c.Logger.With(ctx.Request.Context()).Error(err)
		return errorshandler.InternalServerError("")
	}
	ctx.Response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return ctx.Write(items)
}

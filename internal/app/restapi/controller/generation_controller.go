package controller

import (
	"errors"

	"carizza/internal/pkg/apperror"

	"github.com/minipkg/go-app-common/log"
	ozzo_handler "github.com/minipkg/go-app-common/ozzo_handler"
	"github.com/minipkg/go-app-common/ozzo_handler/errorshandler"

	"carizza/internal/domain"
	"carizza/internal/domain/generation"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

type generationController struct {
	Logger  log.ILogger
	Service generation.IService
}

// RegisterHandlers sets up the routing of the HTTP handlers.
//	GET /api/generations/ - список всех моделей
//	GET /api/generation/{ID} - детали модели
func RegisterGenerationHandlers(r *routing.RouteGroup, service generation.IService, logger log.ILogger, authHandler routing.Handler) {
	c := generationController{
		Logger:  logger,
		Service: service,
	}

	r.Get("/generations", c.list)
	r.Get(`/generation/<id>`, c.get)
	r.Get(`/model/<id>/generations`, c.list)
}

// get method is for getting a one entity by ID
func (c generationController) get(ctx *routing.Context) error {
	id, err := ozzo_handler.ParseUintParam(ctx, "id")
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

	return ctx.Write(entity)
}

// list method is for a getting a list of all entities
func (c generationController) list(ctx *routing.Context) error {
	cond := domain.DBQueryConditions{
		SortOrder: map[string]string{
			"name": "asc",
		},
	}

	modelId, err := ozzo_handler.ParseUintParam(ctx, "modelId")
	if errors.Is(err, apperror.ErrNotFound) {
		modelId, err = ozzo_handler.ParseUintQueryParam(ctx, "modelId")
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
	return ctx.Write(items)
}

package controller

import (
	"errors"

	"carizza/internal/pkg/apperror"
	"carizza/internal/pkg/errorshandler"
	"carizza/internal/pkg/log"
	ozzo_handler "carizza/pkg/ozzo_handler"

	"carizza/internal/domain"
	"carizza/internal/domain/serie"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

type serieController struct {
	Logger  log.ILogger
	Service serie.IService
}

// RegisterHandlers sets up the routing of the HTTP handlers.
//	GET /api/series/ - список всех моделей
//	GET /api/serie/{ID} - детали модели
func RegisterSerieHandlers(r *routing.RouteGroup, service serie.IService, logger log.ILogger, authHandler routing.Handler) {
	c := serieController{
		Logger:  logger,
		Service: service,
	}

	r.Get("/series", c.list)
	r.Get(`/serie/<id>`, c.get)
	r.Get(`/generation/<id>/series`, c.list)
}

// get method is for getting a one entity by ID
func (c serieController) get(ctx *routing.Context) error {
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
func (c serieController) list(ctx *routing.Context) error {
	cond := domain.DBQueryConditions{
		SortOrder: map[string]string{
			"name": "asc",
		},
	}

	generationId, err := ozzo_handler.ParseUintParam(ctx, "generationId")
	if errors.Is(err, apperror.ErrNotFound) {
		generationId, err = ozzo_handler.ParseUintQueryParam(ctx, "generationId")
	}
	if err == nil && generationId > 0 {
		cond.Where = map[string]interface{}{
			"id_car_generation": generationId,
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

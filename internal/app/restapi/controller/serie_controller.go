package controller

import (
	"github.com/go-ozzo/ozzo-routing/v2"

	"carizza/internal/pkg/apperror"
	"carizza/internal/pkg/errorshandler"
	"carizza/internal/pkg/log"

	"carizza/internal/domain"
	"carizza/internal/domain/serie"
)

type serieController struct {
	Controller
	Service serie.IService
}

// RegisterHandlers sets up the routing of the HTTP handlers.
//	GET /api/series/ - список всех моделей
//	GET /api/serie/{ID} - детали модели
func RegisterSerieHandlers(r *routing.RouteGroup, service serie.IService, logger log.ILogger, authHandler routing.Handler) {
	c := serieController{
		Controller: Controller{
			Logger: logger,
		},
		Service: service,
	}

	r.Get("/series", c.list)
	r.Get(`/serie/<id>`, c.get)
}

// get method is for getting a one entity by ID
func (c serieController) get(ctx *routing.Context) error {
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
func (c serieController) list(ctx *routing.Context) error {
	cond := domain.DBQueryConditions{
		SortOrder: map[string]string{
			"name": "asc",
		},
	}

	generationId, err := c.parseUintQueryParam(ctx, "generationId")
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
	ctx.Response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return ctx.Write(items)
}

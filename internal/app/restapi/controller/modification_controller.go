package controller

import (
	"github.com/go-ozzo/ozzo-routing/v2"

	"carizza/internal/pkg/apperror"
	"carizza/internal/pkg/errorshandler"
	"carizza/internal/pkg/log"

	"carizza/internal/domain"
	"carizza/internal/domain/modification"
)

type modificationController struct {
	Controller
	Service modification.IService
}

// RegisterHandlers sets up the routing of the HTTP handlers.
//	GET /api/modifications/ - список всех моделей
//	GET /api/modification/{ID} - детали модели
func RegisterModificationHandlers(r *routing.RouteGroup, service modification.IService, logger log.ILogger, authHandler routing.Handler) {
	c := modificationController{
		Controller: Controller{
			Logger: logger,
		},
		Service: service,
	}

	r.Get("/modifications", c.list)
	r.Get(`/modification/<id>`, c.get)
}

// get method is for getting a one entity by ID
func (c modificationController) get(ctx *routing.Context) error {
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
func (c modificationController) list(ctx *routing.Context) error {
	cond := domain.DBQueryConditions{
		SortOrder: map[string]string{
			"name": "asc",
		},
	}

	serieId, err := c.parseUintQueryParam(ctx, "serieId")
	if err == nil && serieId > 0 {
		cond.Where = map[string]interface{}{
			"id_car_serie": serieId,
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
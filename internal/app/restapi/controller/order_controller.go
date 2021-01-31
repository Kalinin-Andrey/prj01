package controller

import (
	"github.com/go-ozzo/ozzo-routing/v2"

	"carizza/internal/pkg/apperror"
	"carizza/internal/pkg/errorshandler"
	"carizza/internal/pkg/log"

	"carizza/internal/domain"
	"carizza/internal/domain/order"
)

type orderController struct {
	Controller
	Service order.IService
}

// RegisterHandlers sets up the routing of the HTTP handlers.
//	GET /api/orders/ - список всех моделей
//	GET /api/order/{ID} - детали модели
func RegisterOrderHandlers(r *routing.RouteGroup, service order.IService, logger log.ILogger, authHandler routing.Handler) {
	c := orderController{
		Controller: Controller{
			Logger: logger,
		},
		Service: service,
	}

	r.Get("/orders", c.list)
	r.Get(`/order/<id>`, c.get)
}

// get method is for getting a one entity by ID
func (c orderController) get(ctx *routing.Context) error {
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
func (c orderController) list(ctx *routing.Context) error {
	cond := domain.DBQueryConditions{
		SortOrder: map[string]string{
			"name": "asc",
		},
	}

	markId, err := c.parseUintQueryParam(ctx, "markId")
	if err == nil && markId > 0 {
		cond.Where = map[string]interface{}{
			"id_car_mark": markId,
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

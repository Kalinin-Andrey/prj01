package controller

import (
	"errors"

	"github.com/go-ozzo/ozzo-routing/v2"

	"carizza/internal/pkg/apperror"
	"carizza/internal/pkg/errorshandler"
	"carizza/internal/pkg/log"

	"carizza/internal/domain"
	"carizza/internal/domain/supply"
)

type supplyController struct {
	Controller
	Service supply.IService
}

// RegisterHandlers sets up the routing of the HTTP handlers.
//	GET /api/supplys/ - список всех работ
//	GET /api/supply/<ID> - детали модели
//	GET /work/<workId>/supplies - список работ для услуги
func RegisterSupplyHandlers(r *routing.RouteGroup, service supply.IService, logger log.ILogger, authHandler routing.Handler) {
	c := supplyController{
		Controller: Controller{
			Logger: logger,
		},
		Service: service,
	}

	r.Get("/supplys", c.list)
	r.Get(`/supply/<id>`, c.get)
	r.Get(`/work/<workId>/supplies`, c.list)
}

// get method is for getting a one entity by ID
func (c supplyController) get(ctx *routing.Context) error {
	id, err := c.parseUintParam(ctx, "id")
	if err != nil {
		return errorshandler.BadRequest("ID is required to be uint")
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
func (c supplyController) list(ctx *routing.Context) error {
	cond := domain.DBQueryConditions{
		SortOrder: map[string]string{
			"name": "asc",
		},
	}

	workId, err := c.parseUintParam(ctx, "workId")
	if errors.Is(err, apperror.ErrNotFound) {
		workId, err = c.parseUintQueryParam(ctx, "workId")
	}
	if err == nil && workId > 0 {
		cond.Where = map[string]interface{}{
			"workId": workId,
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
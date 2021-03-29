package controller

import (
	"carizza/pkg/selection_condition"
	"errors"

	"carizza/internal/pkg/apperror"

	"github.com/minipkg/go-app-common/log"
	ozzo_handler "github.com/minipkg/go-app-common/ozzo_handler"
	"github.com/minipkg/go-app-common/ozzo_handler/errorshandler"

	"carizza/internal/domain/work"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

type workController struct {
	Logger  log.ILogger
	Service work.IService
}

// RegisterHandlers sets up the routing of the HTTP handlers.
//	GET /api/works/ - список всех работ
//	GET /api/work/<id> - детали модели
//	GET /maintenance/<maintenanceId>/works - список работ для услуги
func RegisterWorkHandlers(r *routing.RouteGroup, service work.IService, logger log.ILogger, authHandler routing.Handler) {
	c := workController{
		Logger:  logger,
		Service: service,
	}

	r.Get("/works", c.list)
	r.Get(`/work/<id>`, c.get)
	r.Get(`/maintenance/<maintenanceId>/works`, c.list)
}

// get method is for getting a one entity by ID
func (c workController) get(ctx *routing.Context) error {
	id, err := ozzo_handler.ParseUintParam(ctx, "id")
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

	return ctx.Write(entity)
}

// list method is for a getting a list of all entities
func (c workController) list(ctx *routing.Context) error {
	cond := &selection_condition.SelectionCondition{
		SortOrder: []map[string]string{{
			"name": "asc",
		}},
	}

	maintenanceId, err := ozzo_handler.ParseUintParam(ctx, "maintenanceId")
	if errors.Is(err, apperror.ErrNotFound) {
		maintenanceId, err = ozzo_handler.ParseUintQueryParam(ctx, "maintenanceId")
	}
	if err == nil && maintenanceId > 0 {
		cond.Where = map[string]interface{}{
			"maintenanceId": maintenanceId,
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

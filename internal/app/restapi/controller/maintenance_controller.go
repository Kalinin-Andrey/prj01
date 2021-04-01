package controller

import (
	"github.com/minipkg/selection_condition"
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing/v2"

	"carizza/internal/pkg/apperror"

	"github.com/minipkg/log"
	"github.com/minipkg/ozzo_routing"
	"github.com/minipkg/ozzo_routing/errorshandler"

	"carizza/internal/domain/maintenance"
)

type maintenanceController struct {
	Logger  log.ILogger
	Service maintenance.IService
}

// RegisterHandlers sets up the routing of the HTTP handlers.
//	GET /api/maintenances/ - список всех услуг
//	GET /api/maintenance/<ID> - детали услуги
//	POST /api/maintenance - создание услуги
//	PUT /api/maintenance/<ID> - обновление услуги
//	DELETE /api/maintenance/<ID> - удаление услуги
func RegisterMaintenanceHandlers(r *routing.RouteGroup, service maintenance.IService, logger log.ILogger, authHandler routing.Handler) {
	c := maintenanceController{
		Logger:  logger,
		Service: service,
	}

	r.Get("/maintenances", c.list)
	r.Get(`/maintenance/<id>`, c.get)

	r.Use(authHandler)

	r.Post("/maintenance", c.create)
	r.Put("/maintenance/<id>", c.update)
	r.Delete(`/maintenance/<id>`, c.delete)

}

// get method is for getting a one entity by ID
func (c maintenanceController) get(ctx *routing.Context) error {
	id, err := ozzo_routing.ParseUintParam(ctx, "id")
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
func (c maintenanceController) list(ctx *routing.Context) error {
	cond := &selection_condition.SelectionCondition{
		SortOrder: []map[string]string{{
			"name": "asc",
		}},
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

func (c maintenanceController) create(ctx *routing.Context) error {
	entity := c.Service.NewEntity()
	if err := ctx.Read(entity); err != nil {
		c.Logger.With(ctx.Request.Context()).Info(err)
		return errorshandler.BadRequest(err.Error())
	}

	if err := entity.Validate(); err != nil {
		return errorshandler.BadRequest(err.Error())
	}

	if err := c.Service.Create(ctx.Request.Context(), entity); err != nil {
		c.Logger.With(ctx.Request.Context()).Info(err)
		return errorshandler.BadRequest(err.Error())
	}

	return ctx.WriteWithStatus(entity, http.StatusCreated)
}

func (c maintenanceController) update(ctx *routing.Context) error {
	id, err := ozzo_routing.ParseUintParam(ctx, "id")
	if err != nil {
		errorshandler.BadRequest("ID is required to be uint.")
	}

	entity := c.Service.NewEntity()
	if err := ctx.Read(entity); err != nil {
		c.Logger.With(ctx.Request.Context()).Info(err)
		return errorshandler.BadRequest(err.Error())
	}

	if id != entity.ID {
		errorshandler.BadRequest("ID in URI and body must be equal.")
	}

	if err := entity.Validate(); err != nil {
		return errorshandler.BadRequest(err.Error())
	}

	if err := c.Service.Update(ctx.Request.Context(), entity); err != nil {
		if err == apperror.ErrNotFound {
			c.Logger.With(ctx.Request.Context()).Info(err)
			return errorshandler.NotFound("")
		}
		c.Logger.With(ctx.Request.Context()).Error(err)
		return errorshandler.InternalServerError("")
	}

	return ctx.WriteWithStatus(errorshandler.SuccessMessage(), http.StatusOK)
}

func (c maintenanceController) delete(ctx *routing.Context) error {
	id, err := ozzo_routing.ParseUintParam(ctx, "id")
	if err != nil {
		errorshandler.BadRequest("ID is required to be uint.")
	}

	if err := c.Service.Delete(ctx.Request.Context(), id); err != nil {
		if err == apperror.ErrNotFound {
			c.Logger.With(ctx.Request.Context()).Info(err)
			return errorshandler.NotFound("")
		}
		c.Logger.With(ctx.Request.Context()).Error(err)
		return errorshandler.InternalServerError("")
	}

	ctx.Response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return ctx.WriteWithStatus(errorshandler.SuccessMessage(), http.StatusOK)
}

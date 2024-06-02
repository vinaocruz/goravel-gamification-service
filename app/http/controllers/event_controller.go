package controllers

import (
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"
)

type EventController struct {
	//Dependent services
}

func NewEventController() *EventController {
	return &EventController{
		//Inject services
	}
}

func (r *EventController) Index(ctx http.Context) http.Response {
	var events []models.Event
	facades.Orm().Query().Get(&events)

	return ctx.Response().Success().Json(events)
}

func (r *EventController) Show(ctx http.Context) http.Response {
	var event models.Event
	err := facades.Orm().Query().FindOrFail(&event, ctx.Request().RouteInt("id"))

	if err == orm.ErrRecordNotFound {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"message": err.Error(),
		})
	}

	return ctx.Response().Success().Json(event)
}

func (r *EventController) Store(ctx http.Context) http.Response {
	var event models.Event
	ctx.Request().Bind(&event)

	facades.Orm().Query().Create(&event)

	return ctx.Response().Json(http.StatusCreated, event)
}

func (r *EventController) Update(ctx http.Context) http.Response {
	var event models.Event
	err := facades.Orm().Query().FindOrFail(&event, ctx.Request().RouteInt("id"))

	if err == orm.ErrRecordNotFound {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"message": err.Error(),
		})
	}

	ctx.Request().Bind(&event)
	facades.Orm().Query().Save(&event)

	return ctx.Response().Success().Json(event)
}

func (r *EventController) Destroy(ctx http.Context) http.Response {
	var event models.Event
	res, _ := facades.Orm().Query().Delete(&event, ctx.Request().RouteInt("id"))

	if res.RowsAffected == 0 {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"message": "record not found",
		})
	}

	return ctx.Response().Json(http.StatusNoContent, nil)
}

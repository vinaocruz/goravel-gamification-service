package controllers

import (
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"
)

type PlayerController struct {
	//Dependent services
}

func NewPlayerController() *PlayerController {
	return &PlayerController{
		//Inject services
	}
}

func (r *PlayerController) Index(ctx http.Context) http.Response {
	var players []models.Player
	facades.Orm().Query().Get(&players)

	return ctx.Response().Success().Json(players)
}

func (r *PlayerController) Show(ctx http.Context) http.Response {
	id := ctx.Request().RouteInt("id")

	var player models.Player
	err := facades.Orm().Query().With("Events").FindOrFail(&player, id)

	if err == orm.ErrRecordNotFound {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"message": err.Error(),
		})
	}

	return ctx.Response().Success().Json(player)
}

func (r *PlayerController) Store(ctx http.Context) http.Response {
	var player models.Player
	ctx.Request().Bind(&player)

	err := facades.Orm().Query().Create(&player)

	if err != nil {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{
			"message": err.Error(),
		})
	}

	return ctx.Response().Success().Json(player)
}

func (r *PlayerController) Update(ctx http.Context) http.Response {
	id := ctx.Request().RouteInt("id")

	var player models.Player
	err := facades.Orm().Query().FindOrFail(&player, id)

	if err == orm.ErrRecordNotFound {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"message": err.Error(),
		})
	}

	ctx.Request().Bind(&player)
	facades.Orm().Query().Save(&player)

	return ctx.Response().Success().Json(player)
}

func (r *PlayerController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().RouteInt("id")

	var player models.Player
	res, _ := facades.Orm().Query().Delete(&player, id)

	if res.RowsAffected == 0 {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"message": "record not found",
		})
	}

	return ctx.Response().Json(http.StatusNoContent, nil)
}

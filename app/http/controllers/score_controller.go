package controllers

import (
	"encoding/json"
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"
)

type ScoreController struct {
	//Dependent services
}

type Score struct {
	PlayerId uint
	Player   string
	Score    int
}

func NewScoreController() *ScoreController {
	return &ScoreController{
		//Inject services
	}
}

func (r *ScoreController) Create(ctx http.Context) http.Response {
	var player models.Player
	err := facades.Orm().Query().FindOrFail(&player, ctx.Request().InputInt("playerId"))

	if err == orm.ErrRecordNotFound {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"message": "player not found",
		})
	}

	var event models.Event
	err = facades.Orm().Query().FindOrFail(&event, ctx.Request().InputInt("eventId"))

	if err == orm.ErrRecordNotFound {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"message": "event not found",
		})
	}

	var score models.PlayerEvents
	score.PlayerID = player.ID
	score.EventID = event.ID
	score.Points = event.Points

	facades.Orm().Query().Create(&score)

	facades.Cache().Forget("ranking")

	return ctx.Response().Json(http.StatusCreated, score)
}

func (r *ScoreController) Show(ctx http.Context) http.Response {
	var scores []Score

	if facades.Cache().Has("ranking") {
		rankingJson := facades.Cache().Get("ranking")
		err := json.Unmarshal([]byte(rankingJson.(string)), &scores)
		if err == nil {
			return ctx.Response().Success().Json(scores)
		}
	}

	var sql = "SELECT player_id, players.name as player, sum(points) as score FROM player_events JOIN players ON player_events.player_id = players.id GROUP BY player_id, players.name ORDER BY score DESC LIMIT 10"
	facades.Orm().Query().Raw(sql).Get(&scores)

	scoresJson, err := json.Marshal(scores)
	if err == nil {
		facades.Cache().Put("ranking", string(scoresJson), 0)
	}

	return ctx.Response().Success().Json(scores)
}

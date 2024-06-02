package routes

import (
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Api() {
	facades.Route().Resource("/players", controllers.NewPlayerController())
	facades.Route().Resource("/events", controllers.NewEventController())

	scoreController := controllers.NewScoreController()
	facades.Route().Post("/scores", scoreController.Create)
	facades.Route().Get("/scores", scoreController.Show)
}

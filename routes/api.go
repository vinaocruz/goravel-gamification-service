package routes

import (
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Api() {
	facades.Route().Resource("/players", controllers.NewPlayerController())
	facades.Route().Resource("/events", controllers.NewEventController())
	facades.Route().Post("/scores", controllers.NewScoreController().Create)
}

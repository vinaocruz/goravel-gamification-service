package routes

import (
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Api() {
	facades.Route().Resource("/players", controllers.NewPlayerController())
}

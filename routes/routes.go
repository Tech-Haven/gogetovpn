package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/tech-haven/gogetovpn/configs"
	"github.com/tech-haven/gogetovpn/controllers"
)

func Routes(e *echo.Echo, configuration *configs.Config) {
	e.POST("/ovpn", controllers.GetOvpn(configuration))
}

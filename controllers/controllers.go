package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tech-haven/gogetovpn/configs"
	"github.com/tech-haven/gogetovpn/responses"
	"github.com/tech-haven/gogetovpn/utils"
)

func GetOvpn(config *configs.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		client := c.FormValue("client")

		if client == "" {
			fmt.Printf("error: no client")
			return c.JSON(http.StatusBadRequest, responses.HTTPResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": "No client"}})
		}

		stdout, stderr, err := utils.ExecShell(fmt.Sprintf("sudo -S env MENU_OPTION='1' CLIENT='%s' PASS='1' %s > /dev/null && cat ~/%s.ovpn", client, configs.OpenVpnInstallScriptDir(), client))

		if err != nil {
			fmt.Printf("error: %v", stderr)
			return c.JSON(http.StatusInternalServerError, responses.HTTPResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": "Error retreiving VPN file"}})
		}

		return c.JSON(http.StatusOK, responses.HTTPResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": stdout}})
	}
}

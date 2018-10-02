package api

import (
	"net/http"

	"github.com/labstack/echo"
)

// Methods interface use to help passing method
type Methods interface {
	RunScript(scriptPath string) error
}

// API use to expose object
type API struct {
	ScriptRunner Methods
}

// Webhook method use for serve webhook
func (api *API) Webhook(c echo.Context) error {
	scriptPath := "./run-with-docker.sh"

	err := api.ScriptRunner.RunScript(scriptPath)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "ok",
	})
}

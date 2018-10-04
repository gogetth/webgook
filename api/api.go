package api

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/labstack/echo"
)

// ScriptRunner struct for  run script
type ScriptRunner struct {
}

// Methods interface use to help passing method
type Methods interface {
	RunScript(scriptPath string) error
	CheckScriptPath(scriptPath string) (string, error)
}

// API use to expose object
type API struct {
	ScriptRunner Methods
	ScriptPath   string
}

// Webhook method use for serve webhook
func (api *API) Webhook(c echo.Context) error {
	scriptPath, err := api.ScriptRunner.CheckScriptPath(api.ScriptPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	err = api.ScriptRunner.RunScript(scriptPath)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "ok",
	})
}

// RunScript for starting shellÎ
func (s ScriptRunner) RunScript(scriptPath string) error {
	cmd := exec.Command(scriptPath)

	err := cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error starting script", err)
		os.Exit(1)
	}
	return err
}

// CheckScriptPath use to check user specify path is legal
func (s ScriptRunner) CheckScriptPath(scriptPath string) (string, error) {
	if scriptPath == "" {
		return "", errors.New("script file must be specify")
	}

	_, err := os.Stat(scriptPath)
	if err != nil {
		return "", err
	}

	if os.IsNotExist(err) {
		return "", errors.New("script file not found")
	}

	return scriptPath, nil
}

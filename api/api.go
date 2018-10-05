package api

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/labstack/echo"
)

// Parameter is struct that we use to pass parameter from command line
type Parameter struct {
	ScriptPath string
	SecretKey  string
}

// ScriptRunner struct for  run script
type ScriptRunner struct {
}

// Methods interface use to help passing method
type Methods interface {
	RunScript(scriptPath string) error
	CheckScriptPath(scriptPath string) (string, error)
	VerifySecretKey(secretKeyFromCommandLine, secretKey string) (bool, error)
}

// API use to expose object
type API struct {
	ScriptRunner Methods
	Parameter    Parameter
}

// Webhook method use for serve webhook
func (api *API) Webhook(c echo.Context) error {
	secretKeyFromHook := c.QueryParam("key")

	virified, err := api.ScriptRunner.VerifySecretKey(api.Parameter.SecretKey, secretKeyFromHook)
	if err != nil {
		return responseInternalServerError(c, err)
	}

	if !virified {
		return responseInternalServerError(c, errors.New("failed to verify secret key"))
	}

	scriptPath, err := api.ScriptRunner.CheckScriptPath(api.Parameter.ScriptPath)
	if err != nil {
		return responseInternalServerError(c, err)
	}

	err = api.ScriptRunner.RunScript(scriptPath)
	if err != nil {
		return responseInternalServerError(c, err)
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

// VerifySecretKey use to verify hooking is legal
func (s ScriptRunner) VerifySecretKey(secretKeyFromCommandLine, secretKey string) (bool, error) {
	if secretKeyFromCommandLine == secretKey {
		return true, nil
	}
	return false, errors.New("secret key does not valid")
}

func responseInternalServerError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, map[string]string{
		"error": err.Error(),
	})
}

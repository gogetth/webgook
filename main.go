package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gogetth/webgook/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// ScriptRunner struct for  run script
type ScriptRunner struct {
}

func main() {
	e := echo.New()

	api := &api.API{
		ScriptRunner: &ScriptRunner{},
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/webhook", api.Webhook)
	e.Logger.Fatal(e.Start(":9000"))
}

// RunScript for starting shellÎ
func (s ScriptRunner) RunScript(scriptPath string) error {
	cmd := exec.Command(scriptPath)

	err := cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting script", err)
		os.Exit(1)
	}
	return err
}

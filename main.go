package main

import (
	"flag"
	"fmt"

	"github.com/gogetth/webgook/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var scriptPath *string

func main() {
	e := echo.New()

	// read flag from command line
	ip := flag.String("ip", "", "Specify ip address if you want")
	port := flag.String("port", "9000", "Specify port (default: 9000)")
	scriptPath = flag.String("script", "", "You must specify script")

	flag.Parse()
	serverString := fmt.Sprintf("%s:%s", *ip, *port)

	api := &api.API{
		ScriptRunner: &api.ScriptRunner{},
		ScriptPath:   *scriptPath,
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/webhook", api.Webhook)
	e.Logger.Fatal(e.Start(serverString))
}

package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/carstenmichel/goadder/backend"
	swaggerui "github.com/carstenmichel/goadder/backend/swagger"
	instana "github.com/instana/go-sensor"
	"github.com/instana/go-sensor/instrumentation/instaecho"
)

func setupEcho() *echo.Echo {
	log.Printf("starting with configuration\n")

	e := echo.New()
	e.HideBanner = true

	e.GET("/q/health", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]interface{}{"status": "UP"})
	})

	e.Use(middleware.Logger())
	swaggerui.Mount(e)
	backend.RegisterWarehouse(e)

	sensor := instana.NewSensor("echo-sensor")
	e.Use(instaecho.Middleware(sensor))
	return e

}

func main() {

	log.SetLevel(log.DebugLevel)
	//log.SetFormatter(&log.JSONFormatter{TimestampFormat: time.RFC3339})
	log.SetReportCaller(true)
	log.Println("Starting")

	e := setupEcho()
	log.Fatal(e.Start(fmt.Sprintf(":%d", 8080)))
}

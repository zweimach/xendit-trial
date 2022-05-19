package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zweimach/xendit-trial/api"
	"github.com/zweimach/xendit-trial/utils"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${status} ${uri} ${error}\n",
	}))
	e.Use(middleware.Recover())
	e.Validator = utils.NewCustomValidator(validator.New())

	e.POST("/", api.AuditTransactions)

	if err := e.Start(":3000"); err != nil {
		e.Logger.Fatal(err)
	}
}

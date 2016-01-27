package common

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type App struct {
	router *echo.Echo
}

func (app *App) Init() error {
	log.Printf("httpapi init")
	app.router = echo.New()
	app.router.Use(middleware.Logger(), middleware.Recover())
	app.router.Get("/balances", app.handleGetBalances)
	return nil
}

func (app *App) Run(addr string) {
	app.router.Run(addr)
}

func (app *App) handleGetBalances(ctx *echo.Context) error {
	balances, err := app.requestBalances()
	if err != nil {
		return err
	}
	ctx.JSON(200, balances)
	return nil
}

func (app *App) handlePostTransaction(ctx *echo.Context) error {
	return nil
}

func (app *App) requestBalances() (map[int64]int64, error) {
	return map[int64]int64{
		1: 100,
		2: 40500,
	}, nil
}
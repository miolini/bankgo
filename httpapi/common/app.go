package common

import (
	"log"

	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/miolini/bankgo/rpc/client"
)

type BalanceEntry struct {
	UserID int64
	Value  int64
}

type App struct {
	router *echo.Echo
	client *client.BalanceStorageClient
}

func (app *App) Init(rpcAddr string) error {
	log.Printf("httpapi init")
	client, err := client.Connect(rpcAddr)
	if err != nil {
		return err
	}
	app.client = client
	app.router = echo.New()
	app.router.Use(middleware.Logger(), middleware.Recover())
	app.router.Use(echoJsonCheckErrorMW())
	app.router.Get("/balances", app.handleGetBalances)
	app.router.Post("/transaction", app.handlePostTransaction)
	return nil
}

func (app *App) Run(addr string) {
	app.router.Run(addr)
}

func (app *App) handleGetBalances(ctx *echo.Context) error {
	records, err := app.client.AllBalances()
	if err != nil {
		return err
	}
	replyJson(ctx, records)
	return nil
}

func (app *App) handlePostTransaction(ctx *echo.Context) error {
	var balance BalanceEntry
	err := json.NewDecoder(ctx.Request().Body).Decode(&balance)
	if err != nil {
		return replyJsonError(ctx, fmt.Errorf("json parse error: %s", err))
	}
	if balance.UserID < 1 {
		return replyJsonError(ctx, "UserID must be greater than 0")
	}
	balance.Value, err = app.client.SetValue(balance.UserID, balance.Value)
	if err != nil {
		return replyJsonError(ctx, err)
	}
	return replyJson(ctx, balance)
}

func replyJson(ctx *echo.Context, v interface{}) error {
	return ctx.JSON(200, map[string]interface{}{"response": v})
}

func replyJsonError(ctx *echo.Context, err interface{}) error {
	return ctx.JSON(400, map[string]interface{}{"error": fmt.Sprintf("%s", err)})
}

func echoJsonCheckErrorMW() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			err := h(c)
			if err != nil {
				return replyJsonError(c, err)
			}
			return nil
		}
	}
}

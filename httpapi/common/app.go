package common

import (
	"bytes"
	"log"
	"io"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/miolini/bankgo/core"
	"encoding/json"
	"fmt"
)

type App struct {
	router *echo.Echo
}

type BalanceEntry struct {
	UserID int64 `json:"UserID"`
	Value  int64
}

type BalanceMap map[int64]int64

func (bm BalanceMap) MarshalJson() ([]byte, error) {
	var err error
	buf := &bytes.Buffer{}
	_, err = io.WriteString(buf, "{")
	if err != nil {
		return nil, err
	}
	for balanceId, balanceValue := range bm {
		io.WriteString(buf, "  \"")
		io.WriteString(buf, strconv.FormatInt(balanceId, 10))
		io.WriteString(buf, "\": ")
		io.WriteString(buf, strconv.FormatInt(balanceValue, 10))
		io.WriteString(buf, ", \n")
	}
	buf.Truncate(buf.Len() - 3)
	io.WriteString(buf, "\n")
	_, err = io.WriteString(buf, "}")
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (app *App) Init() error {
	log.Printf("httpapi init")
	app.router = echo.New()
	app.router.Use(middleware.Logger(), middleware.Recover())
	app.router.Use(core.EchoJsonCheckErrorMW())
	app.router.Get("/balances", app.handleGetBalances)
	app.router.Post("/transaction", app.handlePostTransaction)
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
	log.Printf("balances: %#v", balances)
	data, err := balances.MarshalJson()
	if err != nil {
		return err
	}
	log.Printf("data: %s", data)
	ctx.Response().Header().Add("Content-Type", "application/json")
	ctx.Response().WriteHeader(200)
	ctx.Response().Write(data)
	return nil
}

func (app *App) handlePostTransaction(ctx *echo.Context) error {
	var balance BalanceEntry
	err := json.NewDecoder(ctx.Request().Body).Decode(&balance)
	if err != nil {
		return fmt.Errorf("json parse error: %s", err)
	}
	ctx.JSON(200, balance)
	return nil
}

func (app *App) requestBalances() (BalanceMap, error) {
	return BalanceMap{
		1: 100,
		2: 40500,
	}, nil
}
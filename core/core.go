package core

import "fmt"
import "github.com/labstack/echo"

// EchoReplyJson reply with json standart structure with filed response
func EchoReplyJson(ctx *echo.Context, v interface{}) error {
	return ctx.JSON(200, map[string]interface{}{"response": v})
}

// EchoReplyJsonError reply with json standart error structure
func EchoReplyJsonError(ctx *echo.Context, err interface{}) error {
	return ctx.JSON(400, map[string]interface{}{"error": fmt.Sprintf("%s", err)})
}

func EchoJsonCheckErrorMW() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			err := h(c)
			if err != nil {
				return EchoReplyJsonError(c, err)
			}
			return nil
		}
	}
}

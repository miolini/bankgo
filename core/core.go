package core

import "fmt"
import "github.com/labstack/echo"

// EchoReplyJsonError reply with json standart error structure
func EchoReplyJsonError(ctx *echo.Context, err interface{}) {
	ctx.JSON(400, map[string]interface{}{"error": fmt.Sprintf("%s", err)})
}
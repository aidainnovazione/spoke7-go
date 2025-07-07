package request_id

import (
	"context"

	"github.com/labstack/echo/v4"
)

const RequestIdCtxKey string = "request_id"

func RequestIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqID := c.Request().Header.Get("X-Request-Id")

		ctx := context.WithValue(c.Request().Context(), RequestIdCtxKey, reqID)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}

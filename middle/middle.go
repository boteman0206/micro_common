package middle

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var TraceId = "trace_id"

func TraceIdMiddle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		uid, _ := uuid.NewUUID()
		c.Set(TraceId, uid.String())
		return next(c)
	}
}

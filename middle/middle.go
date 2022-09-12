package middle

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"micro_api/micro_common/es"
)

var TraceId = "trace_id"

func TraceIdMiddle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		method := c.Request().Method
		es.MyLog.Debug("=======", method)
		es.MyLog.Info("method:", method)
		uid, _ := uuid.NewUUID()
		c.Set(TraceId, uid.String())
		return next(c)
	}
}

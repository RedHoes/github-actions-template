package route

import (
	"net/http"

	"github.com/c0x12c/backend-go/pkg/api/response"
	"github.com/labstack/echo/v4"
)

type HealthRouter struct {
}

func (r *HealthRouter) Configure(e *echo.Echo) {
	e.GET("/api/health", r.HealthCheck)
}

func NewHealthRouter() *HealthRouter {
	return &HealthRouter{}
}

// HealthCheck returns the health status of the API
func (r *HealthRouter) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, &response.HealthResponse{
		Status:  "up",
		Version: "1.0.0",
	})
}

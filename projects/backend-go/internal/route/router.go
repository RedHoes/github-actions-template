package route

import "github.com/labstack/echo/v4"

type Router interface {
	// Configure configures the router
	Configure(e *echo.Echo)
}

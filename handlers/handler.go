package handlers

import "github.com/labstack/echo/v4"

func CustomResponseHeaders(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
		c.Response().Header().Set(echo.HeaderAccessControlAllowHeaders, "DELETE,GET,HEAD,OPTIONS,PATCH,POST,PUT")
		return next(c)
	}
}

func Example(name string) string {
	return "Hello " + name
}

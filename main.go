package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func main() {
	e := echo.New()

	// Little bit of middlewares for housekeeping
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	// This will initiate our template renderer
	NewTemplateRenderer(e, "templates/*.html")

	e.GET("/", func(e echo.Context) error {
		return e.Render(http.StatusOK, "index", nil)
	})

	e.GET("/about", func(e echo.Context) error {
		res := map[string]interface{}{
			"LinkedIn": "http://linkedin.com/in/zacpollack/",
		}
		return e.Render(http.StatusOK, "about", res)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

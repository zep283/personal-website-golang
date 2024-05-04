package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func main() {
	e := echo.New()

	// Little bit of middleware for housekeeping
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	// This will initiate our template renderer
	NewTemplateRenderer(e, "templates/*.html")

	mediumStories := ParseMediumRSSFeed()

	info := map[string]interface{}{
		"LinkedIn": "http://linkedin.com/in/zacpollack/",
		"GitHub":   "https://github.com/zep283",
		"Medium":   "https://medium.com/@zep283",
		"Stories":  mediumStories,
	}

	e.GET("/", func(e echo.Context) error {
		return e.Render(http.StatusOK, "index", info)
	})

	e.GET("/about", func(e echo.Context) error {
		return e.Render(http.StatusOK, "about", info)
	})

	e.GET("/blog", func(e echo.Context) error {
		return e.Render(http.StatusOK, "blog", info)
	})

	e.GET("/projects", func(e echo.Context) error {
		return e.Render(http.StatusOK, "projects", info)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

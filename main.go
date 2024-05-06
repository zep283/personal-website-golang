package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

type HtmlPage struct {
	Name string
	Path string
}

func main() {
	e := echo.New()

	// Little bit of middleware for housekeeping
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	htmlPages := []HtmlPage{
		{
			Name: "index",
			Path: "templates/index.html",
		},
		{
			Name: "about",
			Path: "templates/about.html",
		},
		{
			Name: "blog",
			Path: "templates/blog.html",
		},
		{
			Name: "projects",
			Path: "templates/projects.html",
		},
	}

	// This will initiate our template renderer
	e.Renderer = RegisterTemplates(htmlPages)

	mediumStories := ParseMediumRSSFeed()

	info := struct {
		LinkedIn string
		GitHub   string
		Medium   string
		Stories  []Story
	}{
		LinkedIn: "http://linkedin.com/in/zacpollack/",
		GitHub:   "https://github.com/zep283",
		Medium:   "https://medium.com/@zep283",
		Stories:  mediumStories,
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

package main

import (
	"net/http"

	"github.com/zep283/personal-website-golang/internal/common"
	"github.com/zep283/personal-website-golang/internal/medium"
	"github.com/zep283/personal-website-golang/internal/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func main() {
	e := echo.New()

	e.Static("/static", "../web/assets")

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	htmlPages := []common.HtmlPage{
		{
			Name: "index",
			Path: "../web/index.html",
		},
		{
			Name: "about",
			Path: "../web/about.html",
		},
		{
			Name: "blog",
			Path: "../web/blog.html",
		},
		{
			Name: "projects",
			Path: "../web/projects.html",
		},
	}

	e.Renderer = template.RegisterTemplates(htmlPages)

	mediumStories := medium.ParseMediumRSSFeed()

	info := struct {
		LinkedIn      string
		GitHub        string
		Medium        string
		Stories       []common.Story
		LatestStories []common.Story
	}{
		LinkedIn:      "http://linkedin.com/in/zacpollack/",
		GitHub:        "https://github.com/zep283",
		Medium:        "https://medium.com/@zep283",
		Stories:       mediumStories,
		LatestStories: mediumStories[0:3],
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

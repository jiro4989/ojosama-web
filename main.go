package main

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Env struct {
	Port string
}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

var (
	// アプリケーションの起動時刻
	startAt time.Time
)

func main() {
	location := time.FixedZone("Asia/Tokyo", 9*60*60)
	startAt = time.Now().In(location)

	e := echo.New()
	env := newEnv()

	t, err := loadTemplates()
	if err != nil {
		e.Logger.Fatal(err)
	}

	r := &TemplateRenderer{
		templates: t,
	}
	e.Renderer = r

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	e.GET("/", pageGetRoot)
	e.GET("/tos", pageGetTOS)
	e.POST("/api/ojosama", apiPostOjosama)
	e.GET("/api/ping", apiGetPing)
	e.GET("/api/version", apiGetVersion)

	e.Logger.Fatal(e.Start(":" + env.Port))
}

func newEnv() Env {
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	return Env{
		Port: port,
	}
}

func loadTemplates() (*template.Template, error) {
	t := template.New("")
	err := filepath.Walk("./public/views", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			_, err = t.ParseFiles(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return t, nil
}

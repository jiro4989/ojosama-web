package main

import (
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jiro4989/ojosama"
	"github.com/labstack/echo/v4"
)

type Data struct {
	Title       string
	Description string
}

func pageGetRoot(c echo.Context) error {
	data := Data{
		Title:       "Ojosama web converter",
		Description: "テキストを壱百満天原サロメお嬢様風の口調に変換します。",
	}
	return c.Render(http.StatusOK, "index.tmpl.html", data)
}

func pageGetTOS(c echo.Context) error {
	data := Data{
		Title:       "利用規約 - Ojosama web converter",
		Description: "利用規約です。アプリを使う際のお約束事を記載しています。",
	}
	return c.Render(http.StatusOK, "tos.tmpl.html", data)
}

func apiPostOjosama(c echo.Context) error {
	var r RequestPostOjosama
	if err := c.Bind(&r); err != nil {
		return err
	}
	if r.Text == "" {
		return nil
	}

	result, err := ojosama.Convert(r.Text, nil)
	if err != nil {
		return err
	}

	var resp ResponseOjosama
	resp.Result = result
	return c.JSON(http.StatusOK, resp)
}

func apiGetPing(c echo.Context) error {
	resp := ResponsePing{Status: "OK"}
	return c.JSON(http.StatusOK, resp)
}

func apiGetVersion(c echo.Context) error {
	r, err := os.Open("revision.txt")
	if err != nil {
		return err
	}
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	rev := strings.TrimSpace(string(b))
	resp := ResponseVersion{
		StartAt: startAt.Format(time.RFC3339),
		Version: rev,
	}
	return c.JSON(http.StatusOK, resp)
}

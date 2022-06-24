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

func pageGetRoot(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func pageGetTOS(c echo.Context) error {
	return c.Render(http.StatusOK, "tos.html", nil)
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

package main

import (
	"net/http"

	"github.com/jiro4989/ojosama"
	"github.com/labstack/echo/v4"
)

func pageGetRoot(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
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

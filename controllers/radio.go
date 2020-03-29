package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Radio(c echo.Context) error {
	response, erro := http.Get("https://painel.audiovox.pw/api-json/Njc2OCsw")
	if erro != nil {
		panic(erro)
	}

	responseJson, erro := ioutil.ReadAll(response.Body)

	if erro != nil {
		panic(erro)
	}

	return c.String(http.StatusOK, string(responseJson))
}

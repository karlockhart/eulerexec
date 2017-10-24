package api

import (
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/coderzer0h/eulerexec/pkg/host"

	"github.com/labstack/echo"
)

func Start(wg *sync.WaitGroup) {
	e := echo.New()
	e.Static("/", "app/dist")
	e.POST("api/go/lint", Lint)
	e.POST("api/go/fmt", Format)
	e.POST("api/go/run", Run)
	e.Logger.Error(e.Start(":1323"))
	wg.Done()
}

func Lint(c echo.Context) error {
	return c.String(http.StatusOK, "Lint")
}

func Format(c echo.Context) error {
	b, e := ioutil.ReadAll(c.Request().Body)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}
	f, e := host.Format(b)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}
	return c.String(http.StatusOK, string(f))
}

func Run(c echo.Context) error {
	b, e := ioutil.ReadAll(c.Request().Body)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}
	f, e := host.Run(b)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}
	return c.String(http.StatusOK, string(f))
}

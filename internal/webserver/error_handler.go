package webserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)

	response := customHTTPErrorPageModel{
		Code:    code,
		Message: err.Error(),
	}

	c.Render(http.StatusOK, "error.html", response)
	// if err := c.Render(http.StatusOK, "error.html", response); err != nil {
	// c.Logger().Error(err)
	// }

	// errorPage := fmt.Sprintf("%d.html", code)
	// if err := c.File(errorPage); err != nil {
	// 	c.Logger().Error(err)
	// }
}

type customHTTPErrorPageModel struct {
	Code    int
	Message string
}

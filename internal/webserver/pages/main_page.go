package pages

import (
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

func MainPage(c echo.Context) error {

	if _, err := os.Stat("frontend/dist/index.html"); os.IsNotExist(err) {
		return c.String(http.StatusInternalServerError, "index file not found")
	}

	index_bytes, err := os.ReadFile("frontend/dist/index.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	correct_base_url := "var __base_url__ = 'http://localhost:1314';"

	index_lines := strings.Split(string(index_bytes), "\n")

	lines := []string{}
	for _, line := range index_lines {
		if strings.Contains(line, "var __base_url__") {
			lines = append(lines, correct_base_url)
			continue
		}

		lines = append(lines, line)
	}

	output := strings.Join(lines, "\n")

	return c.HTML(http.StatusOK, output)
}

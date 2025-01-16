package pages

import (
	"fmt"
	"hdu/internal/webserver/utils"
	"io"
	"io/fs"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type MainPageConfiguration struct {
	BaseURL  string
	Registry bool
}

const main_page_inject_lines_template = `
	__base_url__ = '%s';
	__version__ = '%s';
	__registry__ = %s;
`

func MainPage(c echo.Context, config MainPageConfiguration, storage fs.FS) error {

	// inject variables to index.html, content of storage - see frontend/dist/
	f, err := storage.Open("index.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	base_url := utils.ResolveURL("/", config.BaseURL)

	registry := "false"
	if config.Registry {
		registry = "true"
	}

	index_lines := strings.Split(string(content), "\n")

	lines := []string{}
	for _, line := range index_lines {
		if strings.Contains(line, "// INJECT") {
			injected_lines := fmt.Sprintf(main_page_inject_lines_template, base_url, "0.0.0", registry)
			lines = append(lines, injected_lines)
			continue
		}

		lines = append(lines, line)
	}

	output := strings.Join(lines, "\n")

	return c.HTML(http.StatusOK, output)
}

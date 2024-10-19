package webserver

import (
	"errors"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// type Template struct {
// 	templates *template.Template

// 	tts map[string]*template.Template
// }

// func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
// 	return t.templates.ExecuteTemplate(w, name, data)
// }

// // --- составление списка шаблонов(обход всего каталога)
// func makeTemplatesList(root_path string) []string {

// 	var all_templates []string

// 	filepath.Walk(root_path, func(path string, info os.FileInfo, err error) error {
// 		filename := info.Name()
// 		if strings.HasSuffix(filename, ".html") {
// 			fmt.Println(path)
// 			all_templates = append(all_templates, path)
// 		}
// 		return nil
// 	})
// 	return all_templates
// }

// TODO: доработать!
// https://stackoverflow.com/questions/11467731/is-it-possible-to-have-nested-templates-in-go-using-the-standard-library
// составляем цепочки шаблонов пока вручную
type Templater struct {
	tts map[string]*template.Template
}

func NewTemplater() *Templater {
	var tpls = make(map[string]*template.Template)

	base_template := template.Must(template.New("base_template").ParseFiles("./views/base.html"))

	from_base_templates := []string{
		"config.html",
		"configs.html",
		"container.html",
		"containers.html",
		"image.html",
		"images.html",
		"main.html",
		"network.html",
		"networks.html",
		"secret.html",
		"secrets.html",
		"volume.html",
		"volumes.html",
	}

	for _, name := range from_base_templates {
		tpls[name] = template.Must(template.Must(base_template.Clone()).ParseFiles("./views/" + name))
	}

	tpls["error.html"] = template.Must(template.New("error").ParseFiles("./views/error.html"))

	return &Templater{
		tts: tpls,
	}
}

func (t *Templater) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	if tt, ok := t.tts[name]; ok {
		return tt.ExecuteTemplate(w, name, data)
	}
	return errors.New("not found template definition: " + name)
}

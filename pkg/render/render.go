package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func GeneratedHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string

	for _, file := range filenames {
		files = append(files, fmt.Sprintf("resources/views/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "resources/views/layout.html", data)
}

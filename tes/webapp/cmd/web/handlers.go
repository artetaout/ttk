package main

import (
	"html/template"
	"net/http"
	"path"
)

var path2Templates = "./templates/"

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "home.page.gohtml", &TemplateData{})

}

type TemplateData struct {
	IP   string
	Data map[string]string
}

func (app *application) render(w http.ResponseWriter, r *http.Request, t string, data *TemplateData) error {
	// parse template
	parsedTemplate, err := template.ParseFiles(path.Join(path2Templates, t))
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return err
	}

	// execute template
	err = parsedTemplate.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}

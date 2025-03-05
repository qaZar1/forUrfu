package renders

import (
	"html/template"
	"net/http"
)

func renderRegister(w http.ResponseWriter, templateName string) {
	tmpl, err := template.ParseFiles("internal/templates/" + templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RenderRegisterFor(w http.ResponseWriter, r *http.Request) {
	renderRegister(w, "employers/register_for_employers.html")
}

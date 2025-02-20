package renders

import (
	"html/template"
	"net/http"

	"github.com/qaZar1/HHforURFU/web/internal/api"
	"github.com/qaZar1/HHforURFU/web/internal/models"
)

type Profile struct {
	Employer models.Employer
}

func renderProfile(w http.ResponseWriter, templateName string, data Profile) {
	tmpl, err := template.ParseFiles("internal/temp/" + templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data) // Здесь можно передать данные в шаблон, если нужно
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RenderProfile(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username_employers")
	if err != nil {
		return
	}

	apiEmployers := api.NewApiEmployers()

	employer, err := apiEmployers.CheckEmployerByUsername(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	renderProfile(w, "employers/profile.html", Profile{
		Employer: employer,
	})
}

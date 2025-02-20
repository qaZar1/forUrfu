package renders

import (
	"html/template"
	"net/http"

	"github.com/qaZar1/HHforURFU/web/internal/api"
	"github.com/qaZar1/HHforURFU/web/internal/models"
)

type Profile struct {
	Seeker models.Seeker
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
	cookie, err := r.Cookie("username")
	if err != nil {
		return
	}

	apiSeekers := api.NewApiSeekers()

	seeker, err := apiSeekers.CheckSeeker(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	renderProfile(w, "seekers/profile.html", Profile{
		Seeker: seeker,
	})
}

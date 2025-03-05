package renders

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/qaZar1/forUrfu/web/internal/api"
	"github.com/qaZar1/forUrfu/web/internal/models"
)

type Profile struct {
	Employer models.Employer
	Rating   models.Rating
}

func renderProfile(w http.ResponseWriter, templateName string, data Profile) {
	tmpl, err := template.ParseFiles("internal/templates/" + templateName)
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
	username, err := r.Cookie("username_employers")
	if err != nil {
		return
	}

	apiEmployers := api.NewApiEmployers()
	apiRatings := api.NewApiRatings()

	employer, err := apiEmployers.CheckEmployerByUsername(username.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	rating, err := apiRatings.CheckRating(username.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if rating == (models.Rating{}) {
		rating.Rating = 0
	}

	ratingStr := fmt.Sprintf("%.1f", rating.Rating)
	ratingFloat, err := strconv.ParseFloat(ratingStr, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	rating.Rating = ratingFloat

	renderProfile(w, "employers/profile.html", Profile{
		Employer: employer,
		Rating:   rating,
	})
}

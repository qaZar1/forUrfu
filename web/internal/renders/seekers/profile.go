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
	Seeker models.Seeker
	Rating models.Rating
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
	cookie, err := r.Cookie("username")
	if err != nil {
		return
	}

	apiSeekers := api.NewApiSeekers()
	apiRatings := api.NewApiRatings()

	seeker, err := apiSeekers.CheckSeeker(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	rating, err := apiRatings.CheckRating(seeker.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	ratingStr := fmt.Sprintf("%.1f", rating.Rating)
	ratingFloat, err := strconv.ParseFloat(ratingStr, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	rating.Rating = ratingFloat

	renderProfile(w, "seekers/profile.html", Profile{
		Seeker: seeker,
		Rating: rating,
	})
}

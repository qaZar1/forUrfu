package renders

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/qaZar1/forUrfu/web/internal/api"
	"github.com/qaZar1/forUrfu/web/internal/models"
)

type Response struct {
	Response models.Response
	Vacancy  models.Vacancy
	Seeker   models.Seeker
	Rating   models.Rating
	IsRated  bool
}

func renderResponse(w http.ResponseWriter, templateName string, data Response) {
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

func RenderResponse(w http.ResponseWriter, r *http.Request) {
	id_str := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(id_str, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	apiVacancies := api.NewApiVacancies()
	apiResponses := api.NewApiResponses()
	apiSeekers := api.NewApiSeekers()
	apiRatings := api.NewApiRatings()
	apiEmloyers := api.NewApiEmployers()

	resp, err := apiResponses.GetResponsesByID(id)
	if err != nil {
		http.Error(w, "Can not get response by username", http.StatusBadRequest)
		return
	}

	vacancy, err := apiVacancies.GetVacancyByVacancyID(resp.VacancyID)
	if err != nil {
		http.Error(w, "Can not get vacancy by id", http.StatusBadRequest)
		return
	}

	seeker, err := apiSeekers.CheckSeeker(resp.Username)
	if err != nil {
		http.Error(w, "Can not get seeker by username", http.StatusBadRequest)
		return
	}

	employer, err := apiEmloyers.CheckEmployer(vacancy.EmployerID)
	if err != nil {
		http.Error(w, "Can not get seeker by username", http.StatusBadRequest)
		return
	}

	rating, err := apiRatings.CheckRating(seeker.Username)
	if err != nil {
		http.Error(w, "Can not get rating by username", http.StatusBadRequest)
		return
	}

	isRated := false
	if rating.FromUser == seeker.Username && rating.ToUser == employer.Username {
		isRated = true
	}

	data := Response{
		Response: resp,
		Vacancy:  vacancy,
		Seeker:   seeker,
		Rating:   rating,
		IsRated:  isRated,
	}

	renderResponse(w, "seekers/response.html", data)
}

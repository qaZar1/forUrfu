package renders

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/qaZar1/forUrfu/web/internal/api"
	"github.com/qaZar1/forUrfu/web/internal/models"
)

type Vacancy struct {
	Vacancy models.Vacancy
	Applied bool
}

func renderResp(w http.ResponseWriter, templateName string, data Vacancy) {
	tmpl, err := template.ParseFiles("internal/templates/" + templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RenderVacancy(w http.ResponseWriter, r *http.Request) {
	username, err := r.Cookie("username")
	if err != nil {
		return
	}

	id_str := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(id_str, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	apiVacancies := api.NewApiVacancies()
	apiResp := api.NewApiResponses()
	responses, err := apiResp.GetAllResponsesByUsername(username.Value)
	isApplied := false
	for _, resp := range responses {
		if resp.VacancyID == id {
			isApplied = true
		}
	}
	vacancy, _ := apiVacancies.GetVacancyByVacancyID(id)

	data := Vacancy{
		Vacancy: vacancy,
		Applied: isApplied,
	}
	renderResp(w, "seekers/vacancy.html", data)
}

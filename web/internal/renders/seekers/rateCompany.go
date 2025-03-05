package renders

import (
	"net/http"
	"strconv"

	"github.com/qaZar1/forUrfu/web/internal/api"
	"github.com/qaZar1/forUrfu/web/internal/models"
	"github.com/sirupsen/logrus"
)

type Rate struct {
	Seeker models.Seeker
	Rating models.Rating
}

func RateCompanyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logrus.Errorf("Error: invalid method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	username, err := r.Cookie("username")
	if err != nil {
		logrus.Errorf("Error: invalid username")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ratingStr := r.FormValue("rating")
	rating, err := strconv.Atoi(ratingStr)
	if err != nil {
		logrus.Errorf("Error: invalid username")
		return
	}

	vacancyIdStr := r.FormValue("vacancy_id")
	vacancyId, err := strconv.Atoi(vacancyIdStr)
	if err != nil {
		logrus.Errorf("Error: invalid username")
		return
	}

	// Проверяем корректность рейтинга (1-5)
	if rating < 1 || rating > 5 {
		http.Error(w, "Недопустимое значение рейтинга", http.StatusBadRequest)
		return
	}

	apiRatings := api.NewApiRatings()
	apiVacancies := api.NewApiVacancies()
	apiEmployers := api.NewApiEmployers()

	vacancy, err := apiVacancies.GetVacancyByVacancyID(int64(vacancyId))
	if err != nil {
		logrus.Errorf("Error: invalid username")
		return
	}

	employer, err := apiEmployers.CheckEmployer(vacancy.EmployerID)
	if err != nil {
		logrus.Errorf("Error: invalid username")
		return
	}

	ratingModel := models.Rating{
		FromUser: username.Value,
		ToUser:   employer.Username,
		Rating:   float64(rating),
	}

	ok, err := apiRatings.AddRating(ratingModel)
	if err != nil {
		logrus.Errorf("Error: invalid username")
		return
	}

	if !ok {
		http.Error(w, "Невозможно добавить рейтинг", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

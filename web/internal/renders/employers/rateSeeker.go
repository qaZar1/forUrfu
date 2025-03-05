package renders

import (
	"net/http"
	"strconv"

	"github.com/qaZar1/forUrfu/web/internal/api"
	"github.com/qaZar1/forUrfu/web/internal/models"
	"github.com/sirupsen/logrus"
)

type RateSeeker struct {
	Seeker models.Seeker
	Rating models.Rating
}

func RateSeekerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logrus.Errorf("Error: invalid method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	username, err := r.Cookie("username_employers")
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

	seekerUsername := r.FormValue("seeker-username")

	// Проверяем корректность рейтинга (1-5)
	if rating < 1 || rating > 5 {
		http.Error(w, "Недопустимое значение рейтинга", http.StatusBadRequest)
		return
	}

	apiRatings := api.NewApiRatings()

	ratingModel := models.Rating{
		FromUser: username.Value,
		ToUser:   seekerUsername,
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

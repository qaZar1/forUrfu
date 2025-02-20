package renders

import (
	"net/http"

	"github.com/qaZar1/HHforURFU/web/internal/api"
	"github.com/qaZar1/HHforURFU/web/internal/models"
	"github.com/sirupsen/logrus"
)

func SaveUpdate(w http.ResponseWriter, r *http.Request) {
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

	apiSeeker := api.NewApiSeekers()

	seeker, err := apiSeeker.CheckSeeker(username.Value)
	if err != nil {
		logrus.Errorf("Error: cant checked seeker")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	updatedResume := r.FormValue("resume")
	if updatedResume == "" || updatedResume == seeker.Resume {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	updatedSeeker := &models.Seeker{
		Resume: updatedResume,
	}

	if err := apiSeeker.EditSeeker(username.Value, updatedSeeker); err != nil {
		logrus.Errorf("Error: cant update seeker")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

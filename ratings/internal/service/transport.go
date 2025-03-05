package service

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Impisigmatus/service_core/utils"
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/forUrfu/ratings/autogen/server"
	"github.com/qaZar1/forUrfu/ratings/internal/database"
	"github.com/qaZar1/forUrfu/ratings/internal/models"

	jsoniter "github.com/json-iterator/go"
)

type Transport struct {
	db database.Database
}

func NewTransport(db *sqlx.DB) server.ServerInterface {
	return &Transport{
		db: *database.NewDatabase(db),
	}
}

// Set godoc
//
// @Router /api/ratings/{username} [get]
// @Summary Получение ответов по username
// @Description При обращении, возвращаются все ответы, которые есть у username
//
// @Tags APIs
// @Produce      application/json
// @Param 	username	path	string	true	"ID пользователя"
//
// @Success 200 {array} rating "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) GetApiRatingsUsername(w http.ResponseWriter, r *http.Request, username string) {
	ratings, err := transport.db.CheckRatingByUsername(username)
	if err != nil {
		if ratings == (models.Rating{}) {
			utils.WriteString(w, http.StatusNoContent, err, "В базе нет рейтингов")
		}

		utils.WriteString(w, http.StatusNoContent, err, "Внутренняя ошибка сервера")
		return
	}

	utils.WriteObject(w, ratings)
}

// Set godoc
//
// @Router /api/ratings/addRating [post]
// @Summary Добавление отклика в БД
// @Description При обращении, добавляется отклик в БД по телу запрсоа
//
// @Tags APIs
// @Produce      application/json
// @Param 	request	body	rating	true	"Тело запроса"
//
// @Success 200 {object} nil "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) PostApiRatingsAddRating(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, fmt.Errorf("Invalid read body: %s", err), "Не удалось прочитать тело запроса")
		return
	}

	var rating models.Rating
	if err := jsoniter.Unmarshal(data, &rating); err != nil {
		utils.WriteString(w, http.StatusBadRequest, fmt.Errorf("Invalid parse body: %s", err), "Не удалось распарсить тело запроса формата JSON")
		return
	}

	ok, err := transport.db.AddRating(rating)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось добавить вакансию")
		return
	}

	if !ok {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось добавить вакансию")
		return
	}

	utils.WriteNoContent(w)
}

package api

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/forUrfu/web/internal/models"
)

type APIRatings struct {
	client *resty.Client
}

func NewApiRatings() *APIRatings {
	return &APIRatings{
		client: resty.New().SetBaseURL("http://localhost:8008/api").SetTimeout(1*time.Minute).SetBasicAuth("dev", "test").SetDisableWarn(true),
	}
}

func (api *APIRatings) AddRating(rating models.Rating) (bool, error) {
	const endpoint = "/ratings/addRating"

	resp, err := api.client.R().SetBody(rating).Post(endpoint)
	if err != nil {
		return false, err
	}

	if resp.IsError() {
		return false, fmt.Errorf("failed to add rating: %s", resp.Status())
	}

	return true, nil
}

func (api *APIRatings) CheckRating(username string) (models.Rating, error) {
	endpoint := "/ratings/%s"

	endpoint = fmt.Sprintf(endpoint, username)
	resp, err := api.client.R().Get(endpoint)
	if err != nil {
		return models.Rating{}, err
	}

	if len(resp.Body()) == 0 {
		return models.Rating{}, nil
	}

	rating := models.Rating{}
	if err := jsoniter.Unmarshal(resp.Body(), &rating); err != nil {
		return models.Rating{}, err
	}

	return rating, nil
}

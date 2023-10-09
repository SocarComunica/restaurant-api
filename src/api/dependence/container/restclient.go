package container

import (
	"restaurant-api/internal/platform/restful"
	"restaurant-api/src/api/config/static"
	"restaurant-api/src/api/config/static/model"
)

type RestClient struct {
	Market restful.RestClient
}

func CreateRestClient() RestClient {
	config := static.GetConfig()
	return RestClient{
		Market: createRestClient(config.MarketClient)}
}

func createRestClient(config model.RestClient) restful.RestClient {
	return restful.NewRestClient(config.BaseURL)
}

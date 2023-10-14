package container

import (
	"restaurant-api/api/config/static"
	"restaurant-api/api/config/static/model"
	"restaurant-api/internal/platform/restful"
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

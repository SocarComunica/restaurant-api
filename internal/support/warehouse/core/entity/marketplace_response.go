package entity

type MarketPlaceResponse struct {
	Message string         `json:"message"`
	Data    map[string]int `json:"data"`
}

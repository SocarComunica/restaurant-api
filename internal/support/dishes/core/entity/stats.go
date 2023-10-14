package entity

type Stats struct {
	Delivered  int `json:"delivered"`
	Finished   int `json:"finished"`
	InProgress int `json:"inProgress"`
	Queued     int `json:"queued"`
}

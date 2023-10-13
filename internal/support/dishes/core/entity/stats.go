package entity

type Stats struct {
	Delivered int `json:"delivered"`
	Finished  int `json:"finished"`
	Queued    int `json:"queued"`
}

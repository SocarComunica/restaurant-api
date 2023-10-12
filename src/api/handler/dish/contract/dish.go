package contract

type Dish struct {
	ID     int            `json:"id"`
	Name   string         `json:"name"`
	Recipe map[string]int `json:"recipe"`
	Stats  Stats          `json:"stats"`
}

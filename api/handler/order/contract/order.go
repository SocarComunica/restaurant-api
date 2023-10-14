package contract

import "time"

type Order struct {
	ID        int       `json:"id"`
	DishID    int       `json:"dish_id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Status    string    `json:"status"`
}

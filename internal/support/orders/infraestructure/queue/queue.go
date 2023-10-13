package queue

import "restaurant-api/internal/support/orders/core/entity"

type Queue []*entity.Order

func (q *Queue) Enqueue(order entity.Order) {
	*q = append(*q, &order)
}

func (q *Queue) Dequeue() (*entity.Order, bool) {
	if len(*q) == 0 {
		return nil, false
	}
	order := (*q)[0]
	*q = (*q)[1:]
	return order, true
}

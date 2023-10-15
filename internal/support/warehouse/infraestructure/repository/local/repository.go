package local

type Repository struct {
	ingredients map[string]int
}

func NewRepository() Repository {
	return Repository{
		ingredients: map[string]int{
			"tomato":  5,
			"lemon":   5,
			"potato":  5,
			"rice":    5,
			"ketchup": 5,
			"lettuce": 5,
			"onion":   5,
			"cheese":  5,
			"meat":    5,
			"chicken": 5,
		},
	}
}

func (r *Repository) GetInventory(name string) (int, error) {
	return r.ingredients[name], nil
}

func (r *Repository) UseInventory(name string, amount int) error {
	r.ingredients[name] -= amount
	return nil
}

func (r *Repository) AddInventory(name string, amount int) error {
	r.ingredients[name] += amount
	return nil
}

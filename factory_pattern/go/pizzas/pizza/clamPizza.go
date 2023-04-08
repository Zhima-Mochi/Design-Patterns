package pizza

type ClamPizza struct {
	PizzaBase
}

func NewClamPizza() *ClamPizza {
	return &ClamPizza{
		PizzaBase{
			Name:  "Clam Pizza",
			Dough: "Thin Crust",
			Sauce: "White Garlic Sauce",
			Toppings: []string{
				"Clams",
				"Grated parmesan cheese",
			},
		},
	}
}

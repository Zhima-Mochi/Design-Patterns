package pizza

type NYStyleCheesePizza struct {
	PizzaBase
}

func NewNYStyleCheesePizza() *NYStyleCheesePizza {
	return &NYStyleCheesePizza{
		PizzaBase{
			Name:  "NY Style Sauce and Cheese Pizza",
			Dough: "Thin Crust Dough",
			Sauce: "Marinara Sauce",
			Toppings: []string{
				"Grated Reggiano Cheese",
			},
		},
	}
}

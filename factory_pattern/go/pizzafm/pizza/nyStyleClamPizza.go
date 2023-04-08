package pizza

type NYSyleClamPizza struct {
	PizzaBase
}

func NewNYStyleClamPizza() *NYSyleClamPizza {
	return &NYSyleClamPizza{
		PizzaBase{
			Name:  "NY Style Clam Pizza",
			Dough: "Thin Crust Dough",
			Sauce: "Marinara Sauce",
			Toppings: []string{
				"Grated Reggiano Cheese",
				"Fresh Clams from Long Island Sound",
			},
		},
	}
}

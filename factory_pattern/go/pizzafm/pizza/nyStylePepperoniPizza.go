package pizza

type NYStylePepperoniPizza struct {
	PizzaBase
}

func NewNYStylePepperoniPizza() *NYStylePepperoniPizza {
	return &NYStylePepperoniPizza{
		PizzaBase{
			Name:  "NY Style Pepperoni Pizza",
			Dough: "Thin Crust Dough",
			Sauce: "Marinara Sauce",
			Toppings: []string{
				"Sliced Pepperoni",
				"Sliced Onion",
				"Grated parmesan cheese",
			},
		},
	}
}

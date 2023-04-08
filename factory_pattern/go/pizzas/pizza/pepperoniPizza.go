package pizza

type PepperoniPizza struct {
	PizzaBase
}

func NewPepperoniPizza() *PepperoniPizza {
	return &PepperoniPizza{
		PizzaBase{
			Name:  "Pepperoni Pizza",
			Dough: "Crust",
			Sauce: "Marinara sauce",
			Toppings: []string{
				"Sliced Pepperoni",
				"Sliced Onion",
				"Grated parmesan cheese",
			},
		},
	}
}

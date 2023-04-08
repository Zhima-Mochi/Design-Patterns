package pizza

type CheesePizza struct {
	PizzaBase
}

func NewCheesePizza() *CheesePizza {
	return &CheesePizza{
		PizzaBase{
			Name:  "Cheese Pizza",
			Dough: "Regular Crust",
			Sauce: "Marinara Pizza Sauce",
			Toppings: []string{
				"Fresh Mozzarella",
				"Parmesan",
			},
		},
	}
}

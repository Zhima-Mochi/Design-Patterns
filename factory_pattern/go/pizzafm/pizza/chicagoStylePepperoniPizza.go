package pizza

import "fmt"

type ChicagoStylePepperoniPizza struct {
	PizzaBase
}

func NewChicagoStylePepperoniPizza() *ChicagoStylePepperoniPizza {
	return &ChicagoStylePepperoniPizza{
		PizzaBase{
			Name:  "Chicago Style Pepperoni Pizza",
			Dough: "Extra Thick Crust Dough",
			Sauce: "Plum Tomato Sauce",
			Toppings: []string{
				"Shredded Mozzarella Cheese",
				"Black Olives",
				"Spinach",
				"Eggplant",
				"Sliced Pepperoni",
			},
		},
	}
}

func (p *ChicagoStylePepperoniPizza) Cut() {
	fmt.Println("Cutting the pizza into square slices")
}

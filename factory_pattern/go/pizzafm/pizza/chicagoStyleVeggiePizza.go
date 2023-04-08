package pizza

import "fmt"

type ChicagoStyleVeggiePizza struct {
	PizzaBase
}

func NewChicagoStyleVeggiePizza() *ChicagoStyleVeggiePizza {
	return &ChicagoStyleVeggiePizza{
		PizzaBase{
			Name:  "Chicago Style Deep Dish Veggie Pizza",
			Dough: "Extra Thick Crust Dough",
			Sauce: "Plum Tomato Sauce",
			Toppings: []string{
				"Shredded Mozzarella Cheese",
				"Black Olives",
				"Spinach",
				"Eggplant",
			},
		},
	}
}

func (p *ChicagoStyleVeggiePizza) Cut() {
	fmt.Println("Cutting the pizza into square slices")
}

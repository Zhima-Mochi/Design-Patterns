package pizza

type VeggiePizza struct {
	PizzaBase
}

func NewVeggiePizza() *VeggiePizza {
	return &VeggiePizza{
		PizzaBase{
			Name:  "Veggie Pizza",
			Dough: "Crust",
			Sauce: "Marinara sauce",
			Toppings: []string{
				"Shredded mozzarella",
				"Grated parmesan",
				"Diced onion",
				"Sliced mushrooms",
				"Sliced red pepper",
				"Sliced black olives",
			},
		},
	}
}

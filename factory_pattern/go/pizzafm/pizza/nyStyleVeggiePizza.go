package pizza

type nyStyleVeggiePizza struct {
	PizzaBase
}

func NewNYStyleVeggiePizza() *nyStyleVeggiePizza {
	return &nyStyleVeggiePizza{
		PizzaBase{
			Name:  "NY Style Veggie Pizza",
			Dough: "Thin Crust Dough",
			Sauce: "Marinara Sauce",
			Toppings: []string{
				"Grated Reggiano Cheese",
				"Garlic",
				"Onion",
				"Mushrooms",
				"Red Pepper",
			},
		},
	}
}

package pizza

import "fmt"

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
	String() string
}

type PizzaBase struct {
	Name     string
	Dough    string
	Sauce    string
	Toppings []string
}

func (p *PizzaBase) GetName() string {
	return p.Name
}

func (p *PizzaBase) Prepare() {
	fmt.Println("Prepare ", p.Name)
	fmt.Println("Tossing dough...")
	fmt.Println("Adding sauce...")
	fmt.Println("Adding toppings: ")
	for _, topping := range p.Toppings {
		fmt.Println("   ", topping)
	}
}

func (p *PizzaBase) Bake() {
	fmt.Println("Bake for 25 minutes at 350")
}

func (p *PizzaBase) Cut() {
	fmt.Println("Cut the pizza into diagonal slices")
}

func (p *PizzaBase) Box() {
	fmt.Println("Place pizza in official PizzaStore box")
}

func (p *PizzaBase) String() string {
	display := "---- " + p.Name + " ----\n"
	display += p.Dough + "\n"
	display += p.Sauce + "\n"
	for _, topping := range p.Toppings {
		display += topping + "\n"
	}
	return display
}

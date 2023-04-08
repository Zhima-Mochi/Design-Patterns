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
	fmt.Println("Preparing ", p.Name)
}

func (p *PizzaBase) Bake() {
	fmt.Println("Baking ", p.Name)
}

func (p *PizzaBase) Cut() {
	fmt.Println("Cutting ", p.Name)
}

func (p *PizzaBase) Box() {
	fmt.Println("Boxing ", p.Name)
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

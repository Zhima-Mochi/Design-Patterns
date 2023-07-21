package main

type Customer struct {
	waitress *Waitress
	cook     *Cook
	o        Order
}

func NewCustomer(waitress *Waitress, cook *Cook) *Customer {
	return &Customer{
		waitress: waitress,
		cook:     cook,
	}
}

func (c *Customer) CreateOrder() {
	c.o = func() {
		c.cook.MakeBurger()
		c.cook.MakeFries()
	}
}

func (c *Customer) Hungry() {
	c.waitress.TakeOrder(c.o)
}

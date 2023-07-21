package main

func main() {
	cook := NewCook()
	waitress := NewWaitress()
	customer := NewCustomer(waitress, cook)
	customer.CreateOrder()
	customer.Hungry()
}

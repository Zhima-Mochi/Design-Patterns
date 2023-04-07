package weather

type DisplayElement interface {
	Display()
}

type Observer interface {
	Update(temp, humidity, pressure float64)
}

type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

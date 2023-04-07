package weather

import (
	"fmt"
)

// Observer
type ForecastDisplay struct {
	currentPressure float64
	lastPressure    float64
	weatherData     Subject
}

func NewForecastDisplay(subject Subject) *ForecastDisplay {
	res := &ForecastDisplay{weatherData: subject}
	subject.RegisterObserver(res)
	return res
}

func (fd *ForecastDisplay) Update(temp, humidity, pressure float64) {
	fd.lastPressure = fd.currentPressure
	fd.currentPressure = pressure
	fd.display()
}

func (fd *ForecastDisplay) display() {
	fmt.Println("Forecast: ")
	if fd.currentPressure > fd.lastPressure {
		fmt.Println("Improving weather on the way!")
	} else if fd.currentPressure == fd.lastPressure {
		fmt.Println("More of the same")
	} else if fd.currentPressure < fd.lastPressure {
		fmt.Println("Watch out for cooler, rainy weather")
	}
}

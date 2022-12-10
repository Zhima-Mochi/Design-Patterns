package weather

import (
	"fmt"
	weatherInf "observer/interface/weather"
)

// Observer
type CurrentConditionDisplay struct {
	temperature float64
	humidity    float64
	weatherData weatherInf.Subject
}

func NewCurrentConditionDisplay(subject weatherInf.Subject) *CurrentConditionDisplay {
	res := &CurrentConditionDisplay{weatherData: subject}
	subject.RegisterObserver(res)
	return res
}

func (ccd *CurrentConditionDisplay) Update(temperature, humidity, pressure float64) {
	ccd.temperature = temperature
	ccd.humidity = humidity
	ccd.display()
}

func (ccd *CurrentConditionDisplay) display() {
	fmt.Println("Current conditions: ", ccd.temperature, "F degrees and ", ccd.humidity, "% humidity")
}

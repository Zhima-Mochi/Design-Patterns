package weather

import (
	"fmt"
	"math"
	weatherInf "observer/interface/weather"
)

// Observer
type StatisticsDisplay struct {
	maxTemp     float64
	minTemp     float64
	tempSum     float64
	numReadings int
	weatherData weatherInf.Subject
}

func NewStatisticsDisplay(subject weatherInf.Subject) *StatisticsDisplay {
	res := &StatisticsDisplay{maxTemp: math.SmallestNonzeroFloat64, minTemp: math.MaxFloat64, weatherData: subject}
	subject.RegisterObserver(res)
	return res
}

func (sd *StatisticsDisplay) Update(temp, humidity, pressure float64) {
	sd.tempSum += temp
	sd.numReadings++

	if temp > sd.maxTemp {
		sd.maxTemp = temp
	}

	if temp < sd.minTemp {
		sd.minTemp = temp
	}

	sd.display()
}

func (sd *StatisticsDisplay) display() {
	fmt.Println("Avg/Max/Min temperature =", (sd.tempSum / float64(sd.numReadings)), "/", sd.maxTemp, "/", sd.minTemp)
}

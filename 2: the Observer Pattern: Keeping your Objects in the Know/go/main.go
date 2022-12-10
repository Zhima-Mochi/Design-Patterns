package main

import (
	"fmt"
	"observer/concret/weather"
)

func main() {
	weatherData := weather.NewWeatherData()

	_ = weather.NewCurrentConditionDisplay(weatherData)
	_ = weather.NewStatisticsDisplay(weatherData)
	forecastDisplay := weather.NewForecastDisplay(weatherData)
	fmt.Println("=== test: SetMeasurements(80, 65, 30.4)")
	weatherData.SetMeasurements(80, 65, 30.4)
	fmt.Println("=== test: SetMeasurements(82, 70, 29.2)")
	weatherData.SetMeasurements(82, 70, 29.2)
	fmt.Println("=== test: SetMeasurements(78, 90, 29.2)")
	weatherData.SetMeasurements(78, 90, 29.2)

	fmt.Println("=== RemoveObserver: forecastDisplay")
	weatherData.RemoveObserver(forecastDisplay)
	fmt.Println("=== test: SetMeasurements(62, 90, 28.1)")
	weatherData.SetMeasurements(62, 90, 28.1)
}

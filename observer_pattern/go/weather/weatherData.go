package weather

// Subject
type WeatherData struct {
	observers   []Observer
	temperature float64
	humidity    float64
	pressure    float64
}

func NewWeatherData() *WeatherData {
	res := &WeatherData{observers: make([]Observer, 0)}
	return res
}

func (wd *WeatherData) RegisterObserver(o Observer) {
	wd.observers = append(wd.observers, o)
}

func (wd *WeatherData) RemoveObserver(o Observer) {
	newObservers := []Observer{}
	for _, observer := range wd.observers {
		if observer != o {
			newObservers = append(newObservers, observer)
		}
	}
	wd.observers = newObservers
}

func (wd *WeatherData) NotifyObservers() {
	for _, observer := range wd.observers {
		observer.Update(wd.temperature, wd.humidity, wd.pressure)
	}
}

func (wd *WeatherData) measurementsChanged() {
	wd.NotifyObservers()
}

func (wd *WeatherData) SetMeasurements(temperature, humidity, pressure float64) {
	wd.temperature = temperature
	wd.humidity = humidity
	wd.pressure = pressure
	wd.measurementsChanged()
}

func (wd *WeatherData) GetTemperature() float64 {
	return wd.temperature
}

func (wd *WeatherData) GetHumidity() float64 {
	return wd.humidity
}

func (wd *WeatherData) GetPressure() float64 {
	return wd.pressure
}

package weather

import (
	"fmt"
	weatherInf "observer/interface/weather"
)

// Observer
type HeatIndexDisplay struct {
	heatIndex   float64
	weatherData weatherInf.Subject
}

func NewHeatIndexDisplay(subject weatherInf.Subject) *HeatIndexDisplay {
	res := &HeatIndexDisplay{weatherData: subject}
	subject.RegisterObserver(res)
	return res
}

func (hid *HeatIndexDisplay) Update(t, rh, pressure float64) {
	hid.heatIndex = hid.computeHeatIndex(t, rh)
	hid.display()
}

func (hid *HeatIndexDisplay) computeHeatIndex(t, rh float64) float64 {
	index := ((16.923 + (0.185212 * t) + (5.37941 * rh) - (0.100254 * t * rh) + (0.00941695 * (t * t)) + (0.00728898 * (rh * rh)) + (0.000345372 * (t * t * rh)) - (0.000814971 * (t * rh * rh)) + (0.0000102102 * (t * t * rh * rh)) - (0.000038646 * (t * t * t)) + (0.0000291583 * (rh * rh * rh)) + (0.00000142721 * (t * t * t * rh)) + (0.000000197483 * (t * rh * rh * rh)) - (0.0000000218429 * (t * t * t * rh * rh)) + 0.000000000843296*(t*t*rh*rh*rh)) - (0.0000000000481975 * (t * t * t * rh * rh * rh)))
	return index
}

func (hid *HeatIndexDisplay) display() {
	fmt.Println("Heat index is ", hid.heatIndex)
}

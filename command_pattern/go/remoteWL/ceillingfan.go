package main

const (
	HIGH = iota
	MEDIUM
	LOW
	OFF
)

type CeilingFan struct {
	location string
	speed    int
}

func NewCeilingFan(location string) *CeilingFan {
	return &CeilingFan{
		location: location,
		speed:    OFF,
	}
}

func (cf *CeilingFan) High() {
	cf.speed = HIGH
}

func (cf *CeilingFan) Medium() {
	cf.speed = MEDIUM
}

func (cf *CeilingFan) Low() {
	cf.speed = LOW
}

func (cf *CeilingFan) Off() {
	cf.speed = OFF
}

func (cf *CeilingFan) GetSpeed() int {
	return cf.speed
}

func (cf *CeilingFan) String() string {
	return cf.location + " ceiling fan is on " + getSpeedString(cf.speed)
}

func getSpeedString(speed int) string {
	switch speed {
	case HIGH:
		return "high"
	case MEDIUM:
		return "medium"
	case LOW:
		return "low"
	case OFF:
		return "off"
	default:
		return "invalid speed"
	}
}

package main

type CeilingFanMediumCommand struct {
	ceilingFan *CeilingFan
	prevSpeed  int
}

func NewCeilingFanMediumCommand(ceilingFan *CeilingFan) *CeilingFanMediumCommand {
	return &CeilingFanMediumCommand{
		ceilingFan: ceilingFan,
	}
}

func (c *CeilingFanMediumCommand) Execute() {
	c.prevSpeed = c.ceilingFan.GetSpeed()
	c.ceilingFan.Medium()
}

func (c *CeilingFanMediumCommand) Undo() {
	switch c.prevSpeed {
	case HIGH:
		c.ceilingFan.High()
	case MEDIUM:
		c.ceilingFan.Medium()
	case LOW:
		c.ceilingFan.Low()
	case OFF:
		c.ceilingFan.Off()
	}
}

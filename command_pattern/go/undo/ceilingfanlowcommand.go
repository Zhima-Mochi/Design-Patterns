package main

type CeilingFanLowCommand struct {
	ceilingFan *CeilingFan
	prevSpeed  int
}

func NewCeilingFanLowCommand(ceilingFan *CeilingFan) *CeilingFanLowCommand {
	return &CeilingFanLowCommand{
		ceilingFan: ceilingFan,
	}
}

func (c *CeilingFanLowCommand) Execute() {
	c.prevSpeed = c.ceilingFan.GetSpeed()
	c.ceilingFan.Low()
}

func (c *CeilingFanLowCommand) Undo() {
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

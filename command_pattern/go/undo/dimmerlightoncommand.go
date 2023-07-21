package main

type DimmerLightOnCommand struct {
	light     *Light
	prevLevel int
}

func NewDimmerLightOnCommand(light *Light) *DimmerLightOnCommand {
	return &DimmerLightOnCommand{
		light: light,
	}
}

func (d *DimmerLightOnCommand) Execute() {
	d.prevLevel = d.light.GetLevel()
	d.light.Dim(75)
}

func (d *DimmerLightOnCommand) Undo() {
	d.light.Dim(d.prevLevel)
}

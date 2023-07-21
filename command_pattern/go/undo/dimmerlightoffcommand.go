package main

type DimmerLightOffCommand struct {
	light     *Light
	prevLevel int
}

func NewDimmerLightOffCommand(light *Light) *DimmerLightOffCommand {
	return &DimmerLightOffCommand{
		light:     light,
		prevLevel: 100,
	}
}

func (d *DimmerLightOffCommand) Execute() {
	d.prevLevel = d.light.GetLevel()
	d.light.Off()
}

func (d *DimmerLightOffCommand) Undo() {
	d.light.Dim(d.prevLevel)
}

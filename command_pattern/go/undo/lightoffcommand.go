package main

type LightOffCommand struct {
	light *Light
	level int
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{
		light: light,
	}
}

func (l *LightOffCommand) Execute() {
	l.level = l.light.GetLevel()
	l.light.Off()
}

func (l *LightOffCommand) Undo() {
	l.light.Dim(l.level)
}

package main

type LightOnCommand struct {
	light *Light
	level int
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{
		light: light,
	}
}

func (l *LightOnCommand) Execute() {
	l.level = l.light.GetLevel()
	l.light.On()
}

func (l *LightOnCommand) Undo() {
	l.light.Dim(l.level)
}

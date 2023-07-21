package main

import (
	"fmt"
	"reflect"
)

type RemoteControlWithUndo struct {
	onCommands  []Command
	offCommands []Command
	undoCommand Command
}

func NewRemoteControlWithUndo() *RemoteControlWithUndo {
	onCammads := make([]Command, 7)
	offCommands := make([]Command, 7)

	noCommand := NewNoCommand()
	for i := 0; i < 7; i++ {
		onCammads[i] = noCommand
		offCommands[i] = noCommand
	}
	undoCommand := noCommand

	return &RemoteControlWithUndo{
		onCommands:  onCammads,
		offCommands: offCommands,
		undoCommand: undoCommand,
	}
}

func (r *RemoteControlWithUndo) SetCommand(slot int, onCommand Command, offCommand Command) {
	r.onCommands[slot] = onCommand
	r.offCommands[slot] = offCommand
}

func (r *RemoteControlWithUndo) OnButtonWasPushed(slot int) {
	r.onCommands[slot].Execute()
	r.undoCommand = r.onCommands[slot]
}

func (r *RemoteControlWithUndo) OffButtonWasPushed(slot int) {
	r.offCommands[slot].Execute()
	r.undoCommand = r.offCommands[slot]
}

func (r *RemoteControlWithUndo) UndoButtonWasPushed() {
	r.undoCommand.Undo()
}

func (r *RemoteControlWithUndo) String() string {
	var str string
	str = "\n------ Remote Control -------\n"
	for i := 0; i < len(r.onCommands); i++ {
		str += fmt.Sprintf("[slot %d] %s\t\t%s\n", i, reflect.TypeOf(r.onCommands[i]).String(), reflect.TypeOf(r.offCommands[i]).String())
	}
	str += fmt.Sprintf("[undo] %s\n", reflect.TypeOf(r.undoCommand).String())
	return str
}

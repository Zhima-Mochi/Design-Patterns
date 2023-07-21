package main

import "fmt"

type RemoteControl struct {
	onCommands  []Command
	offCommands []Command
}

func NewRemoteControl() *RemoteControl {
	rc := &RemoteControl{}
	rc.onCommands = make([]Command, 7)
	rc.offCommands = make([]Command, 7)
	noCommand := func() {}
	for i := 0; i < 7; i++ {
		rc.onCommands[i] = noCommand
		rc.offCommands[i] = noCommand
	}
	return rc
}

func (rc *RemoteControl) SetCommand(slot int, onCommand, offCommand Command) {
	rc.onCommands[slot] = onCommand
	rc.offCommands[slot] = offCommand
}

func (rc *RemoteControl) OnButtonWasPushed(slot int) {
	rc.onCommands[slot]()
}

func (rc *RemoteControl) OffButtonWasPushed(slot int) {
	rc.offCommands[slot]()
}

func (rc *RemoteControl) String() string {
	var str string
	str += "\n------ Remote Control -------\n"
	for i := 0; i < len(rc.onCommands); i++ {
		str += fmt.Sprintf("[slot %d]", i)
	}
	return str
}

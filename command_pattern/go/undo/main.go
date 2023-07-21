package main

import "fmt"

func main() {
	remoteControl := NewRemoteControlWithUndo()

	livingRoomLight := NewLight("Living Room")

	livingRoomLightOn := NewLightOnCommand(livingRoomLight)
	livingRoomLightOff := NewLightOffCommand(livingRoomLight)

	remoteControl.SetCommand(0, livingRoomLightOn, livingRoomLightOff)

	remoteControl.OnButtonWasPushed(0)
	remoteControl.OffButtonWasPushed(0)
	fmt.Println(remoteControl)
	remoteControl.UndoButtonWasPushed()
	remoteControl.OffButtonWasPushed(0)
	remoteControl.OnButtonWasPushed(0)
	fmt.Println(remoteControl)
	remoteControl.UndoButtonWasPushed()

	ceilingFan := NewCeilingFan("Living Room")

	ceilingFanMedium := NewCeilingFanMediumCommand(ceilingFan)
	ceilingFanHigh := NewCeilingFanHighCommand(ceilingFan)
	ceilingFanOff := NewCeilingFanOffCommand(ceilingFan)

	remoteControl.SetCommand(0, ceilingFanMedium, ceilingFanOff)
	remoteControl.SetCommand(1, ceilingFanHigh, ceilingFanOff)

	remoteControl.OnButtonWasPushed(0)
	remoteControl.OffButtonWasPushed(0)
	fmt.Println(remoteControl)
	remoteControl.UndoButtonWasPushed()

	remoteControl.OnButtonWasPushed(1)
	fmt.Println(remoteControl)
	remoteControl.UndoButtonWasPushed()
}

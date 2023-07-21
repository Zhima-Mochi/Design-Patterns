package main

import "fmt"

func main() {
	remoteControl := NewRemoteControl()

	livingRoomLight := NewLight("Living Room")
	kitchenLight := NewLight("Kitchen")
	ceilingFan := NewCeilingFan("Living Room")
	garageDoor := NewGarageDoor("Main house")
	stereo := NewStereo("Living Room")

	remoteControl.SetCommand(0, livingRoomLight.On, livingRoomLight.Off)
	remoteControl.SetCommand(1, kitchenLight.On, kitchenLight.Off)
	remoteControl.SetCommand(2, ceilingFan.High, ceilingFan.Off)

	stereoOnWithCD := func() {
		stereo.On()
		stereo.SetCD()
		stereo.SetVolume(11)
	}
	remoteControl.SetCommand(3, stereoOnWithCD, stereo.Off)
	remoteControl.SetCommand(4, garageDoor.Up, garageDoor.Down)

	fmt.Println(remoteControl)

	remoteControl.OnButtonWasPushed(0)
	remoteControl.OffButtonWasPushed(0)
	remoteControl.OnButtonWasPushed(1)
	remoteControl.OffButtonWasPushed(1)
	remoteControl.OnButtonWasPushed(2)
	remoteControl.OffButtonWasPushed(2)
	remoteControl.OnButtonWasPushed(3)
	remoteControl.OffButtonWasPushed(3)
	remoteControl.OnButtonWasPushed(4)
	remoteControl.OffButtonWasPushed(4)
	remoteControl.OnButtonWasPushed(5)
}

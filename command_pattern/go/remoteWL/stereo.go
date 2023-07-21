package main

import "fmt"

type Stereo struct {
	location string
}

func NewStereo(location string) *Stereo {
	return &Stereo{
		location: location,
	}
}

func (s *Stereo) On() {
	println(s.location + " stereo is on")
}

func (s *Stereo) Off() {
	println(s.location + " stereo is off")
}

func (s *Stereo) SetCD() {
	println(s.location + " stereo is set for CD input")
}

func (s *Stereo) SetDVD() {
	println(s.location + " stereo is set for DVD input")
}

func (s *Stereo) SetRadio() {
	println(s.location + " stereo is set for Radio")
}

func (s *Stereo) SetVolume(volume int) {
	fmt.Printf("%s stereo volume set to %d\n", s.location, volume)
}

func (s *Stereo) String() string {
	return s.location + " stereo"
}

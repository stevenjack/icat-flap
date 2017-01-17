package main

type (
	// Sensor is a type of sensor used by the cat flap.
	Sensor struct {
		pinID uint
		Name  string
		State uint
	}
)

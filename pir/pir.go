package pir

import (
	"log"
	"time"
)

type (
	Sensor struct {
		PinID         uint
		ID            string
		DetectionTime time.Time
		State         uint
	}

	Movement struct {
		Inside       bool
		RFID         bool
		Outside      bool
		Completed    time.Time
		RFIDCount    int
		InitialStart string
	}
)

const INSIDE = "inside"
const OUTSIDE = "outside"
const RFID = "rfid"

const INSIDE_PIN = 20
const RFID_PIN = 16
const OUTSIDE_PIN = 21

const EXITED = 1
const ENTERED = 2
const IN_PROGRESS = 2

const LOW = 0
const HIGH = 1

func (s Sensor) ActivityDetected() bool {
	if s.State == HIGH {
		return true
	} else {
		return false
	}
}

func (s Sensor) Elapsed() float64 {
	elapsed := time.Since(s.DetectionTime)
	return elapsed.Seconds()
}

func (m *Movement) LogMovement(sensor Sensor) {
	log.Printf("[%s] Motion detected for %.2f seconds", sensor.ID, sensor.Elapsed())
	if sensor.Elapsed() > 6 {
		log.Printf("[%s] Logging movement as sufficient", sensor.ID)
		if sensor.ID == INSIDE {
			m.Inside = true
		} else {
			m.Outside = true
		}
	}
}

func (m *Movement) LogRFID(sensor Sensor) {
	m.RFIDCount++
	log.Printf("[%s] activity %d", sensor.ID, m.RFIDCount)
	if m.RFIDCount >= 2 {
		log.Printf("[%s] latch opened", sensor.ID)
		m.RFID = true
	}
}

func (m *Movement) Happened() bool {
	if m.RFID == true && m.Inside == true && m.Outside == true {
		return true
	} else {
		return false
	}
}

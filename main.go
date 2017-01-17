package main

import (
	"github.com/brian-armstrong/gpio"
	"github.com/stevenjack/icat-flap/pir"
	"log"
	"os"
	"time"
)

var sensors map[uint]pir.Sensor
var movement pir.Movement

func main() {
	watcher := gpio.NewWatcher()

	watcher.AddPin(pir.RFID_PIN)
	watcher.AddPin(pir.INSIDE_PIN)
	watcher.AddPin(pir.OUTSIDE_PIN)
	defer watcher.Close()

	sensors = make(map[uint]pir.Sensor)
	sensors[pir.RFID_PIN] = pir.Sensor{pir.RFID_PIN, pir.RFID, time.Now(), 0}
	sensors[pir.INSIDE_PIN] = pir.Sensor{pir.INSIDE_PIN, pir.INSIDE, time.Now(), 0}
	sensors[pir.OUTSIDE_PIN] = pir.Sensor{pir.OUTSIDE_PIN, pir.OUTSIDE, time.Now(), 0}

	movement = pir.Movement{
		false,
		false,
		false,
		time.Now(),
		0,
		"",
	}

	log.SetOutput(os.Stdout)

	log.Print("Starting monitoring...")
	for {
		pin, value := watcher.Watch()
		sensor := sensors[pin]
		sensor.State = value

		if sensor.ActivityDetected() {
			log.Printf("[%s] Activity detected", sensor.ID)
			sensor.DetectionTime = time.Now()
		} else {
			if sensor.ID == pir.RFID {
				movement.LogRFID(sensor)
			} else {
				movement.LogMovement(sensor)
			}
		}

		if movement.Happened() {
			log.Printf("--##### Movement detected, direction '%s' #####--", movement.InitialStart)
			movement = pir.Movement{
				false,
				false,
				false,
				time.Now(),
				0,
				"",
			}
		}

		if movement.InitialStart == "" && sensor.ID != pir.RFID {
			movement.InitialStart = sensor.ID
		}

		sensors[pin] = sensor
	}
}

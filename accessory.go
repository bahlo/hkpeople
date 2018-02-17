package hkpeople

import (
	"time"

	"github.com/bahlo/hkpeople/log"
	"github.com/bahlo/hkpeople/ping"

	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

// Accessory represents an occupancy sensor for people
type Accessory struct {
	*accessory.Accessory
	Sensor *service.OccupancySensor

	Threshold time.Duration
	Interval  time.Duration

	Targets []string

	stopChan    chan bool
	lastContact *time.Time
}

// NewAccessory creates a new Accessory with sane defaults
func NewAccessory(targets ...string) *Accessory {
	info := accessory.Info{
		Name:         "hkpeople",
		Manufacturer: "github.com/bahlo",
	}
	acc := &Accessory{
		Accessory: accessory.New(info, accessory.TypeSensor),
		Sensor:    service.NewOccupancySensor(),
		Interval:  3 * time.Second,
		Threshold: 15 * time.Minute,
		Targets:   targets,
		stopChan:  make(chan bool),
	}
	acc.AddService(acc.Sensor.Service)

	return acc
}

// Start begins to ping all targets
func (a *Accessory) Start() {
loop:
	for {
		err := ping.Any(a.Targets...)
		if err != nil {
			if a.lastContact == nil {
				log.Info.Printf("no contact yet, unoccupied")
				a.SetValue(false)
			} else if time.Since(*a.lastContact) > a.Threshold {
				log.Info.Printf("%s since last contact, unoccupied",
					time.Since(*a.lastContact))
				a.SetValue(false)
			} else {
				log.Info.Printf("%s since last contact, still occupied",
					time.Since(*a.lastContact))
			}
		} else {
			log.Info.Printf("Occupied")
			now := time.Now()
			a.lastContact = &now
			a.SetValue(true)
		}

		select {
		case <-a.stopChan:
			break loop
		case <-time.After(a.Interval):
		}
	}
}

// Stop stops the pinging
func (a *Accessory) Stop() chan bool {
	c := make(chan bool)

	go func() {
		log.Warn.Print("Waiting for ping to complete")
		a.stopChan <- true
		c <- true
	}()

	return c
}

// SetValue sets the value of the channel
func (a *Accessory) SetValue(value bool) {
	state := 0
	if value {
		state = 1
	}

	a.Sensor.OccupancyDetected.SetValue(state)
}

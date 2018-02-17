package main

import (
	"os"
	"strings"

	"github.com/bahlo/hkpeople"
	"github.com/bahlo/hkpeople/log"

	"github.com/brutella/hc"
)

// Pin is the numbers-representation of the letters for hkpeople
// on the numeric keyboard
const Pin = "45736753"

func main() {
	targets := strings.Split(os.Getenv("TARGETS"), ",")
	log.Info.Printf("Setting up accesory with %v", targets)
	s := hkpeople.NewAccessory(targets...)

	config := hc.Config{Pin: Pin, StoragePath: "./db"}
	t, err := hc.NewIPTransport(config, s.Accessory)
	if err != nil {
		panic(err)
	}

	hc.OnTermination(func() {
		log.Warn.Print("Shutting down")
		<-s.Stop()
		<-t.Stop()
	})

	go s.Start()
	t.Start()
}

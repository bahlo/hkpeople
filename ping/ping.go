package ping

import (
	"errors"
	"os/exec"
	"sync"

	"github.com/bahlo/hkpeople/log"
)

// One returns an err if a ping to the target was successful
func One(target string) error {
	return exec.Command("ping", "-t", "1", target).Run()
}

// Any returns nil if any of the targets did respond
func Any(targets ...string) error {
	wg := &sync.WaitGroup{}
	wg.Add(len(targets))

	ec := make(chan error)

	for _, target := range targets {
		go func(target string) {
			defer wg.Done()

			err := One(target)
			if err != nil {
				log.Debug.Printf("could not ping %s: %s", target, err)
			} else {
				log.Debug.Printf("reached %s", target)
			}
			ec <- err
		}(target)
	}

	go func() {
		wg.Wait()
		close(ec)
	}()

	for err := range ec {
		if err == nil {
			return nil
		}
	}

	return errors.New("All pings failed")
}

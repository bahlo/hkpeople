package ping

import "os/exec"

// Ping returns an err if a ping to the target was successful
func Ping(target string) error {
	return exec.Command("ping", "-c", "1", target).Run()
}

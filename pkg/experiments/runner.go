package experiments

import (
	"fmt"
	"time"
)

type Runner struct{}

func NewRunner() *Runner { return &Runner{} }

func (r *Runner) Run(name string, d time.Duration) error {
	switch name {
	case "network_latency":
		fmt.Printf("[chaos] injecting latency for %s\n", d)
	case "db_kill":
		fmt.Printf("[chaos] simulating database unavailable for %s\n", d)
	case "broker_outage":
		fmt.Printf("[chaos] broker API returning 503 for %s\n", d)
			default:
		return fmt.Errorf("unknown experiment: %s", name)
	}
	time.Sleep(d)
	return nil
}

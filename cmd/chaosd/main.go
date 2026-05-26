package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/xerife/chaos-trading-toolkit/pkg/experiments"
)

func main() {
	exp := flag.String("experiment", "network_latency", "experiment name")
	duration := flag.Duration("duration", 30*time.Second, "duration")
	flag.Parse()

	runner := experiments.NewRunner()
	log.Printf("starting chaos experiment=%s duration=%s", *exp, *duration)
	if err := runner.Run(*exp, *duration); err != nil {
		log.Fatal(err)
	}
	fmt.Println("experiment finished — see reports/")
}

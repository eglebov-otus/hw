package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		log.Fatalf("Failed to receive exact time. Reason: %s", err)

		return
	}

	fmt.Printf("current time: %s\n", time.Now())
	fmt.Printf("exact time: %s\n", ntpTime)
}

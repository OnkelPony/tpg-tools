package main

import (
	"battery"
	"fmt"
	"log"
)

func main() {
	status, err := battery.GetStatus()
	if err != nil {
		log.Fatalf("couldn't read battery status %v", err)
	}
	fmt.Printf("Battery %d%% charged\n", status.ChargePercent)
}

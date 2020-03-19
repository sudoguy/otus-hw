package main

import (
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	logger := log.New(os.Stdout, "", log.Lmsgprefix)

	currentTime := time.Now()
	ntpTime, err := ntp.Time("pool.ntp.org")

	if err != nil {
		log.Fatal(err)
	}

	logger.Println("current time:", currentTime)
	logger.Println("exact time:", ntpTime)
}

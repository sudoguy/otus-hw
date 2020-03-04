package main

import (
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	logger := log.New(os.Stdout, "", log.Lmsgprefix)

	exitCode := 0

	currentTime := time.Now()
	ntpTime, err := ntp.Time("pool.ntp.org")

	if err != nil {
		log.Println(err)
		exitCode = 1
	}

	logger.Println("current time:", currentTime)
	logger.Println("exact time:", ntpTime)

	os.Exit(exitCode)
}

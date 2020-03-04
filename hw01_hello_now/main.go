package main

import (
	"github.com/beevik/ntp"
	"log"
	"os"
	"time"
)


func main() {
	logger := log.New(os.Stdout, "", log.Lmsgprefix)

	currentTime := time.Now()
	ntpTime, err := ntp.Time("pool.ntp.org")

	if err != nil {
		log.Println(err)
		defer os.Exit(1)
	}

	logger.Println("current time:", currentTime)
	logger.Println("exact time:", ntpTime)
}

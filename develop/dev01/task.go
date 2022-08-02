package main

import (
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	var host string = "0.beevik-ntp.pool.ntp.org"

	tm, err := getTime(host)

	if err != nil {
		os.Exit(1)
	}

	log.Println(tm)
}

func getTime(host string) (time.Time, error) {
	tm, err := ntp.Time(host)
	if err != nil {
		log.Fatal(err.Error())
	}
	return tm, err
}

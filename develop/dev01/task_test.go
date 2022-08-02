package main

import (
	"testing"
	"time"
)

const host = "0.beevik-ntp.pool.ntp.org"

func TestTime(t *testing.T) {
	tm, err := getTime(host)

	if err != nil {
		t.Errorf("%s\n", err)
	}

	now := time.Now()

	t.Logf("Local time %v\n", tm)
	t.Logf("True time %v \n", now)
	t.Logf("Offset %v\n", tm.Sub(now))
}

func TestFailTime(t *testing.T) {
	_, err := getTime("bad host")
	if err != nil {
		t.Errorf("%s\n", err)
	}
}

package main

import (
	"fmt"
	"testing"
	"time"
)

func TestSetTime(t *testing.T) {

	cl := NewClock()
	exitChan := make(chan struct{})
	exitGoChan := make(chan bool)

	cl.hour = 10
	cl.min = 20
	cl.sec = 0
	RegisterClock(cl, exitGoChan, exitChan)

	time.Sleep(10 * time.Second)
	exitGoChan <- true

	<-exitChan

	fmt.Println("Exiting app ...")
}

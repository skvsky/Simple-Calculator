package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Clock struct {
	ticker *time.Ticker
	sec    int
	min    int
	hour   int
	year   int
	month  time.Month
	day    int
}

func parseTimeStr(cl *Clock, input string) error {

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	cldata := strings.FieldsFunc(input, f)

	if len(cldata) == 3 {
		hr, err := strconv.Atoi(cldata[0])
		if err != nil {
			log.Fatal(err)
		}
		cl.hour = hr
		min, err := strconv.Atoi(cldata[1])
		if err != nil {
			log.Fatal(err)
		}
		cl.min = min
		sec, err := strconv.Atoi(cldata[2])
		if err != nil {
			log.Fatal(err)
		}
		cl.sec = sec
	}

	return nil
}

func NewClock(input string) (*Clock, error) {

	cl := Clock{}

	if len(input) > 0 {
		err := parseTimeStr(&cl, input)
		if err != nil {
			log.Fatal(err)
			return &cl, err
		}
		return &cl, nil
	}

	nowTime := time.Now()

	cl.hour = nowTime.Hour()
	cl.min = nowTime.Minute()
	cl.sec = nowTime.Second()
	cl.hour = nowTime.Hour()
	cl.year = nowTime.Year()
	cl.month = nowTime.Month()
	cl.day = nowTime.Day()

	return &cl, nil
}

func (cl *Clock) displayMsg() {

	cl.sec = cl.sec + 1
	if cl.sec >= 60 {
		cl.sec = 0
		cl.min = cl.min + 1
		if cl.min >= 60 {
			cl.min = 0
			cl.hour = cl.hour + 1
			if cl.hour >= 24 {
				cl.hour = 0
				cl.PrintMsg("bong ... ")
			}
		} else {
			cl.PrintMsg("tock ...")
		}

	} else {
		cl.PrintMsg("tick ... ")
	}
}

func RegisterClock(cl *Clock, exitGoChan chan bool, exitChan chan struct{}) {

	cl.ticker = time.NewTicker(time.Second)

	go func() {
		for {
			select {
			case <-exitGoChan:
				cl.ticker.Stop()
				fmt.Println("Exiting go routinue")
				close(exitChan)
				return

			case <-cl.ticker.C:
				cl.displayMsg()
			}
		}
	}()
}

func (cl *Clock) PrintMsg(msg string) {
	str := fmt.Sprintf("%d:%d:%d ", cl.hour, cl.min, cl.sec)
	fmt.Println("tick ... ", str)

}

func main() {

	var nFlag string
	flag.StringVar(&nFlag, "inputTime", "09:00:00", "Input time")
	flag.Parse()

	cl, err := NewClock(nFlag)
	if err != nil {
		return
	}
	exitChan := make(chan struct{})
	exitGoChan := make(chan bool)
	RegisterClock(cl, exitGoChan, exitChan)

	//time.Sleep(10 * time.Second)
	//exitGoChan <- true

	<-exitChan

	fmt.Println("Exiting app ...")
}

package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		time.Sleep(1 * time.Second)
		go clock()
	}
}

func clock() {
	year, month, date, hour, minute, second := getCurrenttime()
	fmt.Println(year, "|", month, "|", date, "[", hour, ":", minute, ":", second, "]")
}

func getCurrenttime() (int, int, int, int, int, int) {
	currentTime := time.Now()
	timeStampString := currentTime.Format("2006-01-02 15:04:05")
	timeLayout := "2006-01-02 15:04:05"
	timeStamp, err := time.Parse(timeLayout, timeStampString)
	if err != nil {
		fmt.Println(err)
	}
	year, month, date := time.Now().Date()
	hour, minute, second := timeStamp.Clock()
	return year, int(month), date, hour, minute, second
}

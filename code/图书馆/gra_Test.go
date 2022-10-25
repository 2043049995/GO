package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(GetNextdayTime())
}

func p1() {
	fmt.Println("1")
}
func GetNextdayTime1() string {
	time1 := time.Now()
	timeday := time1.Format("2006-01-02")
	t2, _ := time.ParseInLocation("2006-01-02", timeday, time.Local)
	t1 := t2.AddDate(0, 0, +1)

	nextDayTime := t1.Format("2006-01-02")
	//fmt.Println(todayTime, nextDayTime)
	return nextDayTime
}

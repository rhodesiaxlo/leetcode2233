package main

import (
	"fmt"
	"time"
)

func main() {
	//timeFormate()

	//YearMonthDayHour()

	parseStringToTime()
}

// Q1
func timeFormate() {
	tn := time.Now()
	fmt.Println("use formate to formate time ", tn.Format("20060102"))
	fmt.Println("use formate to formate time ", tn.Format("20060102 15:04:05"))
	fmt.Println("use formate to formate time ", tn.Format("20060102 00:00:00"))
}

// Q2
func YearMonthDayHour() {
	a := time.Now()
	year := a.Year()
	month := a.Month()
	fmt.Println("year=", year, "month=", month)
}

func durationCase() {
	time.Sleep(time.Millisecond)
}

func parseStringToTime() {
	parse, _ := time.Parse("2006-01-02 15:04:05", "2022-12-21 01:02:03")
	fmt.Println(parse)

}

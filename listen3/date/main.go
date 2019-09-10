package main

import (
	"fmt"
	"time"
)

func testTime() {
	now := time.Now()
	fmt.Println(now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
}

func testTimestamp(timestamp int64) {

	timeObj := time.Unix(timestamp, 0)
	fmt.Println(timeObj)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	send := timeObj.Second()

	fmt.Printf("current timestamp:%d\n", timestamp)
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, send)

}
func processTask() {
	fmt.Printf("do task\n")
}

func testTicker() {
	ticker := time.Tick(1*time.Second)
	for i := range ticker {
		fmt.Printf("%v\n", i)
		processTask()
	}
}

func testConst() {
	fmt.Printf("nano second:%d\n", time.Nanosecond)
	fmt.Printf("micro second:%d\n", time.Microsecond)
	fmt.Printf("mili second:%d\n", time.Millisecond)
	fmt.Printf("second:%d\n", time.Second)
}

func testFormat() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Format("2006/1/02 15:04"))
}

func testCost1() {
	start := time.Now().UnixNano()
	fmt.Println(start)
	fmt.Println(time.Now().Unix())
	//for i := 0; i < 10; i++ {
	time.Sleep(10*time.Millisecond)
	//}
	end := time.Now().UnixNano()
	cost := (end - start)/1000
	fmt.Printf("code cost:%d us\n", cost)
}
func main() {
	//testTime()
	//timestamp := time.Now().Unix()
	//testTimestamp(timestamp)
	//testTicker()
	//testConst()
	//testFormat()
	testCost1()
}



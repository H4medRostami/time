//test package
package main

import (
	"fmt"
	"time"
)

func localize(timeStr time.Time, location string) time.Time {
	tl, _ := time.LoadLocation(location)
	now := timeStr.In(tl)
	return now
}

func main() {
	q := time.Now().UTC() // sample UTC 0
	temp := localize(q, "Asia/Istanbul")
	fmt.Println(q)
	fmt.Println(temp)
}

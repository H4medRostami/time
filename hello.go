package main

import (
	"fmt"
	"time"
)

func localize(x time.Time , location string ) time.Time {
	tl,_ :=time.LoadLocation(location)
	now :=x.In(tl)
	return now
}

func main(){
	q := time.Now().UTC()
	temp := localize(q,"Asia/Tehran")
	fmt.Println('\t',temp)
}



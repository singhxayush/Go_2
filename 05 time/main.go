package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	createDate := time.Date(2021, time.April, 13, 0, 0, 0, 0, time.Local)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}

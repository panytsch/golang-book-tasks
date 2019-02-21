package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	sleep(5)
	fmt.Println(time.Now())
}

func sleep(second int) {
	<-time.After(time.Second * time.Duration(second))
}

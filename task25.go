package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep
func sleep(i int) {
	<-time.After(time.Duration(i) * time.Second)
}

func main() {
	start := time.Now()
	fmt.Printf("main started %s\n", start.AppendFormat([]byte("time: "), time.Stamp))
	sleep(3)
	fmt.Printf("main finished since %s\n", time.Since(start))
}

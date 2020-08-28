package main

import "fmt"
import "time"

func main() {
	var now = time.Now()
	fmt.Printf("Current time: %02d:%02d:%02d\n", now.Hour(), now.Minute(), now.Second())
}

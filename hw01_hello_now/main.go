package main

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	// Запрашиваю время на сервере NTP
	netTime, _ := ntp.Time("0.ru.pool.ntp.org")
	// Получаю локальное время
	var now = time.Now()
	fmt.Printf("Current time: %02d:%02d:%02d\n", now.Hour(), now.Minute(), now.Second())
	fmt.Printf("Exact time: %02d:%02d:%02d\n", netTime.Hour(), netTime.Minute(), netTime.Second())
}

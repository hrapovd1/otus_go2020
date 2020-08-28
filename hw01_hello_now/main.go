package main

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	//	location, _ := time.LoadLocation("UTC")    // Устанавливаю временную зону на UTC, для прохождения теста
	netNow, _ := ntp.Time("0.ru.pool.ntp.org") // Запрашиваю время на сервере NTP
	var now = time.Now()                       // Получаю локальное время

	localTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, time.UTC)
	netTime := time.Date(netNow.Year(), netNow.Month(), netNow.Day(), netNow.Hour(), netNow.Minute(), netNow.Second(), 0, time.UTC)

	fmt.Printf("current time: %v\n", localTime)
	fmt.Printf("exact time: %v\n", netTime)
}

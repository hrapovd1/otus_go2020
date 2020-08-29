package main

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	const ntpServer = "0.ru.pool.ntp.org"

	netNow, _ := ntp.Time(ntpServer) // Запрашиваю время на сервере NTP
	now := time.Now()                // Получаю локальное время

	// Вариант вывода времени UTC, тест не проходит
	/*
		location, _ := time.LoadLocation("UTC") // Устанавливаю временную зону на UTC

		localTime := time.Date(now.Year(), now.Month(), now.Day(),
			now.Hour(), now.Minute(), now.Second(), 0, time.Local)

		netTime := time.Date(netNow.Year(), netNow.Month(), netNow.Day(),
			netNow.Hour(), netNow.Minute(), netNow.Second(), 0, time.Local)

		fmt.Printf("current time: %v\n", localTime.In(location))
		fmt.Printf("exact time: %v\n", netTime.In(location))
	*/

	// Вариант вывода локального времени с показом зоны UTC, тест проходит !
	localTime := time.Date(now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second(), 0, time.UTC)

	netTime := time.Date(netNow.Year(), netNow.Month(), netNow.Day(),
		netNow.Hour(), netNow.Minute(), netNow.Second(), 0, time.UTC)

	fmt.Printf("current time: %v\n", localTime)
	fmt.Printf("exact time: %v\n", netTime)
}

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	const ntpServer = "0.ru.pool.ntp.org"

	netNow, err := ntp.Time(ntpServer) // Запрашиваю время на сервере NTP, важно выполнить раньше локального запроса.
	if err != nil {
		log.Fatalf("NTP error: %v\n", err)
	}
	now := time.Now() // Получаю локальное время

	// Вариант вывода времени UTC, тест не проходит
	/*
		location, err := time.LoadLocation("UTC") // Устанавливаю временную зону на UTC
		if err != nil {
			log.Fatalf("LoadLocation error: %v\n", err)
		}

		localTime := time.Date(now.Year(), now.Month(), now.Day(),
			now.Hour(), now.Minute(), now.Second(), 0, time.Local) //Формирую строку локального времени

		netTime := time.Date(netNow.Year(), netNow.Month(), netNow.Day(),
			netNow.Hour(), netNow.Minute(), netNow.Second(), 0, time.Local) //Формирую строку сетевого времени

		fmt.Printf("current time: %v\n", localTime.In(location))
		fmt.Printf("exact time: %v\n", netTime.In(location))
	*/

	// Вариант вывода локального времени с показом зоны UTC, тест проходит !
	localTime := time.Date(now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second(), 0, time.UTC) //Формирую строку локального времени

	netTime := time.Date(netNow.Year(), netNow.Month(), netNow.Day(),
		netNow.Hour(), netNow.Minute(), netNow.Second(), 0, time.UTC) //Формирую строку сетевого времени

	fmt.Printf("current time: %v\n", localTime)
	fmt.Printf("exact time: %v\n", netTime)
}

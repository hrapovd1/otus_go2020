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
	localNow := time.Now() // Получаю локальное время

	fmt.Printf("current time: %v\n", localNow.UTC()) // Вывожу локальное время ПК в UTC
	fmt.Printf("exact time: %v\n", netNow.UTC())     // Вывожу сетевое время в UTC
}

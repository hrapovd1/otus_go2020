package main

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	netTime, _ := ntp.Time("0.ru.pool.ntp.org") // Запрашиваю время на сервере NTP
	var now = time.Now()                        // Получаю локальное время

	fmt.Printf("Current time: %02d:%02d:%02d\n", now.Hour(), now.Minute(), now.Second())
	fmt.Printf("Exact time: %02d:%02d:%02d\n", netTime.Hour(), netTime.Minute(), netTime.Second())
}

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func main() {
	t, err := getTime()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(5)
	}

	fmt.Println(t)
}

func getTime() (*time.Time, error) {
	options := ntp.QueryOptions{Timeout: 5 * time.Second}

	response, err := ntp.QueryWithOptions("0.beevik-ntp.pool.ntp.org", options)
	if err != nil {
		return nil, fmt.Errorf("Error do request to ntp: %s", err)
	}

	if response.Stratum == 0 {
		return nil, fmt.Errorf("Ntp response with stratum is zero. Kiss code: %s", response.KissCode)
	}

	time := time.Now().Add(response.ClockOffset)

	return &time, nil
}

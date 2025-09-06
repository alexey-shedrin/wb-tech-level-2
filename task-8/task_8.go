package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

var ntpServers = []string{
	"pool.ntp.org",
	"time.google.com",
	"time.windows.com",
	"time.apple.com",
	"ntp.ubuntu.com",
}

func getTime(servers []string) (time.Time, error) {
	var lastErr error
	for _, server := range servers {
		time, err := ntp.Time(server)
		if err == nil {
			return time, nil
		}
		lastErr = err
	}
	return time.Time{}, fmt.Errorf("Не удалось получить время ни с одного из NTP-серверов: %w", lastErr)
}

func main() {
	time, err := getTime(ntpServers)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Критическая ошибка: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(time)
}

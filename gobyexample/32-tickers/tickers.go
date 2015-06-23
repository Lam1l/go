// tickers.go
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Microsecond * 500)
	
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	
	time.Sleep(time.Microsecond * 5500)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

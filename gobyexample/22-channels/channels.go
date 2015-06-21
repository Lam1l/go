// channels.go
package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	
	go func() {
		for i := 0; i < 10000000000; i++ {
			if i%100000000 == 0 {
				messages <- "ping"
			}
		}
		messages <- "exit"
	}()
	
	for {
		msg := <- messages
		if msg != "exit" {
			fmt.Println(msg)
		} else {
			break
		}
	}
}

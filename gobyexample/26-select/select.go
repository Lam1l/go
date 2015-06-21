// select.go
package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	rand.NewSource(time.Now().UnixNano())
	fmt.Println(rand.Int())
	
	go func() {
		i := rand.Int() % 60
		fmt.Println("one delay:", i)
		time.Sleep(time.Second * time.Duration(i))
		c1 <- "one"
	}()
	
	go func() {
		i := rand.Int() % 60
		fmt.Println("two delay:", i)
		time.Sleep(time.Second * time.Duration(i))
		c2 <- "two"
	}()
	
	for i := 0; i < 2; i++ {
		select {
			case msg1 := <- c1:
				fmt.Println("received", msg1)
			case msg2 := <- c2:
				fmt.Println("received", msg2)
		}
	}
}

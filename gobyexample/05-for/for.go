package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		 i = i + 1
	}
	
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	
	for {
		fmt.Println("Exiting endless loop in...", i)
		i = i - 1
		if i < 1 {
			break
		}
	}
	fmt.Println("Exited successfully")
}
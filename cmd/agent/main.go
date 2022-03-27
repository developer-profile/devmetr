package main

import (
	"fmt"
	"time"
)

func main() {

	GetMetrics(2)

}

func GetMetrics(duration time.Duration) {
	var interval = duration * time.Second
	s := 0  // steps counter
	cs := 0 // send to server counter
	i := 0  // for increment base value
	for {
		<-time.After(interval)

		s += 1
		if i < 5 {
			i += 1
			//s := fmt.Sprintf("%s is %d years old.\n", name, age)
			fmt.Printf("Step #%d collecting data with 2 seconds interval %d \n", s, i)
		} else {
			i = 1
			cs += 1
			fmt.Printf("Step #%d sending data to 127.0.0.1:8080/update: #%d \n", s, cs)
			fmt.Printf("Step #%d collecting data with 2 seconds interval %d \n", s, i)

		}
		if s > 5 {
			fmt.Printf("Total steps: %d \nTotal server update: %d \nExiting.. \n", s, cs)
			return
		}

	}
}

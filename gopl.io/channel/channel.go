package main

import (
	"fmt"
	//	"time"
)

func printCount(c chan int) { // an int type channel is passed
	num := 0
	for num >= 0 {
		num = <-c
		fmt.Print(num, " ")
	}
}

func main() {
	c := make(chan int)

	a := []int{1, 8, 6, 7, 5, 3, 0, 9, 1, 6, 7}

	go printCount(c)

	for _, v := range a {
		c <- v
	}

	fmt.Println("End of main()")

}

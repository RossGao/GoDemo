package main

import (
	channel "Xiang/Concurrency/Channel"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go PrintNumber(i) // Control will give back to the calling mehtod immediatly before printing is done.
	}
	fmt.Println("I am just to show I am here at the moment.")
	var input string
	fmt.Scanln(&input)
}

func PrintNumber(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		sleep := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * sleep)
	}
	fmt.Println("Hi I am done printing.")
	channel.Channel()
}

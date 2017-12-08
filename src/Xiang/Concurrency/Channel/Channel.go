package channel

import (
	"fmt"
	"time"
)

func Channel() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "This is channel 1 brick"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "This is channel 2 brick"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case brick1 := <-c1: // Check which channel contians content.
				fmt.Println(brick1)
			case brick2 := <-c2:
				fmt.Println(brick2)
			}
		}
	}()

	input := ""
	fmt.Scanln(&input)
}

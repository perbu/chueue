package main

import (
	"fmt"
	"github.com/perbu/chueue/chueue"
	"time"
)

func main() {
	c := chueue.NewChueue()
	go consumer1(c)
	go consumer2(c)

	producer(c)
}

func producer(c *chueue.Chueue) {
	for i := 0; i < 10; i++ {
		c.Release()
		time.Sleep(time.Second)
	}
}

func consumer2(c *chueue.Chueue) {
	for {
		c.Wait()
		fmt.Println("consumer2 got released")
	}
}

func consumer1(c *chueue.Chueue) {
	for {
		c.Wait()
		fmt.Println("consumer1 got released")
	}
}

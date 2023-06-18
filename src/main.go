package main

import (
	"os"
	c "task1/src/consumer"
	p "task1/src/producer"
)

func main() {
	go func() {
		p.StartProducer()
	}()
	i := c.StartConsumer()
	os.Exit(i)
}

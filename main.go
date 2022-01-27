package main

import (
	"fmt"
	"sync"
	"time"
)

var messages = [][]string{
	{
		"The fourth office is that of the deacons, who have to do with 1 ",
		" a calf-skin filled with straw, supposed to induce the cow to give milk freely;",
		"hence a term of contempt for one who is used as a dummy for the advantage of another.",
		"just one big hoax.",
	},
	{
		"The fourth office is that of the deacons, who have to do with 1 ",
		" a calf-skin filled with straw, supposed to induce the cow to give milk freely;",
		"hence a term of contempt for one who is used as a dummy for the advantage of another.",
		"just one big hoax.",
	},
	{
		"The fourth office is that of the deacons, who have to do with 1 ",
		" a calf-skin filled with straw, supposed to induce the cow to give milk freely;",
		"hence a term of contempt for one who is used as a dummy for the advantage of another.",
		"just one big hoax.",
	},
}

func produce(link chan<- []string, msg []string, wg *sync.WaitGroup) {
	defer wg.Done()
	link <- msg
}

const producerCount int = 4
const consumerCount int = 3

func consume(link <-chan []string, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range link {
		fmt.Printf("Message \"%v\" is consumed by consumer %v at time %v\n", msg[0], id, time.Now())
	}
}

func main() {

	link := make(chan []string)

	wp := &sync.WaitGroup{}
	wc := &sync.WaitGroup{}

	wp.Add(producerCount)
	wc.Add(consumerCount)

	for i := 0; i < producerCount; i++ {
		go produce(link, messages[i], wp)
	}

	for i := 0; i < consumerCount; i++ {
		go consume(link, i, wc)
	}

	wp.Wait()
	close(link)
	wc.Wait()
}

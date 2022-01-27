package main

import (
	"fmt"
	"sync"
	"time"
)

var messages = [][]string{
	{
		"The world itself's",
		"just one big hoax.",
		"Spamming each other with our",
		"running commentary of bullshit,",
	},
	{
		"but with our things, our property, our money.",
		"I'm not saying anything new.",
		"We all know why we do this,",
		"not because Hunger Games",
		"books make us happy,",
	},
	{
		"masquerading as insight, our social media",
		"faking as intimacy.",
		"Or is it that we voted for this?",
		"Not with our rigged elections,",
	},
	{
		"but because we wanna be sedated.",
		"Because it's painful not to pretend,",
		"because we're cowards.",
		"- Elliot Alderson",
		"Mr. Robot",
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

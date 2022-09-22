package main

import (
	"fmt"
	"sync"
)


var wg = sync.WaitGroup{}

func main() {
	channel := make(chan string)
	wg.Add(2)

	go sortAlphabet(channel)
	go printAlphabet(channel)

	wg.Wait()
}

func sortAlphabet(ch chan string) {
	for l := byte('a'); l <= byte('z'); l  += 1 {
		ch <- string(l)
	}
	wg.Done()
}

func printAlphabet(ch chan string ) {

	for l := byte('a'); l <= byte('z'); l  += 1 {
		fmt.Println(<- ch)
	}
	wg.Done()
}
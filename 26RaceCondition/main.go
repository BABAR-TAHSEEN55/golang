package main

import (
	"fmt"
	"sync"
)

// NOTE : Mutex are used to lock in the file so that not more than one go routine modifies it at the same time
func main() {
	var score = []int{0}
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	wg.Add(3) //Total number of Go Routines
	fmt.Println("Race Condition in Golang")
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("This is First")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("This is Second")

		mut.Lock()
		score = append(score, 1)
		mut.Unlock()

		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("This is Third")

		mut.Lock()
		score = append(score, 1)
		mut.Unlock()

		wg.Done()
	}(wg, mut)
	wg.Wait()
	fmt.Println("score  :  ", score)

}

//NOTE: : When go routines are fired , it is important to wait to print

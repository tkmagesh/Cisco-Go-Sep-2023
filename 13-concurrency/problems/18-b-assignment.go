/*
Modify the below program to generate prime numbers using 3 instances of genPrimes functions and print the generated values
	genPrimes(3,100)
	genPrimes(101,200)
	genPrimes(201,300)
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := genPrimes(3, 100)
	ch2 := genPrimes(101, 200)
	ch3 := genPrimes(201, 300)

	primeNosCh := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		go func() {
			for primeNo := range ch1 {
				primeNosCh <- primeNo
			}
			wg.Done()
		}()

		go func() {
			for primeNo := range ch2 {
				primeNosCh <- primeNo
			}
			wg.Done()
		}()

		go func() {
			for primeNo := range ch3 {
				primeNosCh <- primeNo
			}
			wg.Done()
		}()
	}()

	doneCh := make(chan struct{})
	go func() {
		for primeNo := range primeNosCh {
			fmt.Println(primeNo)
		}
		close(doneCh)
	}()

	wg.Wait()
	close(primeNosCh)
	<-doneCh
	fmt.Println("Done!")
}

func genPrimes(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			if isPrime(no) {
				ch <- no
				time.Sleep(500 * time.Millisecond)
			}
		}
		close(ch)
	}()
	return ch
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

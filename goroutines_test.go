package simplepackage

import (
	"fmt"
	"testing"
	"time"
)

func TestGoRoutines(t *testing.T) {
	sayFiveTimesGone()
	goChannelSum()
	bufferedChannel()
	channelledFibonacci()
	channelledFibonacciWithSelect()
}

/* start of go routine */
func sayFiveTimes(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sayFiveTimesGone() {
	go sayFiveTimes("world")
	sayFiveTimes("hello")
}

/* start of channelled sum */
func summation(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func goChannelSum() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go summation(s[:len(s)/2], c)
	go summation(s[len(s)/2:], c)
	x, y := <-c, <-c //receive from c
	fmt.Println(x, y, x+y)
}

/* start of buffered channel */
func bufferedChannel() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/* start of buffered channel fibonacci */
func fibbonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, y+x
	}
	close(c)
}

func channelledFibonacci() {
	c := make(chan int, 10)
	go fibbonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

/* start of fibonacci with select */
func fibonacciWithSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func channelledFibonacciWithSelect() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciWithSelect(c, quit)
}

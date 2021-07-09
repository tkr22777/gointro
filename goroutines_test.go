package simplepackage

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

func TestGoRoutines(t *testing.T) {
	sayFiveTimesGone()
	goChannelSum()
	bufferedChannel()
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
	//<-c means receive from c, this is a blocking call, i.e if there is nothing on c, this call will block until then
	x, y := <-c, <-c
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

func TestChannelledFibonacci(t *testing.T) {
	channelledFibonacci()
}

/* fibonacci series is defined as 0, 1, 1 (0+1), 2 (1+1), 3 (1+2), 5 (2+3), ..... */
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
	i := 0
	for val := range c {
		fmt.Printf("Channelled Fibo[%d]: %d\n", i, val)
		i = i + 1
	}
}

func TestChannedlledFibonacciWithSelect(t *testing.T) {
	channelledFibonacciWithSelect()
}

/* start of fibonacci with select */
func fibonacciWithSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			time.Sleep(time.Millisecond)
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
		/* after printing for 10 times sending signal to quit */
		for i := 0; i < 10; i++ {
			time.Sleep(2 * time.Millisecond)
			fmt.Printf("Channelled Fibo With Select[%d]: %d\n", i, <-c)
		}
		quit <- 0
	}()
	fibonacciWithSelect(c, quit)
	fmt.Println("done!")
}

func TestErrGroup(t *testing.T) {
	g, _ := errgroup.WithContext(context.TODO())
	testArray := []int64{0, 1, 2, 3, 4, 5, 6, 7}
	newArray := make([]int64, 0, len(testArray))
	for _, val := range testArray {
		newVal := val
		g.Go(func() error {
			newArray = append(newArray, newVal)
			fmt.Println(newArray)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Printf("some error")
	}
	fmt.Printf("success")
}

func TestErrGroupBatch(t *testing.T) {
	g, _ := errgroup.WithContext(context.TODO())
	nums := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	squaredNums := make([]int64, 0, len(nums))
	//numsByNumString := make(map[int64]string, len(nums))
	numsForRead := make(map[int64]string, len(nums))
	batchSize := 3
	smallBatch := make([]int64, 0, batchSize)
	for _, val := range nums {
		numsForRead[val] = strconv.Itoa(int(val))
	}
	for i, val := range nums {
		smallBatch = append(smallBatch, val)
		if (i+1)%batchSize == 0 || (i+1) == len(nums) {
			procBatch := smallBatch
			smallBatch = make([]int64, 0, batchSize)
			g.Go(func() error {
				for _, v := range procBatch {
					//numsByNumString[v] = strconv.Itoa(int(v)) // would panic with concurrent write
					squaredNums = append(squaredNums, v)
					fmt.Printf("Map read: %s\n", numsForRead[v])
				}
				return fmt.Errorf("some error")
			})
		}

	}
	if err := g.Wait(); err != nil {
		fmt.Printf("some error")
	}
	fmt.Println(squaredNums)
}

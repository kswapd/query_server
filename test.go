package main

import (
	"fmt"
	"time"
)

func TestGo(ch chan int) int {
	out := <-ch
	ch <- out + 1
	fmt.Println("TestGo", out)
	return 0
}
func TestChan(n int) int {
	c := make(chan int, 10)

	//go func() {
	c <- 48
	c <- 96
	//time.Sleep(2 * time.Second)
	c <- 200
	//}()

	time.Sleep(1 * time.Second)
	for v := range c {
		fmt.Println(v)
	}

	// 保持持续运行
	//holdRun()
	return 0
}
func main2() {
	var i int
	i = 5
	var j = 66
	ch := make(chan int, 1)
	ch <- i
	fmt.Println("Hello, world.", i, j)
	for m := 0; m < 5; m++ {
		go TestGo(ch)
	}

	TestChan(3)
	time.Sleep(time.Second * 3)

}

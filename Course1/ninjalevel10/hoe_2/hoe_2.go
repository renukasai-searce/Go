package main

import "fmt"

func main() {
	cr := make(chan int)
	go func() {
		cr <- 23
	}()
	fmt.Println(<-cr)
	fmt.Printf("cr\t%T\n", cr)
}

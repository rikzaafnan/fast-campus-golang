package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	// fmt.Println("halo dari main 1")
	// go sayHello()
	// fmt.Println("halo dari main 2")
	//
	// time.Sleep(time.Second * 5)
	// fmt.Println("halo dari main 3")
	go printNumbers(1, &wg)
	wg.Add(1)
	go printNumbers(2, &wg)
	wg.Add(1)
	go printNumbers(3, &wg)
	wg.Add(1)
	go printNumbers(4, &wg)
	wg.Add(1)
	go printNumbers(5, &wg)

	wg.Wait()
	fmt.Println("program selesai")
}

// func sayHello() {
// 	fmt.Println("halo, dari goroutine")
// }

func printNumbers(jobID int, wg *sync.WaitGroup) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("job id: %d, menjalankan task %d\n", jobID, i)
	}
	wg.Done()
}

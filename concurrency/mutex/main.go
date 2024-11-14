package main

import (
	"fmt"
	"runtime"
	"sync"
)

var mutex sync.Mutex
var stock int = 100

func checkout(seq int, quantity int) {

	mutex.Lock()

	if stock >= quantity {
		stock -= quantity
		fmt.Printf("pembelian ke: %d berhasil dengan jumlah: %d\n", seq, quantity)

	} else {
		fmt.Printf("pembelian ke: %d gagal, karena stock tidak cukup\n", seq)
	}

	mutex.Unlock()

}

func main() {
	// for i := 0; i < 10; i++ {
	// 	go checkout(i, 20)
	// }

	// time.Sleep(3 * time.Second)

	test2()
}

type counter struct {
	val int
}

func (c *counter) Add(int) {
	c.val++
}

func (c *counter) Value() int {
	return c.val
}

func test2() {
	runtime.GOMAXPROCS(4)

	var wg sync.WaitGroup
	var meter counter
	var sync sync.Mutex

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 100; j++ {
				sync.Lock()
				meter.Add(1)
				sync.Unlock()
			}

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(meter.Value())
}

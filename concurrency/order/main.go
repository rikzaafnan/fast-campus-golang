package main

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	ID      int
	IsValid bool
}

func main() {

	// handler order
	order := Order{
		ID: 1,
	}
	// validate
	// distribute validation task => fanOut CH
	fanOutCH := make(chan func(), 4)

	// wait all validation task => fanIn CH
	fanInCH := make(chan bool, 4)

	var wg sync.WaitGroup

	wg.Add(4)
	// fanOut : distribute validate task
	go func() {
		fanOutCH <- func() {
			validatePaymentStatus(&order, &wg, fanInCH)
		}

		fanOutCH <- func() {
			validateSellerState(&order, &wg, fanInCH)
		}

		fanOutCH <- func() {
			validateStock(&order, &wg, fanInCH)
		}

		fanOutCH <- func() {
			validateShippingAddress(&order, &wg, fanInCH)
		}

		close(fanOutCH)

	}()

	// fanIn : collect result from all validate task
	go func() {
		wg.Wait()
		close(fanInCH)
	}()

	// 	start worker
	for f := range fanOutCH {
		go f()
	}

	order.IsValid = true
	for result := range fanInCH {
		if !result {
			order.IsValid = false

		}
	}

	if !order.IsValid {
		fmt.Println("order is not valid, can't be processed")
	} else {
		fmt.Println("order is valid")
	}

}

func validatePaymentStatus(order *Order, wg *sync.WaitGroup, fanInCH chan<- bool) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("payment status validated : passed")
	fanInCH <- true
	defer wg.Done()
}

func validateSellerState(order *Order, wg *sync.WaitGroup, fanInCH chan<- bool) {
	time.Sleep(600 * time.Millisecond)
	fmt.Println("seller state validated : passed")
	fanInCH <- true
	defer wg.Done()
}

func validateStock(order *Order, wg *sync.WaitGroup, fanInCH chan<- bool) {
	time.Sleep(400 * time.Millisecond)
	fmt.Println("stock validated : passed")
	fanInCH <- true
	defer wg.Done()
}

func validateShippingAddress(order *Order, wg *sync.WaitGroup, fanInCH chan<- bool) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("shipping address validated : passed")
	fanInCH <- true
	defer wg.Done()
}

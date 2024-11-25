package main

import (
	"context"
	"fmt"
	"time"
)

// example, data driver have data cached on redis
// var (
// 	TopDriverIDs      []uuid.UUID
// 	ReliableDriverIDs []uuid.UUID
// 	NormalDriverIDs   []uuid.UUID
// )

func main() {
	fmt.Println("start...")
	FindNearbyDriverIDs()
	fmt.Println("finished...")
}

func FindNearbyDriverIDs() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	doneCH := make(chan struct{}, 3)

	// findTopDriverIDs()      -- optional
	go func() {
		findTopDriverIDs(ctx)
		doneCH <- struct{}{}

	}()
	// findReliableDriverIDs() -- optional
	go func() {
		findReliableDriverIDs(ctx)
		doneCH <- struct{}{}
	}()
	// findNormalDriverIDs()   -- mandatory
	go func() {
		findNormalDriverIDs(ctx)
		doneCH <- struct{}{}
	}()

	for i := 0; i < 3; i++ {
		<-doneCH

	}

	fmt.Println("all goroutine finished")

}

func findTopDriverIDs(ctx context.Context) (driverIDs []int) {
	// simulate to have a long running query
	select {
	case <-time.After(1200 * time.Millisecond):
		fmt.Println("top driver found")
		return []int{123123, 21312}
	case <-ctx.Done():
		fmt.Println("find top driver cancelled")
		return []int{}
	}

}

func findReliableDriverIDs(ctx context.Context) (driverIDs []int) {
	// simulate to have a long running query
	select {
	case <-time.After(2000 * time.Millisecond):
		fmt.Println("reliable driver found")
		return []int{222, 333}
	case <-ctx.Done():
		fmt.Println("find reliable driver cancelled")
		return []int{}
	}
}

func findNormalDriverIDs(ctx context.Context) (driverIDs []int) {
	fmt.Println("normal driver found")
	return []int{44, 55}
}

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// interruptionCH
	// sigCH

	sigCH := make(chan os.Signal, 1)
	signal.Notify(sigCH, syscall.SIGTERM, syscall.SIGINT)
	interruptionCH := make(chan bool)

	go func() {
		for {
			select {
			case <-interruptionCH:
				fmt.Printf("tTask goroutine diberhentikan ....")
				return
			default:

				fmt.Printf("menjalankan tugas ...\n")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	select {
	case interruptionSignal := <-sigCH:
		fmt.Printf("interruption signal triggered: %v\n", interruptionSignal)
		fmt.Println("shutting down the program gracefully...")
		close(interruptionCH)
	}

}

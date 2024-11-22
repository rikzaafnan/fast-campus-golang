package main

import (
	"fmt"
	"time"
)

func main() {
	// 	tasks channel
	// 	dispatcher
	// 	worker
	// 	live worker

	taskQueueCh := make(chan Notification, 10)
	dispatcher := Dispatcher{
		TaskQueue:  taskQueueCh,
		MaxWorkers: 3,
	}

	dispatcher.StartWorkers()

	go func() {
		for i := 1; ; i++ {
			notification := Notification{
				ID:      i,
				Message: fmt.Sprintf("message ke- %d", i),
			}

			dispatcher.AssignTask(notification)

			time.Sleep(500 * time.Millisecond)

		}
	}()

	select {}
}

type Notification struct {
	ID      int
	Message string
}

func (n *Notification) Send() {
	fmt.Printf("send message dengan content : %s \n", n.Message)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("sending message is successfully")
}

type Worker struct {
	ID        int
	TaskQueue chan Notification
}

func (w *Worker) Start() {
	go func() {
		for notification := range w.TaskQueue {
			fmt.Printf("worker with ID %d sending notif \n", notification.ID)
			notification.Send()
			fmt.Printf("worker with ID %d success send notif \n", notification.ID)

		}
	}()
}

type Dispatcher struct {
	TaskQueue  chan Notification
	MaxWorkers int
}

func (d *Dispatcher) StartWorkers() {
	// 	dispatch task
	for i := 1; i <= d.MaxWorkers; i++ {
		worker := Worker{
			ID:        i,
			TaskQueue: d.TaskQueue,
		}
		worker.Start()
	}
}

func (d *Dispatcher) AssignTask(notification Notification) {
	d.TaskQueue <- notification
}

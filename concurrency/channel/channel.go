package channel

import (
	"fmt"
	"time"
)

func Demo() {
	fmt.Println("program dijalankan")

	// 	flow :
	// 	- buatkan channel --> done
	// 	- buatkan simulasi penerimaan data lewat channel
	// 	- buatkan simulasi pengiriman data lewat channel
	// 	- tunggu sampai program selesai

	messageCh := make(chan string, 3)

	// receiver
	go func() {

		for {
			messageData := <-messageCh
			fmt.Printf("data di terima : %s\n", messageData)
		}

	}()

	// sender
	go func() {

		for i := 0; i <= 12; i++ {
			fmt.Printf("data ke %d dikirim\n", i)
			messageCh <- fmt.Sprintf("ini ada pesan dari goroutine ke:%d", i)
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("program selesai")
}

func Pembagian() {
	fmt.Println("aaaaa")
}

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

	messageCh := make(chan string)

	// receiver
	go func() {

		for i := 1; i <= 5; i++ {
			fmt.Println("menerima pesan")
			messageData := <-messageCh
			fmt.Printf("pesan di terimaa %s\n", messageData)
		}

	}()

	// sender
	go func() {

		fmt.Println("mengirim pesan")
		messageCh <- "ini ada pesan dari goroutine"
		fmt.Println("Pesan terkirim")
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("program selesai")
}

func Pembagian() {
	fmt.Println("aaaaa")
}

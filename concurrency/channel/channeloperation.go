package channel

import (
	"fmt"
	"math/rand"
	"time"
)

func DemoChannel() {

	priceCh := make(chan StockPrice, 100)

	go priceProcessor(priceCh)

	// start fetching stock prices for multiple symbol
	go fetchStockPrices(priceCh, "AAPL")
	go fetchStockPrices(priceCh, "GOOGL")
	go fetchStockPrices(priceCh, "MSFT")

	// 	simulate some delay to allow fetching and processing
	time.Sleep(10 * time.Second)
}

type StockPrice struct {
	Symbol string
	Price  float64
	Time   time.Time
}

func priceProcessor(priceCh <-chan StockPrice) {
	for price := range priceCh {
		fmt.Printf("Processing stock price: %s = %2.f at %s\n", price.Symbol, price.Price, price.Time)

		// 	simulate processing time
		time.Sleep(500 * time.Millisecond)
	}
}

// fetchStockPrices will be called by a scheduler
// this channel is to receive only
func fetchStockPrices(priceCh chan<- StockPrice, symbol string) {
	for {
		price := StockPrice{
			Symbol: symbol,
			Price:  rand.Float64() * 100,
			Time:   time.Now(),
		}
		priceCh <- price

		// 	simulate delay between price updates
		time.Sleep(time.Second)

	}
}

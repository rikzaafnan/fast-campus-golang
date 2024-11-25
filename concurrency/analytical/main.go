package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 	simulate fetching order
	// simulate filtering order
	// simulate analysing order
	// simulate storing the analysis result to big query
	// 	flow::>
	// 	fetch -> filter ->analyze -> store
	orderCH := make(chan Order)
	filterOrderCH := make(chan Item)
	analysisReportCH := make(chan AnalysisReport)

	// run stage of pipeline concurrency
	go fetchOrder(orderCH)
	go filterOrder(orderCH, filterOrderCH)
	go analyseOrder(filterOrderCH, analysisReportCH)
	go storeAnalysisReport(analysisReportCH)

	select {}
}

type (
	Order struct {
		ID    int
		Items []Item
	}

	Item struct {
		ItemID       int
		Category     string // to be filtered, specific for "digital" category
		ProviderName string // telkomsel, indosat,esia
		Price        float64
	}

	AnalysisReport struct {
		Category     string
		AveragePrice float64
		MinPrice     float64
		MaxPrice     float64
	}
)

func fetchOrder(orderCh chan<- Order) {

	for i := 1; ; i++ {
		orderData := Order{
			ID: i,
			Items: []Item{
				{
					ItemID:       rand.Int(),
					Category:     "digital",
					ProviderName: "telkomsel",
					Price:        rand.Float64() * 1000,
				},
			},
		}

		time.Sleep(500 * time.Millisecond)
		orderCh <- orderData
		fmt.Println("fetched the order")
	}

}
func filterOrder(orderCh <-chan Order, filterOrderCH chan<- Item) {
	for order := range orderCh {
		for _, item := range order.Items {
			if item.Category == "digital" {
				time.Sleep(200 * time.Millisecond)
				filterOrderCH <- item
				fmt.Println("filter the order detail")
			}
		}
	}
}

func analyseOrder(filterOrderCH <-chan Item, analysisReport chan<- AnalysisReport) {
	for order := range filterOrderCH {
		// do analysis here
		result := AnalysisReport{
			Category:     order.Category,
			MaxPrice:     order.Price * 2,
			MinPrice:     order.Price,
			AveragePrice: order.Price,
		}

		analysisReport <- result
		fmt.Println("analysis the order detail")

	}
}

func storeAnalysisReport(analysisReport <-chan AnalysisReport) {
	for report := range analysisReport {
		// 	store to big query (example)
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("stored the analysis report to BQ with result: %+v\n\n", report)
	}
}

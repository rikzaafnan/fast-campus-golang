package main

import "sync"

func main() {

}

func FetchPricing() {
	var wg sync.WaitGroup
	wg.Add(1)
	indosatPricingResult, indosatErr := FetchIndosatAPi(&wg)
	wg.Add(1)
	telkomselPricingResult, telkomseslErr := FetchTelkomselAPi(&wg)
	wg.Add(1)
	xlPricingResult, xlErr := FetchXLAPi(&wg)

	wg.Wait()

	// 	cache all to redis
	_ = indosatPricingResult
	_ = indosatErr
	_ = telkomselPricingResult
	_ = telkomseslErr
	_ = xlPricingResult
	_ = xlErr

}

func FetchIndosatAPi(wg *sync.WaitGroup) (data struct{}, err error) {
	defer wg.Done()
	return
}

func FetchTelkomselAPi(wg *sync.WaitGroup) (data struct{}, err error) {
	defer wg.Done()
	return
}

func FetchXLAPi(wg *sync.WaitGroup) (data struct{}, err error) {
	defer wg.Done()
	return
}

package main

import (
	"fmt"

	"github.com/PetarGeorgiev-hash/learning-go/filemanager"
	"github.com/PetarGeorgiev-hash/learning-go/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%0.f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index])
	}

	for _, done := range doneChans {
		<-done
	}

}

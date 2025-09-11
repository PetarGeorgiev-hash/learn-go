package main

import (
	"fmt"

	"github.com/PetarGeorgiev-hash/learning-go/filemanager"
	"github.com/PetarGeorgiev-hash/learning-go/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%0.f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}

}

package main

import (
	"price-calculator/prices"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	taxRates := []float64{0, 0.5, 0.1}

	for _, taxRate := range taxRates {
		pricesJob := prices.NewTaxIncludedPriceJob(taxRate)
		pricesJob.Process()
	}
}

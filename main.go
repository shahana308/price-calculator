package main

import "fmt"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.5, 0.1}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		includedTaxPrice := make([]float64, len(prices))
		for priceIndex, price := range prices {
			includedTaxPrice[priceIndex] = price * (1 + taxRate)
		}
		result[taxRate] = includedTaxPrice
	}
	fmt.Println(result)
}

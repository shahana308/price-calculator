package prices

import (
	"fmt"
	"price-calculator/conversion"
	"price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate          float64
	InputPrice       []float64
	TaxIncludedPrice map[string]string
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := filemanager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println("Error reading file data:", err)
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println("Error converting file data:", err)
	}

	job.InputPrice = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrice {
		taxIncludedString := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedString)
	}
	job.TaxIncludedPrice = result

	filemanager.WriteJSON(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrice: []float64{10, 20, 30},
		TaxRate:    taxRate,
	}
}

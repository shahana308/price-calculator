package prices

import (
	"bufio"
	"fmt"
	"os"
	"price-calculator/conversion"
)

type TaxIncludedPriceJob struct {
	TaxRate          float64
	InputPrice       []float64
	TaxIncludedPrice map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
		file.Close()
		return
	}
	file.Close()

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println("Error converting file data:", err)
		file.Close()
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
	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrice: []float64{10, 20, 30},
		TaxRate:    taxRate,
	}
}

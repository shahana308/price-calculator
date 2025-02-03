package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	prices := make([]float64, len(lines))
	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)

		if err != nil {
			fmt.Println("Converting string to float failed:", err)
			file.Close()
			return
		}

		prices[lineIndex] = floatPrice
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

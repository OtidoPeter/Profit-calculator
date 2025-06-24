package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	outputText("Revenue: ")
	fmt.Scan(&revenue)

	outputText("Expenses: ")
	fmt.Scan(&expenses)

	outputText("Tax Rate: ")
	fmt.Scan(&taxRate)

	operations(revenue, expenses, taxRate)
	/*ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
	*/
}

func outputText(text string) {
	fmt.Print(text)
}

func operations(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	reformattedEbt := fmt.Sprintf("Earnings before tax: %.2f\n", ebt)
	reformattedProfit := fmt.Sprintf("Earnings after tax: %.2f\n", profit)
	reformattedRatio := fmt.Sprintf("Ratio: %.2f\n", ratio)

	fmt.Print(reformattedEbt, reformattedProfit, reformattedRatio)

	return ebt, profit, ratio
}

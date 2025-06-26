package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getUserInput("Revenue: ")
	// fmt.Scan(&revenue)

	expenses = getUserInput("Expenses: ")
	// fmt.Scan(&expenses)

	taxRate = getUserInput("Tax Rate: ")
	// fmt.Scan(&taxRate)

	operations(revenue, expenses, taxRate)
	/*ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
	*/
}

func getUserInput(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	return userInput
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

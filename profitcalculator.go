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

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}

func outputText(text string) {
	fmt.Print(text)
}

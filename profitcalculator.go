package main

import (
	"fmt"
)

func main() {
	var taxRate float64
	var revenue float64
	var expenses float64

	// fmt.Print("Revenue: ")
	revenue = outputNum("Revenue: ")
	// fmt.Scan(&revenue)

	// fmt.Print("Expenses: ")
	expenses = outputNum("Expenses: ")
	// fmt.Scan(&expenses)

	// fmt.Print("Tax rate: ")
	taxRate = outputNum("Tax rate: ")
	// fmt.Scan(&taxRate)

	earningsBeforeTax, earningsAfterTax, ratio := calculateEarnings(revenue, expenses, taxRate)
	// earningsAfterTax := (earningsBeforeTax) * (1 - taxRate/100)
	// ratio := earningsBeforeTax / earningsAfterTax

	fmt.Printf("%.1f\n", earningsBeforeTax)
	fmt.Printf("%.1f\n", earningsAfterTax)
	fmt.Printf("%.1f\n", ratio)
	//fmt.Println(earningsBeforeTax)
	//fmt.Println(earningsAfterTax)
	// fmt.Println(ratio)
}

func outputNum(num string) float64 {
	var userInput float64
	fmt.Print(num)
	fmt.Scan(&userInput)
	return userInput
}

func calculateEarnings(revenue, expenses, taxRate float64) (float64, float64, float64) {
	//
	ebt := revenue - expenses
	eat := (ebt) * (1 - taxRate/100)
	rat := ebt / eat
	return ebt, eat, rat
}

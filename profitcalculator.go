package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

/*Goals
1)validate user input
-show error message and exit if invalid input is provided
-No negative numbers
-Not 0
2) Store calculated results into file*/

func getResultFromFile() float64 {
	data, _ := os.ReadFile("results.txt")
	resultsText := string(data)
	results, _ := strconv.ParseFloat(resultsText, 64)
	return results
}

func main() {
	var taxRate float64
	var revenue float64
	var expenses float64

	revenue, err := outputNum("Revenue: ")

	if err != nil {
		fmt.Println(err)
		return
		//panic(err)
	}

	expenses, err = outputNum("Expenses: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err = outputNum("Tax rate: ")

	if err != nil {
		fmt.Println(err)
	}

	// if err1 != nil || err2 != nil || err3 != nil {
	//	fmt.Println(err1)
	//	return
	//}

	ebt, eat, rat := calculateEarnings(revenue, expenses, taxRate)
	// earningsAfterTax := (earningsBeforeTax) * (1 - taxRate/100)
	// ratio := earningsBeforeTax / earningsAfterTax

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", eat)
	fmt.Printf("%.1f\n", rat)
	storeResults(ebt, eat, rat)
}

func storeResults(ebt, eat, rat float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, eat, rat)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func outputNum(num string) (float64, error) {
	var userInput float64
	fmt.Print(num)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value must be a positive number.")
	}
	return userInput, nil
}

func calculateEarnings(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	eat := (ebt) * (1 - taxRate/100)
	rat := ebt / eat
	return ebt, eat, rat
}

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

const resultsEarningsFile = "results.txt"

func getResultFromFile() float64 {
	data, _ := os.ReadFile(resultsEarningsFile)
	resultsText := string(data)
	results, _ := strconv.ParseFloat(resultsText, 64)
	return results
}

func resultsIntoFile(results float64) {
	resultsText := fmt.Sprint(results)
	os.WriteFile(resultsEarningsFile, []byte(resultsText), 0644)
}

func main() {
	var taxRate float64
	var revenue float64
	var expenses float64

	revenue, err1 := outputNum("Revenue: ")

	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	expenses, err2 := outputNum("Expenses: ")

	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	taxRate, err3 := outputNum("Tax rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}

	ebt, eat, rat := calculateEarnings(revenue, expenses, taxRate)
	// earningsAfterTax := (earningsBeforeTax) * (1 - taxRate/100)
	// ratio := earningsBeforeTax / earningsAfterTax

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", eat)
	fmt.Printf("%.1f\n", rat)
	//fmt.Println(earningsBeforeTax)
	//fmt.Println(earningsAfterTax)
	// fmt.Println(ratio)
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

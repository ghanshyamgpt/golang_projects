package main

import (
	"fmt"
	"strconv"
	"strings"
)

func validateCreditCard(cardNumber string) bool {
	cardNumber = strings.ReplaceAll(cardNumber, " ", "") // Remove spaces

	if _, err := strconv.Atoi(cardNumber); err != nil {
		fmt.Println("Invalid characters in the card number.")
		return false
	}

	// Apply Luhn algorithm
	sum := 0
	parity := len(cardNumber) % 2
	for i, r := range cardNumber {
		digit, _ := strconv.Atoi(string(r))
		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}

func main() {
	var cardNumber string
	fmt.Print("Enter credit card number: ")
	fmt.Scanln(&cardNumber)

	if validateCreditCard(cardNumber) {
		fmt.Println("Valid credit card number.")
	} else {
		fmt.Println("Invalid credit card number.")
	}
}

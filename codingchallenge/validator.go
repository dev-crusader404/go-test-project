package codingchallenge

import (
	"fmt"
	"regexp"
	"strings"
)

func RunCreditCardValidator() {
	var n int
	fmt.Println("Please enter the number of credit card you want to validate: ")
	fmt.Scan(&n)
	if n < 1 || n >= 100 {
		fmt.Printf("\nInvalid input: %d\n", n)
		return
	}

	for i := 0; i < n; i++ {
		var cardNumber string
		fmt.Scan(&cardNumber)
		if isValidCrediCardNumber(cardNumber) {
			fmt.Println("Valid")
		} else {
			fmt.Println("Invalid")
		}
	}

}

func isValidCrediCardNumber(cardNumber string) bool {
	regexMatcher := `^(4|5|6)\d{3}-?\d{4}-?\d{4}-?\d{4}$`

	re := regexp.MustCompile(regexMatcher)
	if !re.MatchString(cardNumber) {
		return false
	}

	cardNumber = strings.ReplaceAll(cardNumber, "-", "")

	if len(cardNumber) != 16 {
		return false
	}

	for i := 0; i < len(cardNumber)-3; i++ {
		if cardNumber[i] == cardNumber[i+1] && cardNumber[i] == cardNumber[i+2] && cardNumber[i] == cardNumber[i+3] {
			return false
		}
	}
	return true
}

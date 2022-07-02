package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	input = strings.Replace(input, " ", "", -1)
	if input == "" {
		return "", fmt.Errorf("an error ocured: %w", errorEmptyInput)
	}
	numbers := getNumbers(input)
	if len(numbers) != 2 {
		return "", fmt.Errorf("an error ocured: %w", errorNotTwoOperands)
	}

	numberOne, err := strconv.Atoi(numbers[0])
	if err != nil {
		return "", fmt.Errorf("an error ocured: %w", err)
	}
	numberTwo, err := strconv.Atoi(numbers[1])
	if err != nil {
		return "", fmt.Errorf("an error ocured: %w", err)
	}

	return fmt.Sprint(numberOne + numberTwo), nil
}

func getNumbers(input string) []string {
	var sb strings.Builder
	numbers := make([]string, 0)
	for i := 0; i < len(input); i++ {
		if (string(input[i]) == "+" || string(input[i]) == "-") && i != 0 {
			numbers = append(numbers, sb.String())
			sb.Reset()
		}
		sb.WriteByte(input[i])
	}
	numbers = append(numbers, sb.String())
	return numbers
}

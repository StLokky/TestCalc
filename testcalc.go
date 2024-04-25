package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isLatin(s string) bool {
	if strings.IndexAny(s, "IVX") == -1 {
		return false
	}
	return true
}

var latinNumerals = map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10, "XI": 11, "XII": 12, "XIII": 13, "XIV": 14,
	"XV": 15, "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20}

func lat2Int(s string) int {
	return latinNumerals[string(s)]
}

func int2Lat(n int) string {
	for key, val := range latinNumerals {
		if val == n {
			return key
		}
	}
	return "" // Возвращаем пустую строку, если значение не найдено.
}

func calculate(str string) (int, bool) {
	var result int
	var flagLatin bool

	elements := strings.Fields(str)

	if len(elements) != 3 {
		panic("Incorrect format. Please enter two numbers and an operator separated by spaces.")
	}

	num1 := elements[0]
	op := elements[1]
	num2 := elements[2]

	if !((isNumeric(num1) && isNumeric(num2)) || (isLatin(num1) && isLatin(num2))) {
		panic("Numbers should be either all Arabic or all Roman numerals")
	}

	num1Int, _ := strconv.Atoi(num1)
	num2Int, _ := strconv.Atoi(num2)

	if isLatin(num1) {
		flagLatin = true
		num1Int = lat2Int(num1)
		num2Int = lat2Int(num2)
	}

	if num1Int < 1 || num1Int > 10 || num2Int < 1 || num2Int > 10 {
		panic("Wrong numbers...")
	}

	switch op {
	case "+":
		result = num1Int + num2Int
	case "-":
		result = num1Int - num2Int
	case "*":
		result = num1Int * num2Int
	case "/":
		result = num1Int / num2Int
	default:
		panic("Invalid operator. Please use +, -, * or /")
	}
	return result, flagLatin
}

func main() {
	var result int
	var resultStr string
	var flagLatin bool

	fmt.Println("\nProgram 'Console Calculator'.")
	fmt.Println("Please enter two numbers and an operator separated by spaces.")
	fmt.Println("You can use Arabic or Latin numerals from 1 to 10")
	fmt.Println("You can only use operators from this list: '+', '-', '*', '/'\n")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter an expression or 'q' to quit: ")
		scanner.Scan()
		input := scanner.Text()
		if input == "q" {
			fmt.Println("Goodbye...")
			break
		}

		result, flagLatin = calculate(input)
		if flagLatin {
			resultStr = int2Lat(result)
			switch {
			case result <= 0:
				panic("There are no negatives in Latin numerals.....")
			case resultStr == "":
				panic("Wrong expession...")
			}
			fmt.Println(input, " = ", int2Lat(result))
		} else {
			fmt.Println(input, " = ", result)
		}
	}
}

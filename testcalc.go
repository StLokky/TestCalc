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
	validStr := "IVXLCDM"
	for _, ch := range s {
		if !strings.ContainsAny(string(ch), validStr) {
			return false
		}
	}
	return true
}

type IntLat struct {
	val int
	lat string
}

var arabicNumerals = []IntLat{{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"}, {100, "C"}, {90, "XC"},
	{50, "L"}, {40, "XL"}, {10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"}}

var latinNumerals = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

func lat2Int(latin string) int {
	result := 0
	temp := 0
	for i := len(latin) - 1; i >= 0; i-- {
		value := latinNumerals[string(latin[i])]
		if value < temp {
			result -= value
		} else {
			result += value
		}
		temp = value
	}
	return result
}

func int2Lat(num int) string {
	result := ""
	for _, lat := range arabicNumerals {
		for num >= lat.val {
			result += lat.lat
			num -= lat.val
		}
	}
	return result
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
				panic("There are no negatives or zero in Latin numerals.....")
			case resultStr == "":
				panic("Wrong expession...")
			}
			fmt.Println(input, " = ", int2Lat(result))
		} else {
			fmt.Println(input, " = ", result)
		}
	}
}

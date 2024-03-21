package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if len(input) == 0 {
		panic("необходимо ввести выражение")
	}

	fmt.Println("Введенное выражение:", input)

	fmt.Println("Результат: ", calculate(input))

	// Раскомментируйте строки ниже и закомментируйте строки выше, чтобы удобно тестировать :)

	// inputs := []string{"1.5 / 1", "VI / III", "I - II", "I + 1", "1", "1 + 2 + 3", "1 + 5", "1 + 10", "10 * 10", "X * X", "X / X"}
	// for _, v := range inputs {
	// 	func() {
	// 		defer func() {
	// 			if err := recover(); err != nil {
	// 				fmt.Println("Recovered from panic:", err)
	// 			}
	// 		}()

	// 		fmt.Println(calculate(v))
	// 	}()
	// }

}

func calculate(input string) string {
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		panic("некорректный формат выражения")
	}

	a := parseNumber(parts[0])
	b := parseNumber(parts[2])
	operator := parts[1]

	isRomanResult := isRoman(parts[0]) && isRoman(parts[2])

	if (isRoman(parts[0]) && !isRoman(parts[2])) || (!isRoman(parts[0]) && isRoman(parts[2])) {
		panic("используются разные системы счисления")
	}

	switch operator {
	case "+":
		if isRomanResult {
			return arabicToRoman(a + b)
		}
		return strconv.Itoa(a + b)
	case "-":
		if isRomanResult {
			if a <= b {
				panic("в римской системе нет отрицательных чисел и нуля")
			}
			return arabicToRoman(a - b)
		}
		return strconv.Itoa(a - b)
	case "*":
		if isRomanResult {
			return arabicToRoman(a * b)
		}
		return strconv.Itoa(a * b)
	case "/":
		if b == 0 {
			panic("деление на ноль")
		}
		if isRomanResult {
			if a < b {
				panic("в римской системе нет дробных чисел")
			}
			return arabicToRoman(a / b)
		}
		return strconv.Itoa(a / b)
	default:
		panic("некорректный операнд")
	}
}

func parseNumber(s string) int {

	var number int
	var err error

	if isRoman(s) {
		number = romanToArabic(s)
	} else {
		number, err = strconv.Atoi(s)
		if err != nil {
			panic("некорректный формат числа (калькулятор может работать только с целыми арабскими и римскими цифрами от 0 до 10 и от I до X)")
		}
	}

	if number > 10 {
		panic("числа должны быть от 0 до 10")
	}

	return number
}

func isRoman(s string) bool {
	for _, c := range s {
		if c != 'I' && c != 'V' && c != 'X' {
			return false
		}
	}
	return true
}

func romanToArabic(s string) int {
	romanValues := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}

	var result int
	var prevValue int
	var temp int

	for i := len(s) - 1; i >= 0; i-- {
		curValue := romanValues[rune(s[i])]

		if curValue >= prevValue {
			temp += curValue
		} else {
			temp -= curValue
		}

		if i == 0 || curValue <= romanValues[rune(s[i-1])] {
			result += temp
			temp = 0
		}

		prevValue = curValue
	}

	if result < 1 || result > 10 {
		panic("число должно быть от I до X")
	}

	return result
}

type romanNumeral struct {
	value  int
	symbol string
}

func arabicToRoman(number int) string {

	numerals := []romanNumeral{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result strings.Builder

	for i, num := range numerals {
		count := number / num.value
		number -= count * num.value
		for j := 0; j < count; j++ {
			result.WriteString(numerals[i].symbol)
		}
	}

	return result.String()
}

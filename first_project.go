package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Приветствую! Это мой калькулятор: ")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение (с пробелами): ")
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)
	
	parts := strings.Fields(input)
	if len(parts) != 3 {
		fmt.Println("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		os.Exit(1)
	}

	aStr, operator, bStr := parts[0], parts[1], parts[2]

	isRomanMode := isRoman(aStr) && isRoman(bStr)
	isArabicMode := !isRoman(aStr) && !isRoman(bStr)

	if !isRomanMode && !isArabicMode {
		fmt.Println("Выдача паники, так как используются одновременно разные системы счисления.")
		os.Exit(1)
	}

	var first, second int
	var err error

	if isRomanMode {
		first, err = romanToInt(aStr)
		if err != nil {
			fmt.Println("Паника:", err)
			os.Exit(1)
		}
		second, err = romanToInt(bStr)
		if err != nil {
			fmt.Println("Паника:", err)
			os.Exit(1)
		}
	} else {
		
		first, err = strconv.Atoi(aStr)
		if err != nil || first < 1 || first > 10 {
			fmt.Println("Выдача паники: некорректное арабское число")
			os.Exit(1)
		}
		second, err = strconv.Atoi(bStr)
		if err != nil || second < 1 || second > 10 {
			fmt.Println("Выдача паники: некорректное арабское число")
			os.Exit(1)
		}
	}

	result, err := calculate(first, second, operator)
	if err != nil {
		fmt.Println("Паника:", err)
		os.Exit(1)
	}

	if isRomanMode {
		romanResult, err := intToRom(result)
		if err != nil {
			fmt.Println("Паника:", err)
			os.Exit(1)
		}
		fmt.Println("Результат:", romanResult)
	} else {
	
		fmt.Println("Результат:", result)
	}
}

var romanToArab = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	"XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15,
	"XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20,
	"XXX": 30, "XL": 40, "L": 50, "LX": 60, "LXX": 70, "LXXX": 80, "XC": 90, "C": 100,
}

var arabicToRom = []struct {
	value int
	sym   string
}{
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

func romanToInt(s string) (int, error) {
	if val, ok := romanToArab[s]; ok {
		return val, nil
	}
	return 0, errors.New("Выдача паники, некорректное римское число")
}

func intToRom(num int) (string, error) {
	if num <= 0 {
		return "", errors.New("Выдача паники, так как в римской системе нет отрицательных чисел.")
	}
	var result strings.Builder
	for _, entry := range arabicToRom {
		for num >= entry.value {
			result.WriteString(entry.sym)
			num -= entry.value
		}
	}
	return result.String(), nil
}

func isRoman(s string) bool {
	_, ok := romanToArab[s]
	return ok
}

func calculate(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("Выдача паники, деление на ноль")
		}
		return a / b, nil
	default:
		return 0, errors.New("Выдача паники, неизвестный оператор")
	}
}

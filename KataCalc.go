package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Введите выражение через пробел: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Split(input, " ")

	var operatorCheck bool = false
	var num1, num2 int
	var num1Check, num2Check string
	operations := []string{"+", "-", "*", "/"}
	romanic := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	arabic := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	if len(parts) < 3 {
		fmt.Println("Cтрока не является математической операцией")
		return
	} else if len(parts) > 3 {
		fmt.Println("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
		return
	}

	for _, operator := range operations {
		if parts[1] == operator {
			operatorCheck = true
			break
		}
	}
	if operatorCheck == false {
		fmt.Println("Некорректная операция")
		return
	}

	for i, number := range romanic {
		if parts[0] == number {
			num1 = i + 1
			num1Check = "romanic"
		}
		if parts[2] == number {
			num2 = i + 1
			num2Check = "romanic"
		}
	}
	for i, number := range arabic {
		if parts[0] == number {
			num1 = i + 1
			num1Check = "arabic"
		}
		if parts[2] == number {
			num2 = i + 1
			num2Check = "arabic"
		}
	}
	if num1Check == "arabic" && num2Check == "arabic" {
		if parts[1] == "+" {
			fmt.Println(num1+num2)
			return
		} else if parts[1] == "-" {
			fmt.Println(num1-num2)
			return
		} else if parts[1] == "*" {
			fmt.Println(num1*num2)
			return
		} else  {
			fmt.Println(num1/num2)
			return
		}
	} else if num1Check == "romanic" && num2Check == "romanic" {
		if parts[1] == "+" {
			fmt.Println(arabicToRoman(num1+num2))
			return
		} else if parts[1] == "-" {
			fmt.Println(arabicToRoman(num1-num2))
			return
		} else if parts[1] == "*" {
			fmt.Println(arabicToRoman(num1*num2))
			return
		} else  {
			fmt.Println(arabicToRoman(num1/num2))
			return
		}

	} else {
		fmt.Println("Используются одновременно разные системы счисления")
		return
	}
}

func arabicToRoman(num int) string {
	if num < 1 {
		return "Римское число не может быть 0 или отрицательным"
	}
	if num > 100{
		return "Недопустимое число"
	}
	val := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	syb := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	romanNum := ""
	i := 0

	for num > 0 {
		if num-val[i] >= 0 {
			num -= val[i]
			romanNum += syb[i]
		} else {
			i++
		}
	}

	return romanNum
}


package main

import (
	"fmt"
)

func main() {
	num1 := 0
	num2 := 0
	operator := ""

	fmt.Print("Введите число num1: ")
	fmt.Scan(&num1)

	fmt.Print("Введите число num2: ")
	fmt.Scan(&num2)

	fmt.Print("Укажите оператор + или - или / или *: ")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		fmt.Println("Сумма чисел ровна: ", num1+num2)
	case "-":
		fmt.Println("Разность чисел ровна: ", num1-num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Делить на ноль нельзя")
		} else {
			fmt.Println("Деление чисел ровно: ", num1/num2)
		}
	case "*":
		fmt.Println("Произведение чисел ровно: ", num1*num2)
	}
}

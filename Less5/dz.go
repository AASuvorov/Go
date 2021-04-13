package main

import (
	"fmt"
)

// Задание №1
func fibbonachi(n uint) uint {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibbonachi(n-1) + fibbonachi(n-2)
}

// ------------------ //

// Задание №2
func fibbonachiMap(z2 uint) uint {

	if z2 == 0 {
		return 0
	}
	if z2 == 1 {
		return 1
	}
	// return fibbonachi(n-1) + fibbonachi(n-2)
	exampleMap := map[uint]uint{}
	exampleMap[z2] = fibbonachiMap(z2-1) + fibbonachiMap(z2-2)

	return exampleMap[z2]
}

// ------------------ //

// Задание №3
type Calc struct {
	FirstOp  float64
	Operator string
	SecondOp float64
}

func (c Calc) Sum() float64 {
	return c.FirstOp + c.SecondOp
}

func (c Calc) Mult() float64 {
	return c.FirstOp * c.SecondOp
}

func (c Calc) Delenie() float64 {
	return c.FirstOp / c.SecondOp
}

func (c Calc) Raznost() float64 {
	return c.FirstOp - c.SecondOp
}

// ------------------ //

func main() {
	// Задание №1
	var z1 uint
	fmt.Println("Введите положительное число")
	fmt.Scanf("%d\n", &z1)

	fmt.Println(fibbonachi(z1))
	// ------------------ //

	// Задание №2
	var z2 uint
	fmt.Println("Введите положительное число")
	fmt.Scanf("%d\n", &z2)

	fmt.Println(fibbonachiMap(z2))
	// ------------------ //

	// Задание №3
	num1 := 0.0
	num2 := 0.0
	operator := ""

	fmt.Println("Введите число num1 затем оператор +,-,*,/ и затем num2")
	fmt.Scan(&num1, &operator, &num2)

	c := Calc{
		FirstOp:  num1,
		Operator: operator,
		SecondOp: num2,
	}

	switch operator {
	case "+":
		fmt.Println(c.Sum())
	case "*":
		fmt.Println(c.Mult())
	case "/":
		fmt.Println(c.Delenie())
	case "-":
		fmt.Println(c.Raznost())
	}
	// ------------------ //
}

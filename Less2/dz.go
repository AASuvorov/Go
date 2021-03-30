package main

import (
	"fmt"
	"math"
)

func main() {
	//	Задание 1
	var x, y, area float64
	fmt.Print("Введите длину прямоугольника: ")
	fmt.Scanln(&x)
	fmt.Print("Введите ширину прямоугольника: ")
	fmt.Scanln(&y)
	area = x * y
	fmt.Println("Площадь прямоугольника: ", area)

	//	Задание 2
	var radius, circle float64
	fmt.Print("Введите радиус: ")
	fmt.Scanln(&radius)
	circle = radius * radius * math.Pi
	fmt.Println("Площадь прямоугольника: ", circle)

	//	Задание 3
	var n, c1, c2, c3 int
	fmt.Print("Введите трехзначное число: ")
	fmt.Scanln(&n)
	if n < 100 {
		fmt.Println("Необходимо ввести трехзначное число")
	} else {
		c1 = n / 100
		fmt.Println("сотен: ", c1)

		c2 = (n / 10) % 10
		fmt.Println("десятков: ", c2)

		c3 = n % 10
		fmt.Println("единиц: ", c3)
	}
}

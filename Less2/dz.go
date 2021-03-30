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
}

package main

import (
	"fmt"
)

type Calc struct {
	FirstOp, SecondOp float64
	Operator          string
}

func (c *Calc) Sum() float64 {
	return (*c).FirstOp + (*c).SecondOp
}

func (c *Calc) Mult() float64 {
	return (*c).FirstOp * (*c).SecondOp
}

func (c *Calc) Delenie() float64 {

	return (*c).FirstOp / (*c).SecondOp
}

func (c *Calc) Raznost() float64 {
	return (*c).FirstOp - (*c).SecondOp
}

func main() {
	n1 := 0.0
	n2 := 0.0
	operator := ""

	fmt.Println("Введите число num1 затем оператор +,-,*,/ и затем num2")
	fmt.Scan(&n1, &operator, &n2)

	c := Calc{
		FirstOp:  n1,
		Operator: operator,
		SecondOp: n2,
	}

	var cPointer *Calc = &c

	switch operator {
	case "+":
		fmt.Println(cPointer.Sum())
	case "*":
		fmt.Println(cPointer.Mult())
	case "/":
		fmt.Println(cPointer.Delenie())
	case "-":
		fmt.Println(cPointer.Raznost())
	}
}

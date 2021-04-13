package calc

type Calc struct {
	FirstOp  float64
	Operator string
	SecondOp float64
}

func (c Calc) Sum() float64 {
	return c.FirstOp + c.SecondOp
}

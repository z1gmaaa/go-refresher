package main

import "github.com/slashpai/go-refresher/anonymousfunctions/calculator"

type MathExpr = string

//Constants
const (
	ADD      = MathExpr("add")
	SUBTRACT = MathExpr("subtract")
	MULTIPLY = MathExpr("multiply")
	DIVIDE   = MathExpr("divide")
)

func main() {
	cal := calculate(ADD)
	println(cal(2, 3))
	cal = calculate(SUBTRACT)
	println(cal(2, 3))
	cal = calculate(MULTIPLY)
	println(cal(2, 3))
	cal = calculate(DIVIDE)
	println(cal(2, 3))
	cal = calculate(DIVIDE)
	println(cal(2, 0))
	cal = calculate(MathExpr("UNKNOWN"))
	println(cal(2, 3))
}

func calculate(expr MathExpr) func(float64, float64) (float64, error) {
	switch expr {
	case ADD:
		return calculator.Add
	case SUBTRACT:
		return calculator.Subtract
	case MULTIPLY:
		return calculator.Multiply
	case DIVIDE:
		return calculator.Divide
	default:
		{
			return func(num1, num2 float64) (float64, error) {
				return 0, nil
			}
		}
	}
}

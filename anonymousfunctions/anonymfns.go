package main

import (
	"fmt"

	"github.com/slashpai/go-refresher/anonymousfunctions/calculator"
)

type MathExpr = string

const (
	ADD      = MathExpr("add")
	SUBTRACT = MathExpr("subtract")
	MULTIPLY = MathExpr("multiply")
	DIVIDE   = MathExpr("divide")
)

func main() {
	fmt.Println("=== Anonymous Functions Demo ===")
	anonymousFunctionDemo()

	fmt.Println("\n=== Calculator Demo ===")
	calculatorDemo()

	fmt.Println("\n=== Stateful Function Demo ===")
	statefulFunctionDemo()
}

// --------- Anonymous Functions -------------
func anonymousFunctionDemo() {
	// simple anonymous function
	func() {
		fmt.Println("first anonymous function")
	}()

	simpleVarFun := func() {
		fmt.Println("function as variables: example variable simpleVarFun")
	}
	simpleVarFun()

	greet := func(name string) {
		fmt.Println("Hello", name)
	}
	greet("Jake")

	add := func(num1, num2 int32) int32 {
		return num1 + num2
	}
	sum := add(1, 2)
	fmt.Println("sum =", sum)
}

// --------- Calculator Example -------------
func calculatorDemo() {
	operations := []MathExpr{ADD, SUBTRACT, MULTIPLY, DIVIDE, DIVIDE, "UNKNOWN"}
	pairs := [][2]float64{{2, 3}, {2, 3}, {2, 3}, {2, 3}, {2, 0}, {2, 3}}

	for i, expr := range operations {
		cal := calculate(expr)
		result, err := cal(pairs[i][0], pairs[i][1])
		if err != nil {
			fmt.Printf("%s: error: %v\n", expr, err)
		} else {
			fmt.Printf("%s: result: %v\n", expr, result)
		}
	}
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
		return func(num1, num2 float64) (float64, error) {
			return 0, fmt.Errorf("unsupported operation: %s", expr)
		}
	}
}

// --------- Stateful Function -------------
func statefulFunctionDemo() {
	ctr := count()
	for i := 0; i < 5; i++ {
		fmt.Println("Counter:", ctr())
	}
}

func count() func() int32 {
	x := 0
	return func() int32 {
		x++
		return int32(x)
	}
}

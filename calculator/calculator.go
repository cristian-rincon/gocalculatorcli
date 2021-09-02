package calculator

import (
	"fmt"
	"os"
)

type Calculator struct {
}

func (Calculator) Sum(a int, b int) int {
	return a + b
}

func (Calculator) Subtract(a int, b int) int {
	return a - b
}

func (Calculator) Multiply(a int, b int) int {
	return a * b
}

func (Calculator) Divide(a int, b int) int {
	return a / b
}

func (c Calculator) Calculate(option string) {

	switch option {
	case "1":
		first, second := AskNumbers("sum")
		Response(c.Sum(first, second))
	case "2":
		first, second := AskNumbers("subtract")
		Response(c.Subtract(first, second))
	case "3":
		first, second := AskNumbers("multiply")
		Response(c.Multiply(first, second))
	case "4":
		first, second := AskNumbers("divide")
		Response(c.Divide(first, second))
	default:
		fmt.Println("¡Bye!")
		os.Exit(0)
	}

}

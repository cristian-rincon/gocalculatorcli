package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Parser(input string) int {
	input_cast, _ := strconv.Atoi(input)
	return input_cast
}

func ReadInput() string {
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	return reader.Text()
}

func Response(result int) {
	fmt.Printf("The result is: %d \n", result)
}

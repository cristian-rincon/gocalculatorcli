package calculator

import "fmt"

type App struct {
}

func (a App) Banner() {

	banner := `
     ___           ___           ___       ___     
    /\  \         /\  \         /\__\     /\  \    
   /::\  \       /::\  \       /:/  /    /::\  \   
  /:/\:\  \     /:/\:\  \     /:/  /    /:/\:\  \  
 /:/  \:\  \   /::\~\:\  \   /:/  /    /:/  \:\  \ 
/:/__/ \:\__\ /:/\:\ \:\__\ /:/__/    /:/__/ \:\__\
\:\  \  \/__/ \/__\:\/:/  / \:\  \    \:\  \  \/__/
 \:\  \            \::/  /   \:\  \    \:\  \      
  \:\  \           /:/  /     \:\  \    \:\  \     
   \:\__\         /:/  /       \:\__\    \:\__\    
    \/__/         \/__/         \/__/     \/__/    
	`
	fmt.Println(banner)
}

func (a App) Line() {
	fmt.Println("====================================================")
}

func (a App) Welcome() {
	a.Line()
	a.Banner()
	a.Line()
	fmt.Println("           Welcome to the calculator!")
	a.Line()
	fmt.Println("Choose the operation you want to do:")
	fmt.Println("1) Sum")
	fmt.Println("2) Subtract")
	fmt.Println("3) Multiply")
	fmt.Println("4) Divide")
	fmt.Println("5) Quit")
	a.Line()
}

func AskNumbers(operation string) (int, int) {
	fmt.Printf("Enter the first number to %v: ", operation)
	first_number := Parser(ReadInput())
	fmt.Printf("Enter the second number to %v: ", operation)
	second_number := Parser(ReadInput())
	return first_number, second_number

}

func (a App) Run() {
	a.Welcome()
	user_option := ReadInput()
	c := Calculator{}
	c.Calculate(user_option)
}

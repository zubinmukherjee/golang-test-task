package main

import "fmt"

func main(){
	var firstNumber, secondNumber float32

	//fetch the user inputs
	fmt.Print("Please enter FIRST number: ")
	_, firstNumberErr := fmt.Scanf("%f", &firstNumber)
	if firstNumberErr != nil {
		fmt.Printf("ERROR: getting an error while fetching the FIRST number, Err(%v)", firstNumberErr)
		return
	}

	fmt.Print("Please enter SECOND number: ")
	_, secondNumberErr := fmt.Scanf("%f", &secondNumber)
	if secondNumberErr != nil {
		fmt.Printf("ERROR: getting an error while fetching the SECOND number, Err(%v)", secondNumberErr)
		return
	}

	fmt.Println("Before Swap:")
	fmt.Println("a =", firstNumber)
	fmt.Println("b =", secondNumber)

	//swapping logic
	firstNumber = firstNumber - secondNumber
	secondNumber = firstNumber + secondNumber
	firstNumber = secondNumber - firstNumber

	fmt.Println("After Swap:")
	fmt.Println("a =", firstNumber)
	fmt.Println("b =", secondNumber)
}

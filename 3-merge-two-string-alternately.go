package main

import "fmt"

func main()  {
	var firstString, secondString string

	//fetch input details
	fmt.Print("Please enter FIRST string: ")
	_, firstStringErr := fmt.Scanf("%s", &firstString)
	if firstStringErr != nil {
		fmt.Printf("Error: expected input as a string, Err(%v)", firstStringErr)
		return
	}

	fmt.Print("Please enter SECOND string: ")
	_, secondStringErr := fmt.Scanf("%s", &secondString)
	if secondStringErr != nil {
		fmt.Printf("Error: expected input as string, Err(%v)", secondStringErr)
		return
	}

	var l1, l2 = len(firstString), len(secondString)
	var result = ""

	//core logic
	for i, j := 0, 0; i < l1 || j < l2; i, j = i + 1, j + 1 {
		if i < l1 {
			result += string(firstString[i])
		}
		if j < l2 {
			result += string(secondString[j])
		}
	}

	//print output
	fmt.Printf("String 1 : %s\n", firstString)
	fmt.Printf("String 2 : %s\n", secondString)
	fmt.Printf("Result :%s\n", result)

	return
}
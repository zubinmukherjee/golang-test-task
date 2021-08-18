package main

import (
	"fmt"
	"math"
)

const suffixString = "[" //this is the next highest ASCII value after "Z"
const maxTestCase = 5 // max test case value

func main() {
	var testCases int
	var stringA, stringB []string

	//fetch input details
	fmt.Print("Please enter the number of test cases you want to perform (max 5): ")
	_, testCasesErr := fmt.Scanf("%d", &testCases)
	if testCasesErr != nil {
		fmt.Printf("Error: expected input as a number, Err(%v)", testCasesErr)
		return
	}

	if testCases > maxTestCase {
		fmt.Println("INVALID INPUT: Test case must be less than or equals to 5")
		return
	}

	for i := 0; i < testCases; i++ {
		var string1, string2 string

		fmt.Printf("Please enter StringA of testcase %d : ", i+1)
		_, string1Err := fmt.Scanf("%s", &string1)
		if string1Err != nil {
			fmt.Printf("Error: expected input as a string only, Err(%v)", string1Err)
			return
		}
		if !IsUpperCase(string1) {
			fmt.Println("Error: expected input string1 in Uppercase only")
			return
		}
		if len(string1) < 1  || len(string1) > int(math.Pow(10, 5)) {
			fmt.Println("Error: String1 values is invalid. Must be greater then 1 and less then 10^5 characters")
			return
		}
		stringA = append(stringA, string1)

		fmt.Printf("Please enter StringB of testcase %d : ", i+1)
		_, string2Err := fmt.Scanf("%s", &string2)
		if string2Err != nil {
			fmt.Printf("Error: expected input as a string in Uppercase only, Err(%v)", string2Err)
			return
		}
		if !IsUpperCase(string2) {
			fmt.Println("Error: expected input string2 in Uppercase only")
			return
		}
		if len(string2) < 1  || len(string2) > int(math.Pow(10, 5)) {
			fmt.Println("Error: String2 values is invalid. Must be greater then 1 and less then 10^5 characters")
			return
		}
		stringB = append(stringB,string2)
	}

	//convert the input into lexicographically string logic
	for j := 0; j < len(stringA); j++ {
		lexicographicallyMinimalString := LexicographicallyMinimalString(stringA[j], stringB[j], "")

		//print output
		fmt.Printf("\n\nINPUT string is %s & %s\n", stringA[j], stringB[j])
		fmt.Printf("Lexicographically Minimal String output is %s\n\n", lexicographicallyMinimalString)
	}

}

//LexicographicallyMinimalString is returns the lexicographically minimal string
func LexicographicallyMinimalString(str1 , str2, finalString string) string {
	str1 += suffixString
	str2 += suffixString
	stringLength := len(str1) + len(str2) - 2

	for i := 0; i < stringLength; i++ {
		if str1 < str2 {
			finalString += str1[0:1]
			str1 = str1[1:]
		} else {
			finalString += str2[0:1]
			str2 = str2[1:]
		}
	}

	return finalString
}

//IsUpperCase it will check the string is in Uppercase or not
func IsUpperCase(s string) bool {
	for _, r := range s {
		if r < 'A' || r > 'Z' {
			return false
		}
	}
	return true
}

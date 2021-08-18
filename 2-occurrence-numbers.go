package main

import (
	"fmt"
	"sort"
)

func main() {
	var elementsArray []int
	var size int

	//fetch input size
	fmt.Printf("Enter size of elements array: ")
	_, sizeErr := fmt.Scanf("%d", &size)
	if sizeErr != nil {
		fmt.Printf("Error: size of Elements array should be an integer, Err(%v)", sizeErr)
		return
	}

	//fetch input details
	for i := 0; i < size; i++ {
		var element int
		fmt.Printf("Please enter an Element %d: ", i+1)
		_, elementErr := fmt.Scanf("%d", &element)
		if elementErr != nil {
			fmt.Printf("Error: Element type should be an integer, Err(%v)", elementErr)
			return
		}
		elementsArray = append(elementsArray, element)
	}

	//sort the array
	sort.Ints(elementsArray)

	//count logic
	uniqueCounts := DuplicateCount(elementsArray)

	//print output
	fmt.Println("number --> total")
	for k, v := range uniqueCounts {
		fmt.Printf("%d   -->   %d \n", k, v)
	}

	return
}

//DuplicateCount is used to count occurrences of the numbers
func DuplicateCount(list []int) map[int]int {

	duplicateFrequency := make(map[int]int)

	for _, item := range list {
		_, exist := duplicateFrequency[item]
		if exist {
			duplicateFrequency[item] += 1
		} else {
			duplicateFrequency[item] = 1
		}
	}
	return duplicateFrequency
}

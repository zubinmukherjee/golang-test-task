package main

import (
	"fmt"
	"math"
)

func main() {
	var n,m int
	var aVal, bVal, kVal int
	var a,b, k []int

	//fetch input details
	fmt.Println("Enter N Value: ")
	_, nErr := fmt.Scanf("%d", &n)
	if nErr != nil {
		fmt.Printf("ERROR: N value should be integer only, Err(%v)", nErr)
		return
	}
	if n < 3  || n > int(math.Pow(10, 7)){
		fmt.Println("N values is invalid. Must be greater then 3 and less then 10^7")
		return
	}

	fmt.Println("Enter M Value: ")
	_, mErr := fmt.Scanln(&m)
	if mErr != nil {
		fmt.Printf("ERROR: M value should be integer only, Err(%v)", mErr)
		return
	}
	if m < 1  || m > 2*int(math.Pow(10, 5)){
		fmt.Println("M values is invalid. Must be greater then 1 and less then 2*10^5")
		return
	}

	fmt.Println("Enter a b k Value: ")
	for i := 0; i < m; i++ {
		_, abkErr := fmt.Scanf("%d %d %d", &aVal,&bVal,&kVal)
		if abkErr != nil {
			fmt.Printf("ERROR: A B K value should be integer only, Err(%v)", abkErr)
			return
		}
		if bVal > n {
			fmt.Println("b values is invalid. Must be less then N")
			return
		}
		if aVal > bVal{
			fmt.Println("a values is invalid. Must be less then b")
			return
		}
		if  aVal < 1 {
			fmt.Println("a values is invalid. Must be greater then 1")
			return
		}
		if  kVal > int(math.Pow(10,9)) {
			fmt.Println("K values is invalid. Must be greater then 10^9")
			return
		}
		a = append(a, aVal)
		b = append(b, bVal)
		k = append(k, kVal)
	}

	//main login to find the max number from the list
	max := FindMax(n, a, b, k, m)

	//print output
	fmt.Println("MAX >>>> ", max)

}

//FindMax function to find the maximum element from updated list
func FindMax(n int, a []int, b []int, k []int, m int) int {
	arr := make([]int, n)
	arr = append(arr, 0)

	for i := 0; i < m; i++ {
		lowerBound := a[i]
		upperBound := b[i]

		for j := lowerBound; j <= upperBound; j++ {
			arr[j] += k[i]
		}
	}

	res := float64(0)
	for i := 0; i < n; i++ {
		res = math.Max(res, float64(arr[i]))
	}

	return int(res)
}

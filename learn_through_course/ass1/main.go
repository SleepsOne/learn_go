package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, val := range arr {
		fmt.Printf("number: %v is ", val)
		if val%2 == 0 {
			fmt.Printf("even\n")
		} else {
			fmt.Printf("odd\n")
		}
	}
}

package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, s := range a {
		if s%2 == 0 {
			fmt.Println(s, "is Even")
		} else {
			fmt.Println(s, "is Odd")
		}
	}
}

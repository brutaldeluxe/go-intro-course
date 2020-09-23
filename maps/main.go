package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	printMap(colors)
}

func printMap(m map[string]string) {
	for key, val := range m {
		fmt.Println("Hex code for ", key, " is ", val)
	}
}

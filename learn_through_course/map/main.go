package main

import "fmt"

func main() {

	// var colors map[string]string
	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	printMap(colors, "white")
}

func printMap(c map[string]string, field string) {
	fmt.Println(c == nil)
	c[field] = "#ffffff"
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}

	delete(c, field)

	c = nil

	fmt.Println(c == nil)
}

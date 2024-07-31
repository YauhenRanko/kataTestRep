package main

import (
	"fmt"
	"os"
)

func main() {
	var action string

	for action != "exit" {
		fmt.Println("enter comand")
		fmt.Fscan(os.Stdin, &action)

		if action == "add" {
			var value int = 3 + 2
			fmt.Println(value)
		}
		if action == "subtracts" {
			var value int = 3 - 2
			fmt.Println(value)
		}
		if action == "multiplie" {
			var value int = 3 * 2
			fmt.Println(value)
		}
		if action == "divide" {
			var value int = 4 / 2
			fmt.Println(value)
		}
	}
}

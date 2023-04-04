// Created by: Mr Coxall
// Created on: Sep 2020
//
// This program finds the area and perimeter of a rectangle

package main

import "fmt"

func main() {
	// This function finds the volume of a right rectangular prism
	var Length int
	var Width int
	var Height int
	var Volume int

	// input
	fmt.Println("This program finds the volume of a right rectangular prism.")
	fmt.Println()
	fmt.Print("Enter the Length (mm): ")
	fmt.Scanln(&Length)
	fmt.Print("Enter the Width (mm): ")
	fmt.Scanln(&Width)
	fmt.Print("Enter the Height (mm): ")
	fmt.Scanln(&Height)

	// process
	Volume = (Length * Width * Height) / 3

	// output
	fmt.Println()
	fmt.Println("The volume is:", Volume, "mm3.")
	fmt.Println("\nDone.")
}

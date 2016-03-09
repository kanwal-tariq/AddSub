package main

import "fmt"
import "time"

func main(){
	
	var i, j, sum, diff int

	fmt.Printf("First number: ")

	fmt.Scan(&i)
	
	fmt.Printf("\n\n")

	fmt.Printf("Second number: ")

	fmt.Scan(&j)
	
	sum = i + j
	diff = i - j

	fmt.Printf("Sum of the two numbers: %d\n\n", sum)

	fmt.Printf("Difference of the two numbers: %d\n\n", diff)

	time.Sleep(1000)
}
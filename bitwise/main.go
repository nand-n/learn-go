package main

import (
	"fmt"
)

const (
	//Create a huge number by shifting a 1 bit left 100 places
	//In other words , the binary number that is 1 followed by 100 zeroes
	Big = 1 << 100
	//Shift it right again 99 places , so we end up with 1<<1 or 2
	Small = Big >> 99
)

// func needInt(x int) int {
// 	return x*10 + 1
// }

// func needFloat(x float64) float64 {
// 	return x * 0.1
// }

func main() {
	// fmt.Println(needInt(Small))
	// fmt.Println(needFloat(Small))
	// fmt.Println(needFloat(Big))

	// Number of powers to generate
	// Number of powers to generate
	numPowers := 10

	fmt.Printf("%-10s %-20s\n", "Binary", "Decimal Value")
	fmt.Println("-----------------------------------")

	for i := 0; i < numPowers; i++ {
		// Calculate the power of 2
		value := 1 << i // This is equivalent to 2^i

		// Print the binary representation and its decimal value
		fmt.Printf("%-10b %-20d\n", value, value)
	}
}

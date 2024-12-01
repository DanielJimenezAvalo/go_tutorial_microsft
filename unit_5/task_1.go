/*Write a FizzBuzz program*/

package main

import (
	"fmt"
)

func FizzBuzz(number int)(){
	switch {
	case number%3 ==0 && number%5 !=0:
		fmt.Println("Fizz")
	case number%5 ==0 && number%3 !=0:
		fmt.Println("Buzz")
	case number%5 ==0 && number%3 ==0:
		fmt.Println("FizzBuzz")
	default:
		fmt.Println(number)
	}
}

func main(){
	FizzBuzz(70)
}
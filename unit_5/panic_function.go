/*
Runtime errors make a Go program panic, such as attempting to access an array by using an out-of-bounds index or 
dereferencing a nil pointer. You can also force a program to panic.

The built-in panic() function stops the normal flow of control in a Go program. 
When you use a panic call, any deferred function calls run normally. 
The process continues up the stack until all functions return. 
The program then crashes with a log message. 
The message includes any error information and a stack trace to help you diagnose the problem's root cause.

*/

package main

import "fmt"

func highlow(high int, low int){
	if high < low {
		fmt.Println("Panic!")
		panic("highlow() low greater than high")
	}
	defer fmt.Println("deferred: highlow(",high,",",low,")")
	fmt.Println("Call: highlow(",high,",",low,")")

	highlow(high,low+1)
}

func main(){
	highlow(2,0)
	fmt.Println("program finished successfully!!!")
}
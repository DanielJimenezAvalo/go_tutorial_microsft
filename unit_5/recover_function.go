/*
Sometimes you might want to avoid a program crash and instead report the error internally. 
Or perhaps you want to clean up the mess before letting the program crash. 
For instance, you might want to close any connection to a resource to avoid more problems.

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
	defer func(){
		handler:=recover()
		if handler!=nil{
			fmt.Println("main(): recover", handler)
		}
	}()

	highlow(2,0)
	fmt.Println("program finished successfully!!!")
}
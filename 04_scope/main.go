package main

import "fmt"

func wrapper() func() int{
	x := 0
	///fmt.Println("hola") executes only once
	return func() int{
		x++
		return x
	}
}

func main(){
	myfunction := wrapper()
	fmt.Println(myfunction())
	fmt.Println(myfunction())
	fmt.Println(myfunction())
	fmt.Println(myfunction())
}
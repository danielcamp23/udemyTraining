package main

import "fmt"

func main(){
	i:=0
	fmt.Println("simple for")
	for i=0; i<10; i++{
		fmt.Println(i)
	}
	i=0
	fmt.Println("\nwhile")
	for i<10{
		i++
		fmt.Println(i)
	}


}
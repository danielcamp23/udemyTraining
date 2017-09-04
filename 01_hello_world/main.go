package main

import "fmt"

func main() {
	//fmt.Println("Hellow world")
	for i:=0; i<200; i++ {
		fmt.Printf("%d \t S%x \t %q\n", i, i, i)
	}
}


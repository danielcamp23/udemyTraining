package main

import (
	"fmt"
)
/*
//func 1
func visit(numbers []int, my_callback func(int)){
	for _, n := range numbers{
		my_callback(n)
	}
}

func main(){
	visit([]int{1,2,3,4,5}, func(n int){
		fmt.Println(n)
	})
}

*/

//func 2

func fun1(the_slice []int, the_func func (int)bool) []int {
	new_slice := []int{}
	//var new_slice []int //valid too!!!
	for _ , val := range the_slice {
		if the_func(val) {
			new_slice = append(new_slice, val)
		}
	}

	return new_slice
}

func main(){
	xfun := fun1([]int{-2,3,1,2,3,4,5}, func(n int)bool{
		return n > 1
	})

	fmt.Println(xfun)
}
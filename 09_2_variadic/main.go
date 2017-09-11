package main

import "fmt"

func average(arg []float32) float32{
	var avr float32 = 0.0
	//for index,value := range arg {
	fmt.Printf("%T", arg)
	for _,value := range arg {
			avr += value
	}

	return float32 (avr) / float32(len(arg))
}

func main(){
	//avr := average(2,3,4,5,6,7,8,9)
	slice := []float32{2,3,4,5,6,7,8,9}
	avr := average(slice)
	fmt.Println("The average is ", avr)
}
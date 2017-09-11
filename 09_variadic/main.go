package main

import "fmt"

func average(arg ...float32) float32{
	var avr float32 = 0.0
	//for index,value := range arg {
	for _,value := range arg {
			avr += value
	}

	return float32 (avr) / float32(len(arg))
}

func main(){
	avr1 := average(2,3,4,5,6,7,8,9)
	fmt.Println("The average1 is ", avr1)
	myslice := []float32{2,3,4,5,6,7,8,9}
	avr2 := average(myslice...)
	fmt.Println("The average2 is ", avr2)

}
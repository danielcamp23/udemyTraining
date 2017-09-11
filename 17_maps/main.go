package main

import "fmt"

func main(){
	//map1 := make(map[string]int) 
	var map1 = make(map[string]int)
	//var map1 = map[string]int{}
	map1["a"] = 2
	map1["b"] = 3
	map1["num"] = 4
	fmt.Println("map1",map1)

	map2 := map[string]int{"a":2, "b":3, "num":4,}
	fmt.Println("map2:",map2)
	
	value1, ok1 := map2["a"]
	value2, ok2 := map2["e"]

	fmt.Println("ok1?",ok1,"value1",value1)
	fmt.Println("ok2?",ok2,"value2",value2)

	delete(map1, "num")
	fmt.Println("map1 after delete",map1)

	if val, okey := map1["num"]; okey{
		fmt.Println("num exists:",val)
	}else{
		fmt.Println("Num does not exist anymore")
	}

		

}
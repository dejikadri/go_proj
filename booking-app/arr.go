package main

import "fmt"

func main(){
	var arrNums [3]int

	arrNums[2] = 23
	for i := 0; i < len(arrNums); i++ {
		fmt.Println(arrNums[i])
	}

	fmt.Println(arrNums)
}
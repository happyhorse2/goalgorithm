package main

import (
	"fmt"
	shuju "goalgorithm/shujuandslice"
)

func main() {

	//testthreeSumClose()
	testnumk()
}

func testthreeSum() {
	//[-1,0,1,2,-1,-4]
	var numberslice = []int{-2, 0, 0, 2, 2}
	shuju.ThreeSum(numberslice)
}

func testthreeSumClose() {
	//[-1,0,1,2,-1,-4]
	var numberslice = []int{-2, 0, 0, 2, 2}
	result := shuju.ThreeSumClosest(numberslice, -3)
	fmt.Println(result)
}

func testnumk() {
	var numberslice = []int{-2, 0, 0, 2, 2}
	result := shuju.GetLeastNumbers(numberslice, 2)
	fmt.Println(result)
}

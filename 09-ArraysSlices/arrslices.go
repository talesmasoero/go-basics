package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	slc := []int{4, 5, 6}

	slc = append(slc, arr[1:2]...)
	fmt.Println(slc)

	testArray := [5]int{1, 2, 3, 4, 5}
	testSlice := testArray[:]

	testSlice[0] = 0
	fmt.Println("Array:", testArray)
	fmt.Println("Slice:", testSlice)

	testArray[1] = 10
	fmt.Println("Array:", testArray)
	fmt.Println("Slice:", testSlice)
}

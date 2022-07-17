package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("initial :  %+v \n", arr)

	/*it shifts left side by rotation amount*/
	leftRotation(arr, 4)
	fmt.Printf("rotated left:  %+v \n", arr)

	/*it shifts right side by rotation amount*/
	rightRotation(arr, 4)
	fmt.Printf("rotated right:  %+v \n", arr)

}

func leftRotation(arr []int, rotation int) {

	var tempArr []int
	for i := 0; i < rotation; i++ {
		tempArr = arr[1:len(arr)]
		tempArr = append(tempArr, arr[0])
		arr = tempArr
	}

}

func rightRotation(arr []int, rotation int) {

	var tempArr []int
	for i := 0; i < rotation; i++ {
		tempArr[0] = arr[len(arr)-1]
		tempArr = append(tempArr, arr[0:len(arr)-1]...)
		arr = tempArr
	}

}

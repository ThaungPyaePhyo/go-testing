package main

import "fmt"

// func add(a int,b int) int {
// 		return a + b
// 	}

// func divide(a int, b int) (int,error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("cannot divide by zero")
// 	}

// 	return a / b, nil
// }

// func main() {
// 	sum := add(3,4)
// 	fmt.Println("sum:",sum)
// 	quotient, err :=  divide(10,2)
// 	if err != nil {
// 		fmt.Println("Error:",err)
// 	}else {
// 		fmt.Println("Quotient:", quotient)
// 	}
// }

//array
// func main() {
// 	var arr [5]int
// 	arr[0] = 1
// 	arr[1] = 2
// 	fmt.Println(arr)
	 
// 	numbers :=  [3] int{1,2,3}
// 	fmt.Println(numbers)
// }

//slice

func main() {
	var slice []int
	slice = append(slice,1,2,3,4)
	fmt.Println(slice)

	a := []int{1,2,3,4,5}
	fmt.Println(a)

	s := make([]int,5)
	s = append(1,2,3)
	fmt.Println(s)

	fruit := []string{"apple","orange","banana"}
	fmt.Println(fruit)
}
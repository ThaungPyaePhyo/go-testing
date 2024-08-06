package main

import "fmt"

func add(a int,b int) int {
		return a + b
	}

func divide(a int, b int) (int,error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a / b, nil
}

func main() {
	sum := add(3,4)
	fmt.Println("sum:",sum)
	quotient, err :=  divide(10,2)
	if err != nil {
		fmt.Println("Error:",err)
	}else {
		fmt.Println("Quotient:", quotient)
	}
}
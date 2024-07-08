package main

import (
	"fmt"
)

func plus(a int,b int) int{
	return a + b	
}

func vals() (int, int) {
	return 1,2
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func sum(nums ...int) {
	fmt.Print(nums,"")
	total := 0
	for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {
	res := plus(1,2)
	fmt.Println("1 +2 =", res)

	res = plusPlus(1,2,3)
	fmt.Println("1+2+3 =", res)

	a, b := vals()
    fmt.Println(a,b)
	_, c := vals()
	fmt.Println(c)
	sum(1, 2)

	nums := []int{1, 2, 3, 4}
	numbers := []int{2,3,4,5}
    sum(nums...)
	sum(numbers...)

}
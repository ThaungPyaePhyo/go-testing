package main

import (
	"fmt"
)

func main() {
	if 7 % 2 == 0 {
		fmt.Println("remainder is soine")
	}else {
		fmt.Println("remainder is ma")
	}

	if num := 9; num <0 {
		fmt.Println(num,"is a negative value.")
	}else if num <10 {
		fmt.Println(num, "is a number.")
	}else {
		fmt.Println(num, "is various numbers.")
	}
}
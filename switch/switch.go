package main

import (
	"fmt"
	"time"
)

func main(){
	i := 1

	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend.")
	default:
		fmt.Println("It's weekday.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon.")
	default:
		fmt.Println("It's afternoon.")
	}


	whatIm := func(i interface{}) {
		switch t := i.(type){
		case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
		}
	} 

	whatIm(true)
	whatIm(1)
	whatIm("helo")
	
}
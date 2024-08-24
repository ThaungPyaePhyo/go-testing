package main

import (
	"fmt"
	"net/http"
)

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
// func main() {
// 	var slice []int
// 	slice = append(slice,1,2,3,4)
// 	fmt.Println(slice)

// 	a := []int{1,2,3,4,5}
// 	fmt.Println(a)

// 	s := make([]int,5)
// 	s = append(1,2,3)
// 	fmt.Println(s)

// 	fruit := []string{"apple","orange","banana"}
// 	fmt.Println(fruit)
// }

//Maps

// func main() {
// 	person := make(map[string]string)
// 	person["name"] = "alice"
// 	person["city"] = "new york"
// 	fmt.Println(person)

// 	countries := map[string]string{"Us": " united States", "CA": "Canada"}
// 	fmt.Println(countries)
// }

//Struct

// type Person struct {
// 	name string
// 	age int
// }
// func main() {
// 	var p Person

// 	p.name = "Alice"
// 	p.age = 30
// 	fmt.Println(p)

// 	p2 := Person{name: "bo bo", age: 20}
// 	fmt.Println(p2)
// }

// Methods
// type Person struct {
// 	name string
// 	age int
// }

// func (p Person) greet() string{
// 	return "Hello, my name is " + p.name
// }

// func main() {
// 	p := Person{name: "koko", age:80}
// 	fmt.Println(p.greet())
// }

// Pointer
// func main() {
// 	var x int = 10
// 	var p *int = &x
// 	fmt.Println("x:", x)
// 	fmt.Println("p:", p)
// 	fmt.Println("*p:", *p)

// 	*p = 20
//     fmt.Println("x after *p = 20:", x)

// }

func handler(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w, "hello")
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)

}
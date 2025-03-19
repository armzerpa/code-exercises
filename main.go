package main

import (
	"fmt"
)

type Person struct {
	name  string
	age   int
	phone string
}

func main() {
	fmt.Println("Hello world")

	p := Person{name: "armando", age: 11, phone: "5551235423"}
	fmt.Printf("Person: %+v\n", p)

	pmap := make(map[string]Person)
	pmap["111"] = Person{name: "armando", age: 11, phone: "5551235423"}
	pmap["222"] = Person{name: "sandra", age: 22, phone: "xxxxxx"}
	fmt.Printf("Person: %+v\n", pmap)

	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	for i := 0; i < len(arr); i++ {
		fmt.Println(i)
	}

	names := []string{"banana", "apple", "orange"}
	for index, value := range names {
		fmt.Printf("for with index %d and %v\n", index, value)
	}

	arr2 := []int{1, 3, 5, 2, 10, 7, 6, 6}
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	arr3 := [2]string{"hola", "hi"}
	for _, value := range arr3 {
		fmt.Println(fmt.Sprintf("this is a string %v", value))
	}

	m := map[string]int{"Jose": 20, "Sandra": 11, "Mati": 2}
	for i, v := range m {
		fmt.Printf("map is i: %v, v: %v\n", i, v)
	}

	m2 := make(map[string]int)
	m2["Armando"] = 22
	m2["Luis"] = 66
	fmt.Println(m2)

	val, exists := m2["otro"]
	if exists {
		fmt.Println(val)
	} else {
		fmt.Println("not exists")
	}

	delete(m2, "Luis")
	fmt.Println(m2)
}

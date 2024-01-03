package main

import "fmt"

func sequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}

func main() {

	int_seq := sequence()
	fmt.Println(int_seq()) //have reference of the variable i
	fmt.Println(int_seq()) //recalls the same function

	int_seq1 := sequence()
	fmt.Println(int_seq1())

}

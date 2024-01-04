package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("Error Encountered")

	if err != nil {
		fmt.Println(err)
	}

}

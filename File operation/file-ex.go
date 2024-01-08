//example to get the information of the file

package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, err := os.Stat("samplefile.txt") //stat is used to get the information
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("File size (in bytes):", fileInfo.Size())
	fmt.Println("File permissions:", fileInfo.Mode())
	fmt.Println("Is directory?", fileInfo.IsDir())
	fmt.Println("File modification time:", fileInfo.ModTime())
}
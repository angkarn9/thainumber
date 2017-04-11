package main

import (
	"fmt"

	"./thainumber"
)

func main() {
	fmt.Printf("Number %d is %s\n", 34543543, thainumber.ToThaiNumber(34543543))
}

package main

import (
	"fmt"
)

func main() {
	loop := 0
	in := 0

	fmt.Scanf("%d", &loop)

	input := []int{}
	for i := 0; i < loop; i++ {
		fmt.Scanf("%d", &in)
		input = append(input, in)
	}

	fmt.Println(input)

}

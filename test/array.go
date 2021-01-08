package main

import (
	"fmt"
	"github.com/a1049990866/go-array/array"
)

func main() {
	arr := array.NewArray([]int{1, 2, 3, 4}, array.IntArray)
	fmt.Println(arr.ToIntMap())
}

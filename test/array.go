package main

import (
	"fmt"
	"go-array/array"
)

func main() {
	arr := array.NewArray([]string{"1", "2", "3", "x", "3"}, array.StringArray).Deduplication(true)
	fmt.Println(arr.Data)
	//fmt.Println(arr.ToIntArray())
	fmt.Println(arr.ToFloat32Map())
}

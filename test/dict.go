package main

import (
	"fmt"
	"github.com/a1049990866/go-array/dict"
)

func main() {
	d := dict.NewDict(nil)
	d.Set("1", false, 2, "3", 4, 5)
	fmt.Println(d.Data, "----------", d.Value.Type())
	d.Set(2, true, 3, "4", 5)
	fmt.Println(d.Data, "----------", d.Value.Type())
	d.Set(3, false, 3, 5)
	fmt.Println(d.Data, "----------", d.Value.Type())
	d.Set(2, true, 3, "x")
	fmt.Println(d.Data, "----------", d.Value.Type())
	d.Set("2", true, "3", "xx")
	fmt.Println(d.Data, "----------", d.Value.Type())
	d.Set("3", true, 3, 4)
	fmt.Println(d.Data, "----------", d.Value.Type())
	fmt.Println(d.Get(3, 4, 0), d.Error)
}

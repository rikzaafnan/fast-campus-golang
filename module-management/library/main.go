package main

import (
	"fmt"
	"github.com/rikzaafnan/fast-campus-golang/module-management/library/pointer"
)

func main() {
	x := pointer.Of("Hello")

	fmt.Println(*x)
}

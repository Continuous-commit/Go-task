package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	fmt.Printf("%d\n", strings.Count(a, "1"))
}

package main

import (
	"fmt"

	"github.com/dzhichko/Contains_module/Contains_function"
)

func main() {
	var d, m string
	fmt.Scan(&d, &m)
	fmt.Println(Contains_function.Contains(d, m))
}

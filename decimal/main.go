package main

import (
	"fmt"
	"github.com/govalues/decimal"
)

func main() {

	a, _ := decimal.Parse("1.1")
	b, _ := decimal.Parse("1.2")
	c, _ := decimal.Parse("1.3")

	a, _ = a.Add(b)
	a, _ = a.Add(c)
	fmt.Println(a.String())

}

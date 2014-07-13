package main

import (
	"flag"
	"fmt"

	"github.com/ukanga/mymath"
)

var val = flag.Float64("val", 2, "Value to get square root value")

func main() {
	flag.Parse()
	fmt.Printf("Hello World!\nSqrt(%v) is %v\n", *val, mymath.Sqrt(*val))
}

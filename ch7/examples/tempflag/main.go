package main

import (
	"flag"
	"fmt"
	"gopl/ch7/examples/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}

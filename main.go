// gothrift project main.go
package main

import (
	"flag"
)

var (
	t int
)

func init() {
	flag.IntVar(&t, "t", 1, "please input -t=1 or -t=0")
}

func main() {
	flag.Parse()
}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int       // a = 0
	var b, c string // b = "", c = ""

	d := false      // d = false (boleano)
	e, f := "", 0.0 // f = "" (string), e = 0 (float)

	var (
		g int     // g = 0
		h float64 // h = 0.0
	)

	const i int = 1000 // i = 1000

	const (
		j string = "Const" // j = "Const"
		k bool   = true    // k = true
	)

	arrVar := []Var{a, b, c, d, e, f, g, h, i, j, k}
	printVar(arrVar)
}

type Var interface {
	any
}

func printVar(vars []Var) {
	for _, v := range vars {
		fmt.Printf("%v - %v\n", v, reflect.TypeOf(v))
	}
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println(10 + 5)
	fmt.Println(10 - 5)
	fmt.Println(10 * 5)
	fmt.Println(10 / 5)
	fmt.Println(10 % 5)
	fmt.Println(10 > 5)
	fmt.Println(10 < 5)
	fmt.Println(10 >= 5)
	fmt.Println(10 <= 5)
	fmt.Println(10 == 5)
	fmt.Println(10 != 5)

	var v, f bool

	//fmt.Println(f && f)
	fmt.Println(f && v)
	fmt.Println(v && f)
	//fmt.Println(v && v)
	//fmt.Println(f || f)
	fmt.Println(f || v)
	fmt.Println(v || f)
	//fmt.Println(v || v)
	fmt.Println(!v)
	fmt.Println(!f)

	var n int

	n++     // n = n + 1
	n += 10 //n = n + 10
	n--     // n = n - 1
	n -= 5  // n = n - 5
	n *= 5  // n = n * 5
	n /= 5  // n = n / 5
	n %= 5  // n = n % 5

	fmt.Println(n)
}

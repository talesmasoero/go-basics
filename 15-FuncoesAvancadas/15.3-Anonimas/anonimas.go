package main

import (
	"fmt"
	"strconv"
)

func main() {
	func() {
		fmt.Println("Sou uma função anônima simples!")
	}()

	str := "H3ll0, W0rld"
	code, nums := func(str string) ([]rune, []int) {
		var ucode []rune
		var nums []int
		for _, char := range str {
			if num, ok := strconv.Atoi(string(char)); ok == nil {
				ucode = append(ucode, char)
				nums = append(nums, num)
			}
		}
		return ucode, nums
	}(str)
	fmt.Println(code)
	fmt.Println(nums)
}

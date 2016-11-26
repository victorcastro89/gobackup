package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(concat("Victor ", "e Let"))
	fmt.Println(arr[2:])

}

func concat(str1 string, str2 string) string {
	length := len(str1) + len(str2)

	alloc := make([]byte, length)
	bl := 0
	bl += copy(alloc[bl:], str1)
	bl += copy(alloc[bl:], str2)
	return string(alloc[:length])
}

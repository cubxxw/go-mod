package main

import "fmt"

func Reverse[T string](s T) T {
	/*
		rune 和 byte 的区别：
		1. rune 是 int32 的别名，byte 是 uint8 的别名
		2. rune 用来表示 Unicode 字符，byte 用来表示 ASCII 字符
		3. rune 用来表示一个字符，byte 用来表示一个字节
		4. rune 用来表示一个 UTF-8 字符，byte 用来表示一个 ASCII 字符
	*/
	// 汉字占3个字节，英文占1个字节
	b := []rune(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return T(b)
}

func main() {
	fmt.Println("反转之前：", "hello")
	fmt.Println("反转之后：", Reverse("hello"))

	fmt.Println("反转之前：", "hello welcome to golang")
	fmt.Println("反转之后：", Reverse("hello welcome to golang"))
}

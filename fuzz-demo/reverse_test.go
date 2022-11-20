package main

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func test() {
	fmt.Println("这是fuzz-demo的reverse_test.go")
}

func TestReverse(t *testing.T) {
	testcases := []struct { //定义一个结构体数组
		in, want string //分别是输入和期望输出
	}{
		{"hello", "olleh"},
		{"hello welcome to golang", "gnalog ot emoclew olleh"},
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"abc!", "!cba"},
		{"abc!def12321fedcba", "abcdef12321fed!cba"},
		// {"我是中国人", "人中国是我"},
	}
	for _, c := range testcases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func FuzzReverse(f *testing.F) {
	// ...这是模糊测试的代码
	// 添加测试集
	testcases := []string{"hello", "hello world", "hello welcome to golang", " ", "123124"}
	for _, c := range testcases {
		//添加测试集
		f.Add(c)
	}

	f.Fuzz(func(t *testing.T, a string) {
		got := Reverse(a)
		inRevers := Reverse(got) //反转两次，应该和原来的一样(a == inRevers)
		if got == inRevers {
			t.Errorf("reverse(%q) == %q, want %q", a, got, a)
		}
		if utf8.ValidString(a) && !utf8.ValidString(got) {
			t.Errorf("Reverse produced invalid UTF-8 string: %q", got)
		}
	})
}

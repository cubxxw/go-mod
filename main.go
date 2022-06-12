package main

import (
	. "fmt"
	"go-mod/hello"
	"go-mod/models"
)

func main() {
	Println("python")
	Println("main主函数")
	hello.Hello()
	Println(models.Name)
	//hello.Hello()
}

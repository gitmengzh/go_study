package main

import (
	"fmt"
)

func main() {

	hellowrod()
	variable()
	arr1()
}

func hellowrod() { // first go
	fmt.Println("Hello world!")
}

func variable() { //变量使用

	var age int //变量声明
	fmt.Println("my age is", age)
	age = 10 //赋值
	fmt.Println("Tom age is", age)
	var age2 int = 30 // 声明变量并赋值
	fmt.Println("Lily age is", age2)
	var age3 = 25 // 推断类型
	fmt.Println("tomy age is", age3)
	var name, age4 = "TOO", 18 // 多变量声明赋值
	fmt.Println(name, "age is", age4)
	name1 := "SONY"
	fmt.Println(name1, "TV")

}

func arr1() {
	var arr1 [5]int
	for i := 1; i < len(arr1); i++ {
		arr1[i] = i * 2
	}
	fmt.Println(arr1) //[0 2 4 6 8]

}

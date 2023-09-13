package main

import "fmt"

// alias to "string" type
type MyString string

func (s MyString) Length() int {
	return len(s)
}

func main() {
	str := MyString("Incididunt nostrud ad magna sit pariatur mollit magna proident. Proident qui minim tempor consectetur irure sit ullamco sit. Nostrud adipisicing ut ullamco fugiat laborum laborum qui cupidatat consectetur labore. Adipisicing occaecat excepteur pariatur cupidatat aliqua deserunt cupidatat et nisi ex amet quis. Laborum aliquip non aute occaecat culpa fugiat consequat veniam. Labore dolor ex aute culpa Lorem ipsum dolor elit in. Proident id Lorem elit adipisicing ut consectetur sunt.")
	fmt.Println(str.Length())
}

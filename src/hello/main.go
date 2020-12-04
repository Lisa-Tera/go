package main

import (
	"test"
	"test/korean"
)
func main(){ // 외부 참조를 위해 함수의 첫 글자는 대문자로
	test.Aaa()
	test.Bbb()
	korean.Ga()
	korean.Na()
}
//여기서는 test써야함
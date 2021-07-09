package main

import (
	"fmt"
	"strings"
)

type upperstring string

func (s upperstring) Upper() string {
	return strings.ToUpper(string(s))

}
func main() {
	s := upperstring("hello world")
	fmt.Println(s.Upper())
}

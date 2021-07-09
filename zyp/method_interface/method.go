package main

import "fmt"

type trangle struct {
	size int
}
type square struct {
	size int
}
type coloredTring struct {
	trangle
	color string
}

func (s square) perimeter() int {
	return s.size * 4
}
func (t trangle) perimeter() int {
	return t.size * 3
}
func (t *trangle) doubleSize() {
	t.size *= 2
}
func (t coloredTring) perimeter() int {
	return t.size * 3 * 2
}
func main() {
	t := trangle{3}
	s := square{3}
	t1 := trangle{3}
	t1.doubleSize()
	fmt.Println("permi", t1.perimeter())
	fmt.Println("perimeter:", t.perimeter())
	fmt.Println("perimeter:", s.perimeter())

	c:=coloredTring{trangle{3},"red"}
	fmt.Println(c.perimeter())
	fmt.Println(c.trangle.perimeter())
}

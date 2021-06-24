package main

import (
	"fmt"
	"log"
)

func main() {
	log.Print("log!")
	//panic("return")
	log.Fatal("Hey, I'm an error log!")
	fmt.Print("can you see me?")
}

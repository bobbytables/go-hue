package main

import (
	"fmt"
	"hue"
)

func main() {
	bridge := NewBridge()
	fmt.Println(fmt.Sprintf("%s", bridge.IP))
}

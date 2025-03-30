package main

import "fmt"

func testRender(sc *Screen) {
	fmt.Println("Rendering 'A' at 0,0")
	sc.PaintFg(0, 0, "A")
}

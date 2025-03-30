package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type Screen struct {
	// x, y | example = [500, 200]
	size [2]int
	fg   [][]string
	bg   [][]string
	loop bool
}

var EmptySpaces = []string{" ", " ", " ", ".", ",", " ", " ", "'", " ", "`", " ", " "}

func getEmptySpace() string {
	i := rand.Intn(len(EmptySpaces))
	return EmptySpaces[i]
}

func NewScreen(x, y int) *Screen {
	s := Screen{
		loop: false,
		size: [2]int{x, y},
		fg:   make([][]string, y),
		bg:   make([][]string, y),
	}
	// y = rows, x = columns
	emptyRow := func() []string {
		s := make([]string, x)
		for i := 0; i < len(s); i++ {
			s[i] = " "
		}
		return s
	}
	for i := 0; i < len(s.bg); i++ {
		s.bg[i] = emptyRow()
		s.fg[i] = emptyRow()
	}
	return &s
}

func (s Screen) SizeY() int {
	return s.size[1]
}

func (s Screen) SizeX() int {
	return s.size[0]
}

func (s *Screen) PaintFg(x, y int, v string) {
	if y >= s.SizeY() {
		log.Fatalf("Y is out of screen. Max y size: %d, got: %d\n", s.SizeY(), y)
	}
	if x >= s.SizeX() {
		log.Fatalf("X is out of screen. Max x size: %d, got: %d\n", s.SizeX(), x)
	}
	s.fg[y][x] = v
}

func (s *Screen) PainAsciiArt(art string) {
	y := 0
	lastX := 0
	for x := 0; x < len(eidMubarak2025); x++ {
		v := string(eidMubarak2025[x])
		switch v {
		case "\n":
			y++
			lastX = x + 1 // +1 because the next x value should be 0 to indicate the starting index.
			break

		case " ":
			break

		default:
			s.PaintFg(x-lastX, y, v)
			break
		}
	}
}

func (s Screen) ClearFg(x, y int) {
	s.PaintFg(x, y, " ")
}

func (s Screen) ClearOutputScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}

func (s *Screen) GetPixel(x, y int) string {
	if s.fg[y][x] != " " {
		return s.fg[y][x]
	}
	return getEmptySpace()
}

func (s Screen) Render() {
	out := ""
	for y, row := range s.bg {
		for x := range row {
			out += s.GetPixel(x, y)
		}
		out += "\n"
	}
	fmt.Printf(out)
}

func (s *Screen) Start() {
	s.loop = true
	s.ClearOutputScreen()
	s.Render()
	go func() {
		for s.loop {
			time.Sleep(time.Millisecond * 150)
			s.Render()
		}
	}()
}

func (s *Screen) Stop() {
	s.loop = false
}

func main() {
	sc := NewScreen(180, 40)
	sc.Start()
	go sc.PainAsciiArt(eidMubarak2025)
	// go eidMubarakRenderer(sc)
	time.Sleep(time.Second * 28)
	sc.Stop()
}

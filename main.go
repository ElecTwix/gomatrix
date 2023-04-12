package main

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

var matrix Matrix

var charset = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type Matrix struct {
	height int
	width  int
	matrix [][]rune
	Drops  []Drop
}

type Drop struct {
	x, y  int
	tail  int
	chars []rune
}

func main() {

	width, height, err := terminal.GetSize(0)

	if err != nil {
		panic(err)
	}

	matrix = Matrix{
		height: height,
		width:  width,
		Drops:  make([]Drop, 35),
	}

	matrix.matrix = make([][]rune, matrix.height)

	for i := 0; i < matrix.height; i++ {
		matrix.matrix[i] = make([]rune, matrix.width)
	}

	rand.Seed(time.Now().UnixNano())

	matrix.CreateDrops()

	for {
		matrix.drawMatrix()
		time.Sleep(50 * time.Millisecond)
		matrix.updateMatrix()
	}
}

func (m *Matrix) drawMatrix() {

	// Clear screen
	fmt.Print("\033[2J")

	for _, v := range m.matrix {
		fmt.Println(string(v))
	}

}

func (m *Matrix) updateMatrix() {

	// Clear matrix
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			m.matrix[y][x] = ' '
		}
	}

	// Update matrix
	for id, drop := range m.Drops {
		m.Drops[id].chars[len(m.Drops[id].chars)-1] = charset[rand.Intn(len(charset))]
		for i, v := range drop.chars {
			if i+drop.y >= m.height {
				m.matrix[i+drop.y-m.height][drop.x] = v
			} else {
				m.matrix[drop.y+i][drop.x] = v
			}
		}
		m.Drops[id].y++
		if drop.y >= m.height {
			m.Drops[id].y = 0
		}
	}
}

func (m *Matrix) CreateDrops() {
	for i := range m.Drops {

		size := rand.Intn(20) + 4

		m.Drops[i] = Drop{
			x:     rand.Intn(m.width),
			y:     rand.Intn(m.height),
			tail:  size,
			chars: make([]rune, size),
		}

		for i2 := 0; i2 < size; i2++ {
			m.Drops[i].chars[i2] = charset[rand.Intn(len(charset))]
		}
	}
}

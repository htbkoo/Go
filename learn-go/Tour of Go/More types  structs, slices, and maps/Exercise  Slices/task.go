package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	/* TODO */

	//p := make([][]uint8, dy)
	//for i, _ := range p {
	//	p[i] = make([]uint8, dx)
	//	for j, _ := range p[i] {
	//		p[i][j] = uint8((i + j) / 2)
	//		fmt.Printf("%d %d %d\n", i, j, p[i][j])
	//	}
	//}
	//
	//return p

	p := make([][]uint8, dy)
	for i, _ := range p {
		p[i] = make([]uint8, dx)
		for j, _ := range p[i] {
			p[i][j] = uint8(i * j)
		}
	}

	return p

	//p := make([][]uint8, dy)
	//for i, _ := range p {
	//	p[i] = make([]uint8, dx)
	//}
	//
	//for i, row := range p {
	//	for j := range row {
	//		row[j] = uint8(i * j)
	//	}
	//}
	//
	//return p

	//p := make([][]uint8, dy)
	//for i := range p {
	//	p[i] = make([]uint8, dx)
	//}
	//
	//for y, row := range p {
	//	for x := range row {
	//		row[x] = uint8(x * y)
	//	}
	//}
	//
	//return p
}

func main() {
	pic.Show(Pic)
}

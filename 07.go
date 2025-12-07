package main
import "fmt"

func main() {
	type pos struct { x, y int }
	var beams, splitters = make(map[pos]int64), make(map[pos]bool)
	var splits, start, y = 0, pos{ 0, 0 }, 0
	for {
		var line string
		if _, err := fmt.Scanln(&line); err != nil { break }
		for x, c := range line {
			if c == 'S' { start = pos{ x, 0 } }
			if c == '^' { splitters[pos{ x, y }] = true }
		}
		y++
	}
	beams[start] = 1
	out:
	for {
		var step = make(map[pos]int64)
		for b, tl := range beams {
			b.y++
			if splitters[b] {
				step[pos{ b.x - 1, b.y }] = step[pos{ b.x - 1, b.y }] + tl
				step[pos{ b.x + 1, b.y }] = step[pos{ b.x + 1, b.y }] + tl
				splits++
			} else {
				step[b] = step[b] + tl
			}
		}
		beams = step
		for b, _ := range beams {
			if b.y == y { break out }
		}
	}
	var timelines = int64(0)
	for _, tl := range beams {
		timelines += tl
	}
	fmt.Println(splits ,timelines)
}

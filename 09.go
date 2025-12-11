package main
import "fmt"
import "slices"

type pos struct { x, y int }
type dir struct { dx, dy int }
var up, right, down, left = dir{ 0, -1 }, dir{ 1, 0 }, dir{ 0, 1 }, dir{ -1, 0 }

func main() {
	var reds = make([]pos, 0)
	for {
		var x, y int
		if _, err := fmt.Scanf("%d,%d", &x, &y); err != nil { break }
		reds = append(reds, pos{ x, y })
	}
	var maxareas [2]int64
	var xs = make([]int, len(reds))
	var ys = make([]int, len(reds))
	for i := 0; i < len(reds); i++ {
		for j := i + 1; j < len(reds); j++ {
			maxareas[0] = max(maxareas[0], area(reds[i], reds[j]))
		}
		xs[i] = reds[i].x
		ys[i] = reds[i].y
	}
	slices.Sort(xs)
	xs = dedup(xs)
	slices.Sort(ys)
	ys = dedup(ys)
	var virt = make([]pos, len(reds))
	copy(virt, reds)
	var maxx, maxy = 1, 1
	for i := 0; i < len(xs); i++ {
		for j, _ := range virt {
			if virt[j].x == xs[i] {
				virt[j].x = maxx
			}
		}
		maxx += 2
	}
	for i := 0; i < len(ys); i++ {
		for j, _ := range virt {
			if virt[j].y == ys[i] {
				virt[j].y = maxy
			}
		}
		maxy += 2
	}
	var colored = make(map[pos]bool)
	for i := 0; i < len(virt); i++ {
		var ra, rb = virt[i], virt[(i + 1) % len(virt)]
		if ra.x == rb.x {
			for y := min(ra.y, rb.y); y <= max(ra.y, rb.y); y++ {
				colored[pos{ra.x, y}] = true
			}
		} else {
			for x := min(ra.x, rb.x); x <= max(ra.x, rb.x); x++ {
				colored[pos{x, ra.y}] = true
			}
		}
	}
	var outside = make(map[pos]bool)
	var q = make([]pos, 0)
	q = append(q, pos{ 0, 0 })
	for len(q) > 0 {
		var p = q[0]
		q = q[1:]
		if outside[p] { continue }
		outside[p] = true
		for _, d := range []dir{ up, right, down, left } {
			var nx = pos{ p.x + d.dx, p.y + d.dy }
			if 0 <= nx.x && nx.x < maxx + 2 && 0 <= nx.y && nx.y < maxy + 2 && !colored[nx] {
				q = append(q, nx)
			}
		}
	}
	var indexes [2]int
	for i := 0; i < len(virt) - 1; i++ {
		for j := i + 1; j < len(virt); j++ {
			var a, b = virt[i], virt[j]
			var inside = true
			for x := min(a.x, b.x); x <= max(a.x, b.x) && inside; x++ {
				inside = !outside[pos{ x, min(a.y, b.y) }] && !outside[pos{ x, max(a.y, b.y) }]
			}
			if !inside { continue }
			for y := min(a.y, b.y); y <= max(a.y, b.y) && inside; y++ {
				inside = !outside[pos{ min(a.x, b.x), y }] && !outside[pos{ max(a.x, b.x), y }]
			}
			if !inside { continue }
			if maxareas[1] < area(a, b) {
				maxareas[1], indexes[0], indexes[1] = area(a, b), i, j
			}
		}
	}
	fmt.Println(maxareas[0], area(reds[indexes[0]], reds[indexes[1]]))
}

func abs(i int) int {
	if i < 0 { return -i }
	return i
}

func area(a, b pos) int64 {
	return int64(abs(a.x - b.x) + 1) * int64(abs(a.y - b.y) + 1)
}

func dedup(s []int) []int {
	var res = s
	for i := 1; i < len(res); i++ {
		if res[i - 1] == res[i] {
			res = append(res[:i], res[i + 1:]...)
		}
	}
	return res
}
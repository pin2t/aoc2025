package main
import "fmt"
import "sort"

func main() {
	var boxes = make([]pos, 0)
	for {
		var x, y, z int
		if _, err := fmt.Scanf("%d,%d,%d", &x, &y, &z); err != nil { break }
		boxes = append(boxes, pos{ x, y, z })
	}
	type conn struct {
		a, b pos
		dist int64
	}
	var conns = make([]conn, 0)
	for i := 0; i < len(boxes) - 1; i++ {
		for j := i + 1; j < len(boxes); j++ {
			conns = append(conns, conn{ boxes[i], boxes[j], dist(boxes[i], boxes[j]) })
		}
	}
	sort.Slice(conns, func(i, j int) bool {
		return conns[i].dist < conns[j].dist
	})
	var nconn, res1 = 0, 0
	var circuits = make([]map[pos]bool, 0)
	for _, cn := range conns {
		var ca, cb = -1, -1
		for i, c := range circuits {
			if c[cn.a] { ca = i }
			if c[cn.b] { cb = i }
		}
		if ca == -1 {
			if cb == -1 {
				var c = make(map[pos]bool)
				c[cn.a] = true
				c[cn.b] = true
				circuits = append(circuits, c)
			} else {
				circuits[cb][cn.a] = true
			}
		} else {
			if cb == -1 {
				circuits[ca][cn.b] = true
			} else if cb != ca {
				for b, _ := range circuits[cb] {
					circuits[ca][b] = true
				}
				circuits[cb] = circuits[len(circuits) - 1]
				circuits = circuits[:len(circuits) - 1]
			}
		}
		nconn++
		if nconn == 1000 {
			var sizes = make([]int, 0)
			for _, c := range circuits {
				sizes = append(sizes, len(c))
			}
			sort.Slice(sizes, func (i, j int) bool {
				return sizes[i] > sizes[j]
			})
			res1 = sizes[0] * sizes[1] * sizes[2]
		}
		if len(circuits[0]) == len(boxes) {
			fmt.Println(res1, cn.a.x * cn.b.x)
			break
		}
	}
}

type pos struct { x, y, z int }

func dist(a, b pos) int64 {
	return int64((b.x - a.x) * (b.x - a.x) + (b.y - a.y) * (b.y - a.y) + (b.z - a.z) * (b.z - a.z))
}

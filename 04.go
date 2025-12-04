package main
import "bufio"
import "os"
import "fmt"

var grid = make([][]byte, 0)

func roll(r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[r]) && grid[r][c] == '@'
}

func adj(r, c int) int {
	var cnt = 0
	if roll(r - 1, c) { cnt++ }
	if roll(r + 1, c) { cnt++ }
	if roll(r, c - 1) { cnt++ }
	if roll(r, c + 1) { cnt++ }
	if roll(r - 1, c - 1) { cnt++ }
	if roll(r + 1, c + 1) { cnt++ }
	if roll(r + 1, c - 1) { cnt++ }
	if roll(r - 1, c + 1) { cnt++ }
	return cnt
}

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	var accesible = 0
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	for r, row := range grid {
		for c, _ := range row {
			if roll(r, c) && adj(r, c) < 4 { accesible++ }
		}
	}
	var total = 0
	for {
		var removed = make([][]byte, len(grid))
		for r, row := range grid {
			var cp = make([]byte, len(row))
			copy(cp, row)
			removed[r] = cp
		}
		var nrem = 0
		for r, row := range grid {
			for c, _ := range row {
				if roll(r, c) && adj(r, c) < 4 {
					nrem++
					removed[r][c] = '.'
				}
			}
		}
		total += nrem
		if nrem == 0 { break }
		grid = removed
	}
	fmt.Println(accesible, total)
}

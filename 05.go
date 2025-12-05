package main
import "bufio"
import "os"
import "strings"
import "strconv"
import "fmt"
import "sort"

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	var ranges = make([][]int64, 0)
	for scanner.Scan() {
		var line = scanner.Text()
		if line == "" { break }
		var parts = strings.Split(line, "-")
		var start, _ = strconv.ParseInt(parts[0], 10, 64)
		var end, _ = strconv.ParseInt(parts[1], 10, 64)
		ranges = append(ranges, []int64{ start, end })
	}
	var n, total = 0, int64(0)
	for scanner.Scan() {
		var id, _ = strconv.ParseInt(scanner.Text(), 10, 64)
		for _, r := range ranges {
			if r[0] <= id && id <= r[1] {
				n++
				break
			}
		}
	}
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	var merged = make([][]int64, 1)
	merged[0] = ranges[0]
	for i := 1; i < len(ranges); i++ {
		if ranges[i][0] <= merged[len(merged) - 1][1] {
			merged[len(merged) - 1][1] = max(merged[len(merged) - 1][1], ranges[i][1])
		} else {
			merged = append(merged, ranges[i])
		}
	}
	for _, r := range merged {
		total += r[1] - r[0] + 1
	}
	fmt.Println(n, total)
}

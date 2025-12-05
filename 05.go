package main
import "bufio"
import "os"
import "strings"
import "strconv"
import "fmt"
import "sort"

type rng struct { first, last int64 }

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	var ranges = make([]rng, 0)
	for scanner.Scan() {
		var line = scanner.Text()
		if line == "" { break }
		var pair = strings.Split(line, "-")
		var r rng
		r.first, _ = strconv.ParseInt(pair[0], 10, 64)
		r.last, _ = strconv.ParseInt(pair[1], 10, 64)
		ranges = append(ranges, r)
	}
	var fresh, total = 0, int64(0)
	for scanner.Scan() {
		var id, _ = strconv.ParseInt(scanner.Text(), 10, 64)
		for _, r := range ranges {
			if r.first <= id && id <= r.last {
				fresh++
				break
			}
		}
	}
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].first < ranges[j].first
	})
	var merged = make([]rng, 1)
	merged[0] = ranges[0]
	for i := 1; i < len(ranges); i++ {
		var l = &merged[len(merged) - 1]
		if ranges[i].first <= l.last {
			l.last = max(l.last, ranges[i].last)
		} else {
			merged = append(merged, ranges[i])
		}
	}
	for _, r := range merged {
		total += r.last - r.first + 1
	}
	fmt.Println(fresh, total)
}

package main

import "fmt"
import "strings"
import "strconv"

func main() {
	var line string
	_, _ = fmt.Scanln(&line)
	var ranges = strings.Split(line, ",")
	var ids, ids2 int64
	for _, rang := range ranges {
		var se = strings.Split(rang, "-")
		if len(se) != 2 {
			continue
		}
		var start, end int64
		start, _ = strconv.ParseInt(se[0], 10, 64)
		end, _ = strconv.ParseInt(se[1], 10, 64)
		for n := start; n <= end; n++ {
			var s = strconv.FormatInt(n, 10)
			if len(s)%2 == 0 && s[:len(s)/2] == s[len(s)/2:] {
				ids += n
			}
			for i, j := 1, 0; i <= len(s)/2; i++ {
				if len(s)%i != 0 {
					continue
				}
				for j = 0; j < len(s) && s[j:j+i] == s[:i]; j += i {
				}
				if j == len(s) {
					ids2 += n
					break
				}
			}
		}
	}
	fmt.Println(ids, ids2)
}

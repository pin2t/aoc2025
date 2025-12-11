package main
import	"bufio"
import "os"
import "regexp"
import "fmt"

var conns = make(map[string][]string)
var cache = make(map[string]int64)

func paths(path []string, to string) (result int64) {
	var cur = path[len(path) - 1]
 	if cur == to {
 		return 1
 	}
 	if c, found := cache[cur + to]; found {
 		return c
 	}
 	var np = make([]string, len(path) + 1)
 	copy(np, path)
 	for _, out := range conns[cur] {
 		np[len(np) - 1] = out
 		result += paths(np, to)
 	}
 	cache[cur + to] = result
 	return result
 }

 func main() {
	var reName = regexp.MustCompile("[a-z]*")
	var scanner = bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
		var line = scanner.Text()
		var parts = reName.FindAllString(line, -1)
		var outs = make([]string, 0)
		for i := 2; i < len(parts); i++ {
			outs = append(outs, parts[i])
		}
		conns[parts[0]] = outs
	}
	fmt.Println(paths([]string{ "you" }, "out"),
		paths([]string{ "svr" }, "fft") * paths([]string{ "fft" }, "dac") * paths([]string{ "dac" }, "out"))
}

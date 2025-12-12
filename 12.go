package main
import "bufio"
import "os"
import "strings"
import "regexp"
import "strconv"
import "fmt"

func toint(s string) int {
	n64, _ := strconv.ParseInt(s, 10, 64)
	return int(n64)
}

func main() {
	var tiles = []int{ 5, 7, 7, 7, 6, 7 }
	var reNum = regexp.MustCompile("\\d+")
	scanner := bufio.NewScanner(os.Stdin)
	var res = 0
	for scanner.Scan() {
		var line = scanner.Text()
		if !strings.ContainsRune(line, 'x') { continue }
		var sn = reNum.FindAllString(line, -1)
		var area = toint(sn[0]) * toint(sn[1])
		var st = 0
		for i := 2; i < len(sn); i++ {
			st += toint(sn[i]) * tiles[i - 2]
		}
		if st <= area {
			res++
		}
	}
	fmt.Println(res)
}
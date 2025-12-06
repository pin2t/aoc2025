package main
import "bufio"
import "os"
import "regexp"
import "fmt"
import "strconv"

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	var reNum = regexp.MustCompile(`\d+`)
	var lines = make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	var ops = lines[len(lines) - 1]
	var res int64
	for j := 0; j < len(ops); j ++ {
		var op = ops[j]
		if op == '+' || op == '*' {
			var numbers = make([]int64, 0)
			for i := 0; i < len(lines) - 1; i ++ {
				numbers = append(numbers, toint(reNum.FindString(lines[i][j:])))
			}
			res += eval(op, numbers)
		}
	}
	var op byte
	var res2 int64
	var vn = make([]int64, 0)
	for j := 0; j < len(lines[0]); j++ {
		if j < len(ops) && ops[j] != ' ' {
			op = ops[j]
		}
		var s string
		for i := 0; i < len(lines) - 1; i++ {
			if lines[i][j] != ' ' {
				s += string(lines[i][j])
			}
		}
		if len(s) > 0 {
			vn = append(vn, toint(s))
		} else {
			res2 += eval(op, vn)
			vn = make([]int64, 0)
		}
	}
	res2 += eval(op, vn)
	fmt.Println(res, res2)
}

func eval(op byte, numbers []int64) (res int64) {
	if op == '*' { res = 1 }
	for _, n := range numbers {
		switch op {
		case '+': res += n
		case '*': res *= n
		}
	}
	return
}

func toint(s string) (n int64) {
	n, _ = strconv.ParseInt(s, 10, 64)
	return
}

package main
import "fmt"
import "regexp"
import "strconv"
import "bufio"
import "os"
import "slices"
import z3 "github.com/mitchellh/go-z3"

func main() {
	type machine struct {
		lights []bool
		buttons [][]int
		joltage []int
	}
	var machines = make([]machine, 0)
	var reWire = regexp.MustCompile("\\(.*?\\)")
	var reParts = regexp.MustCompile("^\\[([.#]+)] ([()\\d, ]+) \\{([\\d,]+)}$")
	var reNum = regexp.MustCompile("\\d+")
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var line = scanner.Text()
		var parts = reParts.FindStringSubmatch(line)
		var lights = make([]bool, len(parts[1]))
		for i, c := range parts[1] {
			lights[i] = c == '#'
		}
		var bs = reWire.FindAllString(parts[2], -1)
		var buttons = make([][]int, len(bs))
		for i, b := range bs {
			var ns = reNum.FindAllString(b, -1)
			buttons[i] = make([]int, len(ns))
			for j, n := range ns {
				buttons[i][j] = toint(n)
			}
		}
		var ns = reNum.FindAllString(parts[len(parts) - 1], -1)
		var joltage = make([]int, len(ns))
		for i, n := range ns {
			joltage[i] = toint(n)
		}
		machines = append(machines, machine { lights, buttons, joltage })
	}
	var presses = []int{ 0, 0 }
	for _, m := range machines {
		type state struct {
			lights []bool
			presses int
		}
		var init = make([]bool, len(m.lights))
		for i, _ := range init {
			init[i] = false
		}
		var q = make([]state, 0)
		q = append(q, state{ init, 0 })
		for len(q) > 0 {
			var st = q[0]
			q = q[1:]
			if slices.Equal(m.lights, st.lights) {
				presses[0] += st.presses
				break
			}
			for _, b := range m.buttons {
				var lights = make([]bool, len(st.lights))
				copy(lights, st.lights)
				for _, bl := range b {
					lights[bl] = !lights[bl]
				}
				q = append(q, state{ lights, st.presses + 1 })
			}
		}
	}
	var config = z3.NewConfig()
	var ctx = z3.NewContext(config)
	config.Close()
	defer ctx.Close()
	for _, m := range machines {
		var s = ctx.NewSolver()
		defer s.Close()
		var zero = ctx.Int(0, ctx.IntSort())
		var vars = make([]*z3.AST, len(m.buttons))
		for i, _ := range m.buttons {
			vars[i] = ctx.Const(ctx.Symbol(fmt.Sprintf("n%d", i)), ctx.IntSort())
			s.Assert(vars[i].Gt(zero))
		}
		for i, j := range m.joltage {
			var equation = ctx.Int(0, ctx.IntSort())
			for _, b := range m.buttons {
				var found = false
				for _, out := range b {
					if out == i {
						found = true
						break
					}
				}
				if found {
					equation.Add(vars[j])
				}
			}
			equation.Eq(ctx.Int(j, ctx.IntSort()))
			s.Assert(equation)
		}
		if v := s.Check(); v != z3.True {
			fmt.Println("Unsolvable", m)
		} else {
			var model = s.Model()
			var assignments = model.Assignments()
			model.Close()
			var pr = 0
			for _, a := range assignments {
				pr += a.Int()
			}
			presses[1] += pr
		}
	}
	fmt.Println(presses)
}

func toint(s string) int {
	n64, _ := strconv.ParseInt(s, 10, 64)
	return int(n64)
}

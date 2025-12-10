package main
import "fmt"
import "regexp"
import "strings"
import "strconv"
import "bufio"
import "os"
import "slices"

func main() {
	type machine struct {
		lights []bool
		buttons [][]int
		joltage []int
	}
	var machines = make([]machine, 0)
	var reWire = regexp.MustCompile("\\(.*?\\)")
	var reJolts = regexp.MustCompile("{.*?}")
	var reNum = regexp.MustCompile("\\d+")
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var line = scanner.Text()
		var l = line[strings.Index(line, "[") + 1:strings.LastIndexByte(line, ']')];
		var lights = make([]bool, len(l))
		for i, c := range l { lights[i] = c == '#' }
		var bs = reWire.FindAllString(line, -1)
		var buttons = make([][]int, len(bs))
		for i, b := range bs {
			var ns = reNum.FindAllString(b, -1)
			buttons[i] = make([]int, len(ns))
			for j, n := range ns { buttons[i][j] = toint(n) }
		}
		var js = reJolts.FindString(line)
		var ns = reNum.FindAllString(js, -1)
		var joltage = make([]int, len(ns))
		for i, n := range ns { joltage[i] = toint(n) }
		machines = append(machines, machine { lights, buttons, joltage })
	}
	var presses = []int{ 0, 0 }
	for _, m := range machines {
		type state struct {
			lights []bool
			presses int
		}
		var inilights = make([]bool, len(m.lights))
		for i, _ := range inilights {
			inilights[i] = false
		}
		var queue = make([]state, 0)
		queue = append(queue, state{ inilights, 0 })
		for len(queue) > 0 {
			var st = queue[0]
			queue = queue[1:]
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
				queue = append(queue, state{ lights, st.presses + 1 })
			}
		}
	}
	//for _, m := range machines {
	//	type state struct {
	//		jlts []int
	//		presses int
	//	}
	//	var init = make([]int, len(m.joltage))
	//	for i, _ := range init {
	//		init[i] = 0
	//	}
	//	var queue = make([]state, 0)
	//	queue = append(queue, state{ init, 0 })
	//	for len(queue) > 0 {
	//		var st = queue[0]
	//		queue = queue[1:]
	//		fmt.Println(st.jlts, m.joltage, st.presses)
	//		if slices.Equal(m.joltage, st.jlts) {
	//			presses[1] += st.presses
	//			break
	//		}
	//		for _, b := range m.buttons {
	//			var stt = st
	//			var exceed = false
	//			for !exceed {
	//				var jlts = make([]int, len(stt.jlts))
	//				copy(jlts, stt.jlts)
	//				for _, ji := range b {
	//					jlts[ji]++
	//					exceed = jlts[ji] > m.joltage[ji]
	//				}
	//				if !exceed {
	//					queue = append(queue, state{ jlts, stt.presses + 1 })
	//					stt = state{ jlts, stt.presses + 1 }
	//				}
	//			}
	//		}
	//	}
	//}
	fmt.Println(presses)
}

func toint(s string) int {
	n64, _ := strconv.ParseInt(s, 10, 64)
	return int(n64)
}

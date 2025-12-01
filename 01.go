package main

import "fmt"

func main() {
	var n, count, clicks = 50, 0, 0
	for {
		var turn, step = 0, 1
		var dir byte
		if _, err := fmt.Scanf("%c%d", &dir, &turn); err != nil { break }
		if dir == 'L' { step = 99 }
		for range turn {
			n = (n + step) % 100
			if n == 0 { clicks++ }
		}
		if n == 0 { count++ }
	}
	fmt.Println(count, clicks)
}

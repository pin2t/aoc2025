package main

import "bufio"
import "fmt"
import "os"
import "strconv"

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	var n, count, clicks = 50, 0, 0
	for scanner.Scan() {
		var line = scanner.Text()
		c, _ := strconv.Atoi(line[1:])
		switch (line[0]) {
		case 'L': 
			for range c {
				n = (n - 1 + 100) % 100
				if n == 0 { clicks++ }
			}
		case 'R': 
			for range c {
				n = (n + 1) % 100
				if n == 0 { clicks++ }
			}
		}	
		if n == 0 { count++ }
	}
	fmt.Println(count, clicks)
}

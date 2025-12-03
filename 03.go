package main
import "fmt"

func main() {
	var jolts = [2]int64{ 0, 0 }
	for {
		var bank string
		if _, err := fmt.Scanln(&bank); err != nil { break }
		var m = int64(0)
		for i := 0; i < len(bank) - 1; i++ {
			for j := i + 1; j < len(bank); j++ {
				m = max(m, int64((bank[i]-'0') * 10 + (bank[j]-'0')))
			}
		}
		jolts[0] += m
		m = 0
		var dm = -1
		for i := 1; i <= 12; i++ {
			dm++
			for j := dm + 1; j < len(bank) - 12 + i; j++ {
				if bank[j] > bank[dm] {
					dm = j
				}
			}
			m = m * 10 + int64(bank[dm] - '0')
		}
		jolts[1] += m
	}
	fmt.Println(jolts)
}
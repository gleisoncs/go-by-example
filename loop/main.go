package loop

import "fmt"

const (
	limit = 10000000000
)

func serialSum() int {
	sum := 0
	for i := 0; i < limit; i++ {
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(serialSum())
}

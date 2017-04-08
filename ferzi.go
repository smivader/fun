package main

import "fmt"

var result chan []int

func generate(a []int, n int, N int) {
	if n == 0 {
		result <- append([]int{}, a...)
		return
	}

	for i := 1; i <= N; i++ {
		duplicate := false

		for _, x := range a {
			if x == i {
				duplicate = true
				break
			}
		}

		if duplicate {
			continue
		}

		generate(append(a, i), n-1, N)
	}
}

func небьют(a []int) bool {
	for i := range a {
		for j := range a {
			if i == j {
				continue
			}

			if a[i] == a[j]-j+i || a[i] == a[j]+j-i {
				return false
			}
		}
	}

	return true
}

func main() {
	result = make(chan []int)

	go func() {
		generate([]int{}, 8, 8)
		close(result)
	}()

	for s := range result {
		if небьют(s) {
			//fmt.Printf("%v\n", s)
			fmt.Println("=========")
			for _, x := range s {
				for i := 1; i < x; i++ {
					fmt.Printf(" o ")
				}
				fmt.Printf(" X ")
				for i := x + 1; i <= 8; i++ {
					fmt.Printf(" o ")
				}
			}
		}

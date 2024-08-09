package main

import (
	"fmt"
	"time"
)

// Slice length
const Length = 100_000_000

// Slice capacity
const Capacity = 100_000_000

func main() {
	start := time.Now()

	Method1()
	fmt.Printf("Complete Method1\n")

	Method2()
	fmt.Printf("Complete Method2\n")

	Method3()
	fmt.Printf("Complete Method3\n")

	fmt.Printf("Process Time: %v\n", time.Since(start))
}

// Method1 : 前もってスライスを作成して，for文でインデックスを代入する．このとき，append関数を使う．また，スライスはmake関数で作成するが，引数を指定しない．
func Method1() []int {
	s := make([]int, 0)
	for i := range Length {
		s = append(s, i)
	}
	return s
}

// Method2 : 前もってスライスを作成して，for文でインデックスを代入する．このとき，append関数を使わない．また，スライスはmake関数で作成するが，引数を指定する(lengthのみ)．
func Method2() []int {
	s := make([]int, Length)
	for i := range Length {
		s[i] = i
	}
	return s
}

// Method3 : 前もってスライスを作成して，for文でインデックスを代入する．このとき，スライスはmake関数で作成するが，引数を指定する(capacityのみ)．
func Method3() []int {
	s := make([]int, 0, Capacity)
	for i := range Length {
		s = append(s, i)
	}
	return s
}

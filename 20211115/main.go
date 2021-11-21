package main

import (
	"fmt"
	"sort"
)

func main() {
	var a []int = []int{100, 6, 3, 5}
	var b []int = []int{4, 7, 1, 3, 5}

	c := hebing(a, b)
	// fmt.Println(c)
	var temp map[int]int = make(map[int]int)
	for i := 0; i < len(c); i++ {
		if _, ok := temp[c[i]]; ok {
			temp[c[i]] = temp[c[i]] + 1
		} else {
			temp[c[i]] = 1
		}
	}
	var d []int = make([]int, 0)
	for k := range temp {
		d = append(d, k)
	}

	sort.Ints(d)
	fmt.Println(d)

}

func hebing(a, b []int) (c []int) {

	if len(a) > len(b) {
		a = append(a, b...)
		return a
	} else {
		b = append(b, a...)
		return b
	}
}

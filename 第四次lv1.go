package main

import (
	"fmt"
	"strings"
)

var (
	str string
	num string
	c   int
	b   string = " "
)

func main() {
	fmt.Scan(&c)
	for i := 0; i < c; i++ {
		fmt.Scan(&num)
		str = str + num + b
	}
	m1 := make(map[string]int)
	for _, v := range strings.Fields(str) {
		if _, ok := m1[v]; ok {
			m1[v] = m1[v] + 1
		} else {
			m1[v] = 1
		}
	}
	for key, value := range m1 {
		fmt.Printf("%q:%d\n", key, value)
	}

}

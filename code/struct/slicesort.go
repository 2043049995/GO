package _struct

import (
	"fmt"
	"sort"
)

type UserSort struct {
	ID    int
	Name  string
	Score float64
}

func SliceSort() {
	list := [][2]int{{1, 2}, {2, 4}, {5, 2}, {6, 8}, {9, 3}}

	sort.Slice(list, func(i, j int) bool {
		fmt.Println(i, j) //i和j的值如何传入
		return list[i][1] < list[j][1]

	})
	fmt.Println(list)
}

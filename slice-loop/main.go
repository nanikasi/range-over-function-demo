package main

import (
	"fmt"
)

func main() {
    // Set型: 重複しない要素の集合を表すコンテナ型
	mySet := NewSet()

    // 要素を追加
	mySet.Add(1)
	mySet.Add(2)
	mySet.Add(3)

    // mySetの要素を2倍して出力するループ処理
	for _, value := range mySet.ToSlice() {
		fmt.Println(value * 2)
	}
}

type Set struct {
    elements map[int]struct{}
}

func NewSet() *Set {
    return &Set{
        elements: make(map[int]struct{}),
    }
}

func (s *Set) Add(value int) {
    s.elements[value] = struct{}{}
}

// Setの全要素をスライスとして取得するメソッド
func (s *Set) ToSlice() []int {
    keys := make([]int, 0, len(s.elements))
    for key := range s.elements {
        keys = append(keys, key)
    }
    return keys
}

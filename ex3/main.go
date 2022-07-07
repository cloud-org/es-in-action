package main

import (
	"sync"
)

var mu sync.Mutex

func fetchData() []int {
	return []int{1, 2, 3}
}

func overlay() {
	var m sync.Mutex
	m.Lock()
	m.Unlock()
	mu = m
}

func main() {
	//go func() {
	//	for {
	//		time.Sleep(1 * time.Second)
	//		overlay()
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		time.Sleep(1000 * 1000)
	//		_, _ = dict.Load(1)
	//	}
	//}()
	overlay()
	//mu.Lock()
	mu.Unlock()

}

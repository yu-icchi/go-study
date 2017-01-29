package main

import (
	"sync"
	"time"
)

var mu sync.RWMutex
var data map[string]string

func main() {
	data = map[string]string{"hoge": "fuga"}
	mu = sync.RWMutex{}
	go read(1)
	go read(2)
	go write(3)
	go read(4)
	time.Sleep(5 * time.Second)
}

func read(n int) {
	println("read_start", n)
	mu.RLock()
	defer mu.RUnlock()
	time.Sleep(1 * time.Second)
	println("read_complete", data["hoge"], n)
}

func write(n int) {
	println("write_start", n)
	mu.Lock()
	defer mu.Unlock()
	time.Sleep(2 * time.Second)
	data["hoge"] = "piyo"
	println("write_complete", n)
}

package dmap

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func Test(t *testing.T) {
	var n = 10
	wg := sync.WaitGroup{}
	wg.Add(n)

	var dict = NewMap[string, string]()
	for i := 1; i <= n; i++ {
		go func(i int) {
			defer func() {
				wg.Done()
			}()
			dict.Del("name")
			dict.Set("name", "dong")
			dict.Get("name")
		}(i)
	}
	fmt.Print(dict.internal)
	wg.Wait()
}

func BenchmarkDMap(b *testing.B) {
	var n = b.N
	wg := sync.WaitGroup{}
	wg.Add(n)

	var dict = NewMap[*string, string]()
	keyName := "dong"
	for i := 1; i <= n; i++ {
		go func(i int) {
			defer func() {
				wg.Done()
			}()
			dict.Set(&keyName, keyName)
			dict.Del(&keyName)
			dict.Get(&keyName)
		}(i)
	}
	wg.Wait()
	dict.Close()
}

func BenchmarkCmd(b *testing.B) {
	var n = b.N
	for i := 1; i <= n; i++ {
		cmd := NewCommand[string, string]("CLOSE", "dong", "dong")
		cmd.kind = "GET"
	}
}

func BenchmarkAlloc(b *testing.B) {
	var n = b.N
	for i := 1; i <= n; i++ {
		m := NewMap[string, int]()
		m.Close()
	}
}

func BenchmarkOriginalMap(b *testing.B) {
	var n = b.N
	m := make(map[string]string)
	for i := 1; i <= n; i++ {
		key := strconv.Itoa(i)
		m[key] = key
		_ = m[key]
		delete(m, key)
	}
}

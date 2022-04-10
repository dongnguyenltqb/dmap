package dmap

import (
	"fmt"
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

func BenchmarkOps(b *testing.B) {
	var n = b.N
	wg := sync.WaitGroup{}
	wg.Add(n)

	var dict = NewMap[string, string]()
	for i := 1; i <= n; i++ {
		go func(i int) {
			defer func() {
				wg.Done()
			}()
			dict.Set("name", "dong")
			dict.Del("name")
			dict.Get("name")
		}(i)
	}
	wg.Wait()
	dict.Close()
}

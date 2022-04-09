package dmap

import (
	"sync"
	"testing"
)

func Test(t *testing.T) {
	var n = 100
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
}

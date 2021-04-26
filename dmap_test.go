package dmap

import (
	"strconv"
	"sync"
	"testing"
)

func Test(t *testing.T) {
	wg := sync.WaitGroup{}
	// test concurency
	n := 1000
	wg.Add(n)
	m := NewMap()
	for i := 1; i <= n; i++ {
		go func(i int) {
			defer wg.Done()
			m.Set("firstName", strconv.Itoa(i))
		}(i)
	}
	wg.Wait()
	// test value
	wg.Add(2)
	go func() {
		defer wg.Done()
		m.Set("firstName", "Dong")
		m.Set("age", 24)
	}()
	go func() {
		defer wg.Done()
		m.Set("lastName", "nguyen")
		m.Set("age", 25)
	}()
	wg.Wait()

	m.Del("age")
	firstName := m.Get("firstName").(string)
	lastName := m.Get("lastName").(string)
	age := m.Get("age")
	m.Close()
	if firstName != "Dong" || lastName != "nguyen" || age != nil {
		t.Error("Expect Dong nguyen <nil> result = ", firstName, lastName, age)
	}

}

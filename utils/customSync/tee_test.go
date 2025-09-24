package customSync

import (
	"sync"
	"testing"
)

func TestTee(t *testing.T) {
	values := []int{0, 1, 2, 3, 10, 20, 30, 100, 200, 300}

	source := make(chan int)

	out1, out2 := Tee(source)
	go func() {
		for _, value := range values {
			source <- value
		}
	}()
	wg := sync.WaitGroup{}
	wg.Go(
		func() {
			for i := range len(values) {
				if v := <-out1; v != values[i] {
					t.Errorf(`extpected v %d to be %d`, v, values[i])
				}
			}
		})
	wg.Go(func() {
		for i := range len(values) {
			if v := <-out2; v != values[i] {
				t.Errorf(`extpected v %d to be %d`, v, values[i])
			}
		}
	})
	wg.Wait()
}

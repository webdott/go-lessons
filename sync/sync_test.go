package main

import (
	"sync"
	"testing"
)

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, &counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup

		for i := 0; i < wantedCount; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				counter.Inc()
			}()
		}

		wg.Wait()

		assertCounter(t, &counter, wantedCount)
	})
}

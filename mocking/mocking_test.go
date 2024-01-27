package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("ensures that we print 3, 2, 1, Go!", func(t *testing.T) {
		buffer := bytes.Buffer{}
		spySleeper := SpySleeper{}

		CountDown(&buffer, &spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("ensures that we sleep before each write", func(t *testing.T) {
		spySleepPrinter := SpyCountdownOperations{}

		CountDown(&spySleepPrinter, &spySleepPrinter)

		got := spySleepPrinter.Calls
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

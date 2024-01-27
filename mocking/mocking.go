package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

type SpyCountdownOperations struct {
	Calls []string
}

type SpySleeper struct {
	Calls int
}

const (
	finalWord      = "Go!"
	countdownStart = 3
	sleep          = "sleep"
	write          = "write"
)

func (spyT *SpyTime) Sleep(duration time.Duration) {
	spyT.durationSlept = duration
}

func (cs *ConfigurableSleeper) Sleep() {
	cs.sleep(cs.duration)
}

func (sco *SpyCountdownOperations) Sleep() {
	sco.Calls = append(sco.Calls, sleep)
}

func (sco *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	sco.Calls = append(sco.Calls, write)
	return
}

func (ss *SpySleeper) Sleep() {
	ss.Calls += 1
}

func CountDown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}

	fmt.Fprint(w, finalWord)
}

func main() {
	cSleeper := ConfigurableSleeper{1 * time.Second, time.Sleep}
	CountDown(os.Stdout, &cSleeper)
}

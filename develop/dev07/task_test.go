package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	workTime := time.Since(start)
	wantWorkTime := 1 * time.Second
	offsetWorkTime := 300 * time.Millisecond
	if workTime > wantWorkTime+offsetWorkTime {
		t.Fatalf("Want: %v, got: %v\n", wantWorkTime, workTime)
	}
}

func TestOrOneChan(t *testing.T) {
	start := time.Now()
	<-or()
	workTime := time.Since(start)
	wantWorkTime := 0 * time.Second
	offsetWorkTime := 100 * time.Millisecond
	if workTime > wantWorkTime+offsetWorkTime {
		t.Fatalf("Want: %v, got: %v\n", wantWorkTime, workTime)
	}
}

func TestOrNoChan(t *testing.T) {
	start := time.Now()
	<-or(sig(2 * time.Second))
	workTime := time.Since(start)
	wantWorkTime := 2 * time.Second
	offsetWorkTime := 300 * time.Millisecond
	if workTime > wantWorkTime+offsetWorkTime {
		t.Fatalf("Want: %v, got: %v\n", wantWorkTime, workTime)
	}
}

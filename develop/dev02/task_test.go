package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestUnpack(t *testing.T) {
	str := "a4bc2d5e"
	want := "aaaabccddddde"
	result, err := UnpackString(str)
	if want != result {
		t.Fatalf("Error: %v, case: %s, want: %s, got: %s", err, str, want, result)
	}
}

func TestNoUnpack(t *testing.T) {
	str := "abcd"
	want := "abcd"
	result, err := UnpackString(str)
	if want != result {
		t.Fatalf("Error: %v, case: %s, want: %s, got: %s", err, str, want, result)
	}
}

func TestOneLetter(t *testing.T) {
	str := "a"
	want := "a"
	result, err := UnpackString(str)
	if want != result {
		t.Fatalf("Error: %v, case: %s, want: %s, got: %s", err, str, want, result)
	}
}

func TestTwoDigits(t *testing.T) {
	str := "a4bc2d55e"
	want := errors.New("Incorrect string")
	result, err := UnpackString(str)
	if fmt.Sprint(want) != fmt.Sprint(err) || err == nil {
		t.Fatalf("Error: %v, case: %s, want: %s, got: %s", err, str, want, result)
	}
}

func TestFirstDigit(t *testing.T) {
	str := "1a4bc2d5e"
	want := errors.New("Incorrect string")
	result, err := UnpackString(str)
	if fmt.Sprint(want) != fmt.Sprint(err) || err == nil {
		t.Fatalf("Error: %v, case: %s, want: %s, got: %s", err, str, want, result)
	}
}

func TestDigits(t *testing.T) {
	str := "45"
	want := errors.New("Incorrect string")
	result, err := UnpackString(str)
	if fmt.Sprint(want) != fmt.Sprint(err) || err == nil {
		t.Fatalf("Error: %v, case: %s, want: %s, got: %s", err, str, want, result)
	}
}

func TestEmpty(t *testing.T) {
	str := ""
	want := ""
	result, err := UnpackString(str)
	if want != result {
		t.Fatalf("Error: %v, case: %s, want: %s, got: %s", err, str, want, result)
	}
}

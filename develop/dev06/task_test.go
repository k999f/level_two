package main

import (
	"reflect"
	"testing"
)

func TestCutFlagF(t *testing.T) {
	lines := []string{
		"111 222 333 444",
		"55 66 77",
		"8 9 10 11 12",
		"13 14 15",
		"16",
	}
	want := []string{
		"222",
		"66",
		"9",
		"14",
	}
	flags.F = 2
	flags.D = " "
	result := cutLines(lines)
	flags.F = -1
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Case:\n%s\nWant:\n%s\nGot:\n%s", getPrettyString(lines), getPrettyString(want), getPrettyString(result))
	}
}

func TestCutFlagD(t *testing.T) {
	lines := []string{
		"111.222.333.444",
		"55.66.77",
		"8.9.10.11.12",
		"13.14.15",
		"16",
	}
	want := []string{
		"333",
		"77",
		"10",
		"15",
	}
	flags.F = 3
	flags.D = "."
	result := cutLines(lines)
	flags.F = -1
	flags.D = " "
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Case:\n%s\nWant:\n%s\nGot:\n%s", getPrettyString(lines), getPrettyString(want), getPrettyString(result))
	}
}

func TestCutFlagS(t *testing.T) {
	lines := []string{
		"111 222 333 444",
		"55.66.77",
		"8 9 10 11 12",
		"13 14.15",
		"16",
	}
	want := []string{
		"66",
		"15",
	}
	flags.F = 2
	flags.D = "."
	flags.S = true
	result := cutLines(lines)
	flags.F = -1
	flags.D = " "
	flags.S = false
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Case:\n%s\nWant:\n%s\nGot:\n%s", getPrettyString(lines), getPrettyString(want), getPrettyString(result))
	}
}

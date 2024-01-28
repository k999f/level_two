package main

import (
	"reflect"
	"testing"
)

func TestSortNoFlags(t *testing.T) {
	lines := []string{
		"aword dword cword",
		"dword bword aword",
		"cword eword bword",
		"eword cword cword",
		"bword fword bword",
	}
	want := []string{
		"aword dword cword",
		"bword fword bword",
		"cword eword bword",
		"dword bword aword",
		"eword cword cword",
	}
	result := sortNoFlags(lines)
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Case:\n%s\nWant:\n%s\nGot:\n%s", getPrettyString(lines), getPrettyString(want), getPrettyString(result))
	}
}

func TestSortFlagK(t *testing.T) {
	flags.K = 2

	lines := []string{
		"aword dword cword",
		"dword bword aword",
		"cword eword bword",
		"eword cword cword",
		"bword fword bword",
	}
	want := []string{
		"dword bword aword",
		"eword cword cword",
		"aword dword cword",
		"cword eword bword",
		"bword fword bword",
	}
	result := sortFlagK(lines)
	flags.K = 0
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Case:\n%s\nWant:\n%s\nGot:\n%s", getPrettyString(lines), getPrettyString(want), getPrettyString(result))
	}
}

func TestSortFlagN(t *testing.T) {
	lines := []string{
		"1word dword cword",
		"3word bword aword",
		"5word eword bword",
		"4word cword cword",
		"2word fword bword",
	}
	want := []string{
		"1word dword cword",
		"2word fword bword",
		"3word bword aword",
		"4word cword cword",
		"5word eword bword",
	}
	result := sortFlagN(lines)
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Case:\n%s\nWant:\n%s\nGot:\n%s", getPrettyString(lines), getPrettyString(want), getPrettyString(result))
	}
}

func TestSortFlagR(t *testing.T) {
	lines := []string{
		"aword dword cword",
		"dword bword aword",
		"cword eword bword",
		"eword cword cword",
		"bword fword bword",
	}
	want := []string{
		"eword cword cword",
		"dword bword aword",
		"cword eword bword",
		"bword fword bword",
		"aword dword cword",
	}
	result := sortFlagR(lines)
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Case:\n%s\nWant:\n%s\nGot:\n%s", getPrettyString(lines), getPrettyString(want), getPrettyString(result))
	}
}

func TestSortFlagU(t *testing.T) {
	lines := []string{
		"aword dword cword",
		"cword eword bword",
		"dword bword aword",
		"cword eword bword",
		"eword cword cword",
		"bword fword bword",
	}
	want := []string{
		"aword dword cword",
		"cword eword bword",
		"dword bword aword",
		"eword cword cword",
		"bword fword bword",
	}
	result := sortFlagU(lines)
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Case:\n%s\nWant:\n%s\nGot:\n%s", getPrettyString(lines), getPrettyString(want), getPrettyString(result))
	}
}

func TestSortFlagM(t *testing.T) {
	lines := []string{
		"January dword cword",
		"March eword bword",
		"December bword aword",
		"February eword bword",
		"July cword cword",
		"June fword bword",
	}
	want := []string{
		"January dword cword",
		"February eword bword",
		"March eword bword",
		"June fword bword",
		"July cword cword",
		"December bword aword",
	}
	result := sortFlagM(lines)
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Case:\n%s\nWant:\n%s\nGot:\n%s", getPrettyString(lines), getPrettyString(want), getPrettyString(result))
	}
}

func TestSortFlagB(t *testing.T) {
	lines := []string{
		" aword dword cword ",
		" dword bword aword  ",
		"   cword eword bword   ",
		" eword cword cword ",
		"  bword fword bword",
	}
	want := []string{
		" aword dword cword ",
		"  bword fword bword",
		"   cword eword bword   ",
		" dword bword aword  ",
		" eword cword cword ",
	}
	result := sortFlagB(lines)
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Case:\n%s\nWant:\n%s\nGot:\n%s", getPrettyString(lines), getPrettyString(want), getPrettyString(result))
	}
}

func TestSortFlagH(t *testing.T) {
	lines := []string{
		"123.456 aword dword cword",
		"123.123 dword bword aword",
		"123 cword eword bword",
		"13 eword cword cword",
		"8 bword eword bword",
	}
	want := []string{
		"123 cword eword bword",
		"123.123 dword bword aword",
		"123.456 aword dword cword",
		"13 eword cword cword",
		"8 bword eword bword",
	}
	result := sortFlagB(lines)
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Case:\n%s\nWant:\n%s\nGot:\n%s", getPrettyString(lines), getPrettyString(want), getPrettyString(result))
	}
}

func TestSortFlagC(t *testing.T) {
	lines := []string{
		"aword dword cword",
		"bword fword bword",
		"cword eword bword",
		"dword bword aword",
		"eword cword cword",
	}
	want := true
	result := sortFlagC(lines)
	if result != want {
		t.Fatalf("Case:\n%s\nWant:\n%t\nGot:\n%t", getPrettyString(lines), want, result)
	}
}

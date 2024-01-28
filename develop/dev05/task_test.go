package main

import (
	"testing"
)

func TestGrepRegexp(t *testing.T) {
	lines := []string{
		"abcdefgh test test test",
		"HIJK lmnopq helloA TEST",
		"RsTuVwXyZ 123456",
		"67890 qwerty",
		"1830 aaa",
		"+++",
	}
	want := "1. abcdefgh test test test\n5. 1830 aaa"
	pattern := "[a-c]"
	flags.N = true
	resultLines, resultLinesIdx := grepLines(lines, pattern)
	flags.N = false
	result := getPrettyString(resultLines, resultLinesIdx)
	if result != want {
		t.Fatalf("Case:\n%s\nWant:\n%s\nGot:\n%s", getPrettyString(lines, nil), want, result)
	}
}

func TestGrepRegexpFlagI(t *testing.T) {
	lines := []string{
		"abcdefgh test test test",
		"HIJK lmnopq helloA TEST",
		"RsTuVwXyZ 123456",
		"67890 qwerty",
		"1830 aaa",
		"+++",
	}
	want := "1. abcdefgh test test test\n2. HIJK lmnopq helloA TEST\n5. 1830 aaa"
	pattern := "[a-c]"
	flags.I = true
	flags.N = true
	resultLines, resultLinesIdx := grepLines(lines, pattern)
	flags.I = false
	flags.N = false
	result := getPrettyString(resultLines, resultLinesIdx)
	if result != want {
		t.Fatalf("Case:\n%s\nWant:\n%s\nGot:\n%s", getPrettyString(lines, nil), want, result)
	}
}

func TestGrepFixed(t *testing.T) {
	lines := []string{
		"abcdefgh test test test",
		"HIJK lmnopq helloA TEST",
		"[a-zA-Z]",
		"67890 qwerty",
		"1830 aaa",
		"+++",
	}
	want := "3. [a-zA-Z]"
	pattern := "[a-zA-Z]"
	flags.N = true
	flags.F = true
	resultLines, resultLinesIdx := grepLines(lines, pattern)
	flags.N = false
	flags.F = false
	result := getPrettyString(resultLines, resultLinesIdx)
	if result != want {
		t.Fatalf("Case:\n%s\nWant:\n%s\nGot:\n%s", getPrettyString(lines, nil), want, result)
	}
}

func TestGrepFixedFlagI(t *testing.T) {
	lines := []string{
		"abcdefgh test test test",
		"HIJK lmnopq helloA TEST",
		"AbCdE",
		"abcde",
		"1830 aaa",
		"+++",
	}
	want := "3. AbCdE\n4. abcde"
	pattern := "ABCDE"
	flags.N = true
	flags.F = true
	flags.I = true
	resultLines, resultLinesIdx := grepLines(lines, pattern)
	flags.N = false
	flags.F = false
	flags.I = false
	result := getPrettyString(resultLines, resultLinesIdx)
	if result != want {
		t.Fatalf("Case:\n%s\nWant:\n%s\nGot:\n%s", getPrettyString(lines, nil), want, result)
	}
}

func TestGrepFlagVFlagI(t *testing.T) {
	lines := []string{
		"abcdefgh test test test",
		"HIJK lmnopq helloA TEST",
		"AbCdE",
		"1830",
		"abcde",
		"+++",
	}
	want := "4. 1830\n6. +++"
	pattern := "[a-z]"
	flags.N = true
	flags.I = true
	flags.V = true
	resultLines, resultLinesIdx := grepLines(lines, pattern)
	flags.N = false
	flags.I = false
	flags.V = false
	result := getPrettyString(resultLines, resultLinesIdx)
	if result != want {
		t.Fatalf("Case:\n%s\nWant:\n%s\nGot:\n%s", getPrettyString(lines, nil), want, result)
	}
}

func TestGrepFlagAFlagB(t *testing.T) {
	lines := []string{
		"abcdefgh test test test",
		"HIJK lmnopq helloA TEST",
		"AbCdE",
		"1830",
		"abcde",
		"+++",
	}
	want := "4. 1830\n3. AbCdE\n5. abcde\n6. +++"
	pattern := "1830"
	flags.N = true
	flags.F = true
	flags.B = 1
	flags.A = 2
	resultLines, resultLinesIdx := grepLines(lines, pattern)
	flags.N = false
	flags.F = false
	flags.B = 0
	flags.A = 0
	result := getPrettyString(resultLines, resultLinesIdx)
	if result != want {
		t.Fatalf("Case:\n%s\nWant:\n%s\nGot:\n%s", getPrettyString(lines, nil), want, result)
	}
}

func TestGrepFlagContextFlagF(t *testing.T) {
	lines := []string{
		"abcdefgh test test test",
		"HIJK lmnopq helloA TEST",
		"AbCdE",
		"1830",
		"abcde",
		"+++",
	}
	want := "4. 1830\n3. AbCdE\n5. abcde"
	pattern := "1830"
	flags.N = true
	flags.F = true
	flags.Context = 1
	resultLines, resultLinesIdx := grepLines(lines, pattern)
	flags.N = false
	flags.F = false
	result := getPrettyString(resultLines, resultLinesIdx)
	if result != want {
		t.Fatalf("Case:\n%s\nWant:\n%s\nGot:\n%s", getPrettyString(lines, nil), want, result)
	}
}

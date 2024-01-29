package main

import (
	"testing"
)

func TestCdWrongPath(t *testing.T) {
	input := []string{
		"cd",
		"some.wrong.path",
	}
	err := shell(input)
	if err == nil {
		t.Fatalf("Case: %s, expected error, got: %v", input, err)
	}
}

func TestPwdExtraArgs(t *testing.T) {
	input := []string{
		"pwd",
		"some",
		"extra",
		"arguments",
	}
	err := shell(input)
	if err == nil {
		t.Fatalf("Case: %s, expected error, got: %v", input, err)
	}
}

func TestEchoExtraArgs(t *testing.T) {
	input := []string{
		"echo",
	}
	err := shell(input)
	if err == nil {
		t.Fatalf("Case: %s, expected error, got: %v", input, err)
	}
}

func TestKillWrongPid(t *testing.T) {
	input := []string{
		"kill",
		"wrong_pid",
	}
	err := shell(input)
	if err == nil {
		t.Fatalf("Case: %s, expected error, got: %v", input, err)
	}
}

func TestPsExtraArgs(t *testing.T) {
	input := []string{
		"ps",
		"some",
		"extra",
		"arguments",
	}
	err := shell(input)
	if err == nil {
		t.Fatalf("Case: %s, expected error, got: %v", input, err)
	}
}

func TestExecWrongCommand(t *testing.T) {
	input := []string{
		"wrongcommand",
	}
	err := shell(input)
	if err == nil {
		t.Fatalf("Case: %s, expected error, got: %v", input, err)
	}
}

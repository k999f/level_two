package main

import (
	"testing"
)

func TestCorrectServer(t *testing.T) {
	address := "0.beevik-ntp.pool.ntp.org"
	_, err := getNtpTime(address)
	if err != nil {
		t.Fatalf("Case: %s, want: %v, got: %s", address, nil, err)
	}
}

func TestIncorrectServer(t *testing.T) {
	address := "incorrect.address.com"
	_, err := getNtpTime(address)
	if err == nil {
		t.Fatalf("Case: %s, want: %v, got: %s", address, nil, err)
	}
}

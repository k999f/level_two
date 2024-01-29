package main

import (
	"testing"
)

func TestWget(t *testing.T) {
	siteURL := "https://example.com/"
	_, err := wget(siteURL)
	if err != nil {
		t.Fatalf("Case: %s, expected error, got: %v", siteURL, err)
	}
}

func TestWgetError(t *testing.T) {
	siteURL := "https://this.is.some/wrong/url/"
	_, err := wget(siteURL)
	if err == nil {
		t.Fatalf("Case: %s, expected error, got: %v", siteURL, err)
	}
}

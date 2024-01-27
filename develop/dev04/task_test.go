package main

import (
	"reflect"
	"testing"
)

func TestAnagramsNoRepeat(t *testing.T) {
	words := []string{
		"пятак",
		"пятка",
		"тяпка",
		"листок",
		"слиток",
		"столик"}
	want := map[string][]string{
		"листок": []string{"листок", "слиток", "столик"},
		"пятак":  []string{"пятак", "пятка", "тяпка"}}
	result := getAnagramSets(&words)
	modifiedResult := make(map[string][]string)
	for k, v := range *result {
		modifiedResult[k] = *v
	}
	if !reflect.DeepEqual(modifiedResult, want) {
		t.Fatalf("Case:\n%v\nWant:\n%v\nGot:\n%v", words, want, modifiedResult)
	}
}

func TestAnagramsOneWordSet(t *testing.T) {
	words := []string{
		"пятак",
		"пятка",
		"тяпка",
		"листок",
		"слиток",
		"столик",
		"тест"}
	want := map[string][]string{
		"листок": []string{"листок", "слиток", "столик"},
		"пятак":  []string{"пятак", "пятка", "тяпка"}}
	result := getAnagramSets(&words)
	modifiedResult := make(map[string][]string)
	for k, v := range *result {
		modifiedResult[k] = *v
	}
	if !reflect.DeepEqual(modifiedResult, want) {
		t.Fatalf("Case:\n%v\nWant:\n%v\nGot:\n%v", words, want, modifiedResult)
	}
}

func TestAnagramswithSort(t *testing.T) {
	words := []string{
		"абвг",
		"гвба",
		"бвга",
		"вбга",
		"слиток",
		"столик"}
	want := map[string][]string{
		"абвг":   []string{"абвг", "бвга", "вбга", "гвба"},
		"слиток": []string{"слиток", "столик"}}
	result := getAnagramSets(&words)
	modifiedResult := make(map[string][]string)
	for k, v := range *result {
		modifiedResult[k] = *v
	}
	if !reflect.DeepEqual(modifiedResult, want) {
		t.Fatalf("Case:\n%v\nWant:\n%v\nGot:\n%v", words, want, modifiedResult)
	}
}

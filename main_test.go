package main

import (
	"reflect"
	"testing"
)

func TestGetQuestionsFromCSV(t *testing.T) {
	questions, err := getQuestionsFromCSV("./test.csv")
	if err != nil {
		t.Errorf("Failed to get Questions from csv %s", err.Error())
	}

	got := *questions
	want := []question{{ask: "Iama", answer: 1}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CSV data not as expected got: %v want: %v", got, want)
	}
}

func TestReadFile(t *testing.T) {
	got, err := readFile("./test.csv")
	if err != nil {
		t.Errorf("File could not be read %s", err.Error())
	}

	want := "Iama,1"

	if got != want {
		t.Errorf("Output of file does not match. got: %s wanted:%s", got, want)
	}
}

func TestParseCSV(t *testing.T) {
	result, err := parseCSV("Iama,1")
	if err != nil {
		t.Errorf("Failed to parse CSV reason: %s", err.Error())
	}

	got := *result
	want := []question{{ask: "Iama", answer: 1}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CSV data not as expected got: %v want: %v", got, want)
	}

}

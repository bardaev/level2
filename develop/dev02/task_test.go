package main

import (
	"errors"
	"testing"
)

var TestFailed error = errors.New("Test failed")

func TestUnpack(t *testing.T) {
	// "a4bc2d5e" => "aaaabccddddde"
	// "abcd" => "abcd"
	// "45" => "" (некорректная строка)
	// "" => ""

	var input1 string = "a4bc2d5e"
	var output1 string = "aaaabccddddde"

	var input2 string = "abcd"
	var output2 string = "abcd"

	result1, err := Unpack(input1)
	if err != nil || result1 != output1 {
		t.Error(TestFailed)
	}
	t.Logf("\nInput string: %s\nOutput string: %s\n", input1, result1)

	result2, err := Unpack(input2)
	if err != nil || result2 != output2 {
		t.Error(TestFailed)
	}
	t.Logf("\nInput string: %s\nOutput string: %s\n", input2, result2)
}

func TestFailedUnpackNumber(t *testing.T) {
	var input string = "45"

	result, err := Unpack(input)

	if err != nil {
		t.Logf("Input string %s returned error", input)
	}

	if result != "" {
		t.Errorf("Function returned bad result: %s", result)
	}
}

func TestFailedUnpackEmptyString(t *testing.T) {
	var input string = ""

	result, err := Unpack(input)

	if err != nil {
		t.Logf("Input string %s returned error", input)
	}

	if result != "" {
		t.Errorf("Function returned bad result: %s", result)
	}
}

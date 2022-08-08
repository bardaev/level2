package main

import (
	"testing"
)

// Данная программа не возвращает результат, поэтому прохождение тестов необходимо определять по выводу в консоли
// Тесты привязаны к содержимому test.txt

func TestAfter(t *testing.T) {
	var after int = 2
	var before int = 0
	var context int = 0
	var count bool = false
	var ignoreCase bool = false
	var invert bool = false
	var fixed bool = false
	var lineNum bool = false
	var template string = "ii"
	var filename string = "test.txt"
	Start(&after, &before, &context, &count, &ignoreCase, &invert, &fixed, &lineNum, template, filename)
}

func TestBefore(t *testing.T) {
	var after int = 0
	var before int = 3
	var context int = 0
	var count bool = false
	var ignoreCase bool = false
	var invert bool = false
	var fixed bool = false
	var lineNum bool = false
	var template string = "ii"
	var filename string = "test.txt"
	Start(&after, &before, &context, &count, &ignoreCase, &invert, &fixed, &lineNum, template, filename)
}

func TestContext(t *testing.T) {
	var after int = 0
	var before int = 0
	var context int = 3
	var count bool = false
	var ignoreCase bool = false
	var invert bool = false
	var fixed bool = false
	var lineNum bool = false
	var template string = "ii"
	var filename string = "test.txt"
	Start(&after, &before, &context, &count, &ignoreCase, &invert, &fixed, &lineNum, template, filename)
}

func TestCount(t *testing.T) {
	var after int = 0
	var before int = 0
	var context int = 0
	var count bool = true
	var ignoreCase bool = false
	var invert bool = false
	var fixed bool = false
	var lineNum bool = false
	var template string = "ii"
	var filename string = "test.txt"
	Start(&after, &before, &context, &count, &ignoreCase, &invert, &fixed, &lineNum, template, filename)
}

func TestInvert(t *testing.T) {
	var after int = 0
	var before int = 0
	var context int = 0
	var count bool = false
	var ignoreCase bool = false
	var invert bool = true
	var fixed bool = false
	var lineNum bool = false
	var template string = "ii"
	var filename string = "test.txt"
	Start(&after, &before, &context, &count, &ignoreCase, &invert, &fixed, &lineNum, template, filename)
}

func TestLineNum(t *testing.T) {
	var after int = 0
	var before int = 0
	var context int = 0
	var count bool = false
	var ignoreCase bool = false
	var invert bool = false
	var fixed bool = false
	var lineNum bool = true
	var template string = "ii"
	var filename string = "test.txt"
	Start(&after, &before, &context, &count, &ignoreCase, &invert, &fixed, &lineNum, template, filename)
}

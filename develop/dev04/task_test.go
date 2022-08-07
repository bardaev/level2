package main

import (
	"errors"
	"testing"
)

func TestAnagramm(t *testing.T) {
	var input []string = []string{"Автор", "пЯтка", "тяПка", "тоВАр", "пятак", "Отвар"}
	var output map[string][]string = map[string][]string{
		"автор": {"автор", "отвар", "товар"},
		"пятка": {"пятак", "пятка", "тяпка"},
	}

	var ArrNotEquals error = errors.New("Arrays not equal")
	var result map[string][]string = Anagramm(input)

	for k, v := range output {
		if val, ok := result[k]; ok {
			if len(val) != len(v) {
				t.Error(ArrNotEquals)
			}
			for i := 0; i < len(val); i++ {
				if val[i] != v[i] {
					t.Error(ArrNotEquals)
				}
			}
		}
	}
	t.Log("Test completed")
}

package main

import (
	"fmt"
	"unicode"
)

func main() {
	var arr []string = []string{"пЯтка", "тяПка", "пятак"}
	Anagramm(arr)
}

func Anagramm(a []string) map[string]string {
	var arr [][]rune = make([][]rune, 0)

	for _, item := range a {
		arr = append(arr, []rune(item))
	}

	for index, item := range arr {
		for jindex, jitem := range item {
			arr[index][jindex] = unicode.ToLower(jitem)
		}
	}

	for _, i := range arr {
		fmt.Println(string(i))
	}
	// fmt.Println(arr)
	return nil
}

func removeItem(arr [][]rune, i int) [][]rune {
	var tmp [][]rune = make([][]rune, len(arr))
	copy(tmp, arr)
	result := tmp[:i]
	result = append(result, tmp[i+1:]...)
	return result
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

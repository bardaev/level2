package main

import (
	"sort"
	"unicode"
)

// Anagramm is general function
func Anagramm(a []string) map[string][]string {
	var arr sortArrRunes = make(sortArrRunes, 0)

	// Переводим строку в руны и пишем в массив рун
	for _, item := range a {
		arr = append(arr, sortRunes{arr: []rune(item)})
	}

	// В нижний регистр
	for index, item := range arr {
		for jindex, jitem := range item.arr {
			arr[index].arr[jindex] = unicode.ToLower(jitem)
		}
	}

	// Данная мапа хранит индексы слайса arr, которые использовались для добавления анаграмм в результирующую мапу
	var usedValues map[int]int = make(map[int]int)
	// Результирующая мапа в виде рун, далее сконвертируем в мапу строк и вернем как результат
	var resultRuneMap map[*sortRunes]sortArrRunes = make(map[*sortRunes]sortArrRunes)

	for index, item := range arr {
		// Если этот индекс использовался пропускаем итерацию
		if _, ok := usedValues[index]; ok {
			continue
		}
		usedValues[index] = index

		// Кпируем слово
		var itemRune sortRunes = item

		// Добавляем в результирующую мапу
		resultRuneMap[&itemRune] = make(sortArrRunes, 0)
		// И добавляем его в массив значений
		resultRuneMap[&itemRune] = append(resultRuneMap[&itemRune], itemRune)

		// Итерируемся по слайсу arr
		for i := index + 1; i < len(arr); i++ {
			// Если в мапе уже есть использованный индекс, пропускаем итерацию
			if _, ok := usedValues[i]; ok {
				continue
			}
			// Если это анаграмма то пишем в результирующую мапу
			if equalsSortRune(getSortedAnagram(itemRune), getSortedAnagram(arr[i])) {
				usedValues[i] = i
				resultRuneMap[&itemRune] = append(resultRuneMap[&itemRune], arr[i])
			}
		}
		// Сортирум массив в мапе
		sort.Sort(resultRuneMap[&itemRune])
	}

	var resultStringMap map[string][]string = make(map[string][]string)

	// Конвертируем из рун в строку
	for k, v := range resultRuneMap {
		var keyString string = string(k.arr)
		resultStringMap[keyString] = make([]string, 0)
		for _, value := range v {
			resultStringMap[keyString] = append(resultStringMap[keyString], string(value.arr))
		}
	}

	return resultStringMap
}

// Сортируем буквы в словах, чтобы найти анаграммы
// напимер, если подадим 2 слова: тяпка и пятак, то получим оба "акптя"
// и если равны, то это анаграмма
func equalsSortRune(a sortRunes, b sortRunes) bool {
	a = getSortedAnagram(a)
	b = getSortedAnagram(b)

	if a.Len() != b.Len() {
		return false
	}

	for index, value := range a.arr {
		if value != b.arr[index] {
			return false
		}
	}

	return true
}

// Тут сортируем буквы в порядке возрастания, чтобы найти анаграммы
func getSortedAnagram(sr sortRunes) sortRunes {
	var copyitem sortRunes = sortRunes{arr: make([]rune, sr.Len())}
	copy(copyitem.arr, sr.arr)
	sort.Sort(copyitem)
	return copyitem
}

// Этот тип нужен, чтобы реализовать интерфейс сортировки для рун
type sortRunes struct {
	arr []rune
}

func (s sortRunes) Less(i, j int) bool {
	return s.arr[i] < s.arr[j]
}

func (s sortRunes) Swap(i, j int) {
	s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
}

func (s sortRunes) Len() int {
	return len(s.arr)
}

// Этот тип нужен для сортировки массивов рун
type sortArrRunes []sortRunes

func (sa sortArrRunes) Less(i, j int) bool {
	var iLen int = len(sa[i].arr)
	var jLen int = len(sa[j].arr)
	var minLen int
	if iLen < jLen {
		minLen = iLen
	} else {
		minLen = jLen
	}

	for ii, jj := 0, 0; ii < minLen; ii, jj = ii+1, jj+1 {
		if sa[i].arr[ii] > sa[j].arr[jj] {
			break
		} else if sa[i].arr[ii] < sa[j].arr[jj] {
			return true
		}
	}
	return false
}

func (sa sortArrRunes) Swap(i, j int) {
	sa[i], sa[j] = sa[j], sa[i]
}

func (sa sortArrRunes) Len() int {
	return len(sa)
}

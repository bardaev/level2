package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
	-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
	-n — сортировать по числовому значению
	-r — сортировать в обратном порядке
	-u — не выводить повторяющиеся строки
*/

func main() {
	var k *int = flag.Int("k", 0, "Колонка для сортировки")
	var n *bool = flag.Bool("n", false, "Сортировка по числовому значению")
	var r *bool = flag.Bool("r", false, "Сортировка в обратном порядке")
	var u *bool = flag.Bool("u", false, "Не выводить повторяющиеся строки")

	flag.Parse()

	var file string = flag.Arg(0)

	fmt.Printf("k %d, n %t, r %t, u %t, fname %s", *k, *n, *r, *u, file)

	Start(k, n, r, u, file)
}

// Start точка входа
func Start(k *int, n *bool, r *bool, u *bool, file string) {
	var in io.Reader
	if filename := file; filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error opening file: ", err)
			os.Exit(1)
		}
		defer f.Close()
		in = f
	} else {
		fmt.Println("Empty argument")
		os.Exit(1)
	}

	buf := bufio.NewReader(in)
	var lines []string = make([]string, 0)

	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		lines = append(lines, line)
	}

	var Output handler = &outputHandler{}
	var Reverse handler = &reverseHandler{
		Reverse: *r,
		Next:    Output,
	}
	var Sort handler = &sortHandler{
		Column: *k,
		Number: *n,
		Next:   Reverse,
	}
	var Unique handler = &uniqueHandler{
		Unique: *u,
		Next:   Sort,
	}

	Unique.SendRequest(lines)
}

type handler interface {
	SendRequest(data []string)
}

// Обработка уникальных значений
type uniqueHandler struct {
	Unique bool
	Next   handler
}

func (u *uniqueHandler) SendRequest(data []string) {
	if u.Unique {
		var uniqueStringMap map[string]int = make(map[string]int)
		for index, value := range data {
			uniqueStringMap[value] = index
		}
		data = make([]string, 0)
		for k := range uniqueStringMap {
			data = append(data, k)
		}
	}
	u.Next.SendRequest(data)
}

// Сортировка
type sortHandler struct {
	Column int
	Number bool
	Next   handler
}

func (s *sortHandler) SendRequest(data []string) {
	// Разбиваем на колонки
	var splitData [][]string = make([][]string, 0)
	for _, value := range data {
		var split []string = make([]string, 0)
		split = strings.Fields(value)
		splitData = append(splitData, split)
	}
	var sortSlice *sortSliceString = &sortSliceString{
		Data:   splitData,
		Column: s.Column,
		Number: s.Number,
	}
	// Сортируем
	sort.Sort(*sortSlice)
	// Собираем в массив
	data = make([]string, 0)
	for _, value := range sortSlice.Data {
		data = append(data, strings.Join(value, " "))
	}
	s.Next.SendRequest(data)
}

// В обратный порядок
type reverseHandler struct {
	Reverse bool
	Next    handler
}

func (r *reverseHandler) SendRequest(data []string) {
	if r.Reverse {
		for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
			data[i], data[j] = data[j], data[i]
		}
	}
	r.Next.SendRequest(data)
}

// Вывод в файл
type outputHandler struct{}

func (o *outputHandler) SendRequest(data []string) {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	writer := bufio.NewWriter(file)
	defer file.Close()

	for _, row := range data {
		writer.WriteString(row)
		writer.WriteString("\n")
	}
	writer.Flush()
}

type sortSliceString struct {
	Data   [][]string
	Column int
	Number bool
}

func (s sortSliceString) Less(i, j int) bool {
	// Сортировка по колонке и по числу
	if s.Column > 0 && s.Number {
		if len(s.Data[i]) < s.Column || len(s.Data[j]) < s.Column {
			return false
		}
		iNumber, ierr := strconv.Atoi(s.Data[i][s.Column-1])
		jNumber, jerr := strconv.Atoi(s.Data[j][s.Column-1])
		if ierr != nil || jerr != nil {
			return false
		}
		return iNumber < jNumber
	}

	// Сортировка по колонке
	if s.Column > 0 {
		if len(s.Data[i]) < s.Column || len(s.Data[j]) < s.Column {
			return false
		}
		var iColumnLen int = len(s.Data[i][s.Column-1])
		var jColumnLen int = len(s.Data[j][s.Column-1])
		var minColumnLen int = iColumnLen
		if !(iColumnLen < jColumnLen) {
			minColumnLen = jColumnLen
		}
		for ii, jj := 0, 0; ii < minColumnLen; ii, jj = ii+1, jj+1 {
			if s.Data[i][s.Column-1][ii] > s.Data[j][s.Column-1][jj] {
				break
			} else if s.Data[i][s.Column-1][ii] < s.Data[j][s.Column-1][jj] {
				return true
			}
		}
		return false
	}

	if s.Number {
		var iNumStr string = strings.Join(s.Data[i], " ")
		var jNumStr string = strings.Join(s.Data[j], " ")

		iNumber, ierr := strconv.Atoi(iNumStr)
		jNumber, jerr := strconv.Atoi(jNumStr)

		if ierr != nil || jerr != nil {
			return false
		}

		return iNumber < jNumber
	}

	// Сортировка если нет аргументов
	var tmpI string = strings.Join(s.Data[i], " ")
	var tmpJ string = strings.Join(s.Data[j], " ")

	var tmpILen int = len(tmpI)
	var tmpJLen int = len(tmpJ)
	var minTmpLen int = tmpILen

	if !(tmpILen < tmpJLen) {
		minTmpLen = tmpJLen
	}

	for ii, jj := 0, 0; ii < minTmpLen; ii, jj = ii+1, jj+1 {
		if tmpI[ii] > tmpJ[jj] {
			break
		} else if tmpI[ii] < tmpJ[jj] {
			return true
		}
	}
	return false
}

func (s sortSliceString) Len() int {
	return len(s.Data)
}

func (s sortSliceString) Swap(i, j int) {
	s.Data[i], s.Data[j] = s.Data[j], s.Data[i]
}

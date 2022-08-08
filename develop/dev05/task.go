package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// -A - "after" печатать +N строк после совпадения
// -B - "before" печатать +N строк до совпадения
// -C - "context" (A+B) печатать ±N строк вокруг совпадения
// -c - "count" (количество строк)
// -i - "ignore-case" (игнорировать регистр)
// -v - "invert" (вместо совпадения, исключать)
// -F - "fixed", точное совпадение со строкой, не паттерн
// -n - "line num", напечатать номер строки

/*
	Данная программа использует паттерн цепочка вызовов.
	На каждый флаг предусмотрен свой обработчик.
	Есть некоторая комбинация флагов, которые противоречат друг другу и непонятно как они обрабатываются в оригинальном grep.
	Поэтому при присутствии конфликтных флагов одни будут обнулять другие.
	Их всего 2: при присутствии after или before они будут обнулять contex и fixed обнулит ignorecase.
	В цеочке вызовов передаются 3 аргумента:
		- lines: строки в файле
		- template: регулярное выражение или точное совпадение при аргумента -F
		- result: итоговый результат
	result представлени в виде map[int]map[int]string
	В первой мапе в качестве ключа хранится индекс найденной строки.
	Во второй мапе в качестве ключа хранится индекс найденной строки и в виде значения сама строка.
	Такая структура необходима для работы after, before, context и lineNum, чтобы вставлять педыдущие и следующие строки вокруг найденного значения и нумерации строк.
*/

func main() {
	var after *int = flag.Int("A", 0, "Печать n строк после совпадения")
	var before *int = flag.Int("B", 0, "Печать n строк до совпадения")
	var context *int = flag.Int("C", 0, "Печать n строк вокруг совпадения")
	var count *bool = flag.Bool("c", false, "Количество строк")
	var ignoreCase *bool = flag.Bool("i", false, "Игнорировать регистр")
	var invert *bool = flag.Bool("v", false, "Исключать совпадения")
	var fixed *bool = flag.Bool("F", false, "Точное совпадение со строкой")
	var lineNum *bool = flag.Bool("n", false, "Печать номера строки")

	flag.Parse()

	var template string = flag.Arg(0)
	var filename string = flag.Arg(1)

	Start(after, before, context, count, ignoreCase, invert, fixed, lineNum, template, filename)
}

// Start точка входа для тестов
func Start(
	after *int,
	before *int,
	context *int,
	count *bool,
	ignoreCase *bool,
	invert *bool,
	fixed *bool,
	lineNum *bool,
	template string,
	file string,
) {

	if *after > 0 || *before > 0 {
		*context = 0
	}

	if *fixed {
		*ignoreCase = false
	}

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

	var Print handler = &printHandler{}
	var Count handler = &countHandler{
		Count: *count,
		Next:  Print,
	}
	var LineNum handler = &lineNumHandler{
		LineNum: *lineNum,
		Next:    Count,
	}
	var Context handler = &contextHandler{
		Context: *context,
		Next:    LineNum,
	}
	var Before handler = &beforeHandler{
		Before: *before,
		Next:   Context,
	}
	var After handler = &afterHandler{
		After: *after,
		Next:  Before,
	}
	var Invert handler = &invertHandler{
		Invert: *invert,
		Next:   After,
	}
	var IgnoreCase handler = &ignoreCaseHandler{
		IgnoreCase: *ignoreCase,
		Next:       Invert,
	}
	var Fixed handler = &fixedHandler{
		Fixed: *fixed,
		Next:  IgnoreCase,
	}
	var Find handler = &findHandler{
		Next: Fixed,
	}

	var result map[int]map[int]string = map[int]map[int]string{}
	Find.SendRequest(lines, template, result)
}

type handler interface {
	SendRequest(lines []string, template string, result map[int]map[int]string)
}

type findHandler struct {
	Next handler
}

func (a *findHandler) SendRequest(lines []string, template string, result map[int]map[int]string) {
	for index, value := range lines {
		matched, err := regexp.MatchString(template, value)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if matched {
			var innerMap map[int]string = map[int]string{}
			innerMap[index] = value
			result[index] = innerMap
		}
	}
	a.Next.SendRequest(lines, template, result)
}

type fixedHandler struct {
	Fixed bool
	Next  handler
}

func (a *fixedHandler) SendRequest(lines []string, template string, result map[int]map[int]string) {
	if a.Fixed {
		result = map[int]map[int]string{}
		for index, value := range lines {
			if strings.Contains(value, template) {
				var innerMap map[int]string = map[int]string{}
				innerMap[index] = value
				result[index] = innerMap
			}
		}
	}
	a.Next.SendRequest(lines, template, result)
}

type ignoreCaseHandler struct {
	IgnoreCase bool
	Next       handler
}

func (i *ignoreCaseHandler) SendRequest(lines []string, template string, result map[int]map[int]string) {
	if i.IgnoreCase {
		result = map[int]map[int]string{}
		for index, value := range lines {
			if strings.Contains(strings.ToLower(value), strings.ToLower(template)) {
				var innerMap map[int]string = map[int]string{}
				innerMap[index] = value
				result[index] = innerMap
			}
		}
	}
	i.Next.SendRequest(lines, template, result)
}

type invertHandler struct {
	Invert bool
	Next   handler
}

func (i *invertHandler) SendRequest(lines []string, template string, result map[int]map[int]string) {
	if i.Invert {
		var invertResult map[int]map[int]string = make(map[int]map[int]string)
		invertResult[0] = make(map[int]string)
		for index := range lines {
			if _, ok := result[index]; !ok {
				invertResult[0][index] = lines[index]
			}
		}
		result = invertResult
	}
	i.Next.SendRequest(lines, template, result)
}

type beforeHandler struct {
	Before int
	Next   handler
}

func (b *beforeHandler) SendRequest(lines []string, template string, result map[int]map[int]string) {
	var with int
	if b.Before > 0 {
		for k := range result {
			with = k - b.Before
			if with < 0 {
				with = 0
			}
			var tmp []string = make([]string, len(lines[with:k]))
			copy(tmp, lines[with:k])
			for _, value := range tmp {
				result[k][with] = value
				with++
			}
		}
	}
	b.Next.SendRequest(lines, template, result)
}

type afterHandler struct {
	After int
	Next  handler
}

func (a *afterHandler) SendRequest(lines []string, template string, result map[int]map[int]string) {
	var to int
	if a.After > 0 {
		for k := range result {
			to = k + a.After
			if to > len(lines) {
				to = len(lines)
			}
			var tmp []string = make([]string, len(lines[k:to]))
			copy(tmp, lines[k+1:to+1])
			var toIndex int = k + 1
			for _, value := range tmp {
				result[k][toIndex] = value
				toIndex++
			}
		}
	}
	a.Next.SendRequest(lines, template, result)
}

type contextHandler struct {
	Context int
	Next    handler
}

func (c *contextHandler) SendRequest(lines []string, template string, result map[int]map[int]string) {
	var with int
	var to int
	if c.Context > 0 {
		for k := range result {
			with = k - c.Context
			if with < 0 {
				with = 0
			}
			to = k + c.Context
			if to > len(lines) {
				to = len(lines)
			}
			var tmp1 []string = make([]string, len(lines[with:k]))
			copy(tmp1, lines[with:k])
			for _, value := range tmp1 {
				result[k][with] = value
				with++
			}
			var tmp2 []string = make([]string, len(lines[k:to]))
			copy(tmp2, lines[k+1:to+1])
			var toIndex int = k + 1
			for _, value := range tmp2 {
				result[k][toIndex] = value
				toIndex++
			}
		}
	}
	c.Next.SendRequest(lines, template, result)
}

type lineNumHandler struct {
	LineNum bool
	Next    handler
}

func (l *lineNumHandler) SendRequest(lines []string, template string, result map[int]map[int]string) {
	if l.LineNum {
		for key := range result {
			for k := range result[key] {
				result[key][k] = strconv.Itoa(k) + ":" + result[key][k]
			}
		}
	}
	l.Next.SendRequest(lines, template, result)
}

type countHandler struct {
	Count bool
	Next  handler
}

func (c *countHandler) SendRequest(lines []string, template string, result map[int]map[int]string) {
	if c.Count {
		var count int = len(result)
		result = make(map[int]map[int]string)
		result[0] = make(map[int]string)
		result[0][0] = strconv.Itoa(count)
	}
	c.Next.SendRequest(lines, template, result)
}

type printHandler struct {
}

func (p *printHandler) SendRequest(lines []string, template string, result map[int]map[int]string) {
	var keys []int
	for k := range result {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var count int = len(keys)
	var i int = 0
	for _, k := range keys {
		var innerKeys []int
		for innerK := range result[k] {
			innerKeys = append(innerKeys, innerK)
		}
		sort.Ints(innerKeys)

		for _, iV := range innerKeys {
			fmt.Print(result[k][iV])
		}
		i++
		if i != count && count > 1 {
			fmt.Println("--")
		}
	}
	if count == 1 {
		fmt.Println()
	}
}

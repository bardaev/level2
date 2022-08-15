package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	-f - "fields" - выбрать поля (колонки)
	-d - "delimiter" - использовать другой разделитель
	-s - "separated" - только строки с разделителем
*/

/*
	Данная программа считывает из консоли строки и выводит вырезанные из них подстроки
	Для обозночения конца ввода необходимо нажать ctrl+c
	Пример запуска:
		go run .\task.go -f 1-2 -d ":" -s
		go run .\task.go -f 1,2,3 -d ":" -s
		go run .\task.go -f 1 -d ":" -s
*/

func main() {
	var f *string = flag.String("f", "", "Выбрать поля")
	var d *string = flag.String("d", "\t", "Использовать другой разделитель")
	var s *bool = flag.Bool("s", false, "Только строки с разделителем")
	flag.Parse()
	Start(f, d, s)
}

// Start точка входа
func Start(f *string, d *string, s *bool) {
	// Сюда будем записывать номера колонок, которые хотим вывести
	var fields []int = make([]int, 0)

	// Парсим параметр -f
	if strings.Contains(*f, ",") {
		arrInts := strings.Split(*f, ",")
		for _, v := range arrInts {
			field, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("Неправильный аргумент")
				os.Exit(1)
			}
			fields = append(fields, field)
		}
	} else if strings.Contains(*f, "-") {
		arrInts := strings.Split(*f, "-")

		start, err := strconv.Atoi(arrInts[0])
		if err != nil {
			fmt.Printf("Неправильный аргумент")
			os.Exit(1)
		}
		end, err := strconv.Atoi(arrInts[1])
		if err != nil {
			fmt.Printf("Неправильный аргумент")
			os.Exit(1)
		}
		for i := start; i <= end; i++ {
			fields = append(fields, i)
		}
	} else if val, err := strconv.Atoi(*f); err == nil {
		fields = append(fields, val)
	} else {
		fmt.Printf("Неправильный аргумент")
		os.Exit(1)
	}

	// Считываем ввод из STFIN
	var lines []string = make([]string, 0)

	fmt.Println("Введите строки. Для ввода следующей строки нажмите Enter. Для прекращения ввода ctrl+c")
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	var Print handler = &handlerPrint{}
	var Fields handler = &handlerFields{
		Field:     fields,
		Delimiter: *d,
		Next:      Print,
	}
	var Separated handler = &handlerSeparated{
		Separated: *s,
		Delimiter: *d,
		Next:      Fields,
	}
	Separated.SendRequest(lines)
}

type handler interface {
	SendRequest(str []string)
}

type handlerSeparated struct {
	Separated bool
	Delimiter string
	Next      handler
}

// Определяем какие строки не содержат разделителя
func (s handlerSeparated) SendRequest(str []string) {
	if s.Separated {
		var result []string = make([]string, 0)
		for _, value := range str {
			if strings.Contains(value, s.Delimiter) {
				result = append(result, value)
			}
		}
		str = result
	}
	s.Next.SendRequest(str)
}

type handlerFields struct {
	Field     []int
	Delimiter string
	Next      handler
}

// Вырезаем подстроки согласно параметрам
func (f handlerFields) SendRequest(str []string) {
	for index, value := range str {
		if !strings.Contains(value, f.Delimiter) {
			continue
		}
		// Делим строку по разделителям
		substr := strings.Split(value, f.Delimiter)
		var lenDubstr int = len(substr)
		// Сюда запишем подстроки, номера которых указаны в параметре -f
		var resultSubstring []string = make([]string, 0)
		for _, v := range f.Field {
			if v-1 > lenDubstr {
				continue
			}
			resultSubstring = append(resultSubstring, substr[v-1])
		}
		// Объеденяем массив строк в строку
		str[index] = strings.Join(resultSubstring, f.Delimiter)
	}
	f.Next.SendRequest(str)
}

// Выводим результат
type handlerPrint struct{}

func (s handlerPrint) SendRequest(str []string) {
	fmt.Println("Результат выполнения:")
	for _, val := range str {
		fmt.Println(val)
	}
}

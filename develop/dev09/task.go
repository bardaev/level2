package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func main() {
	var sourceUrlArg string = os.Args[1]

	// Указываем ссылку
	// var u string = "https://www.google.ru/"
	sourceUrl, err := url.Parse(sourceUrlArg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(sourceUrl)

	// Создаем файл страницы
	file, errFile := os.Create("index.html")
	if errFile != nil {
		fmt.Println(errFile)
		os.Exit(1)
	}

	// Делаем запрос к сайту
	resp, errGet := http.Get(sourceUrl.String())
	if errGet != nil {
		fmt.Println(errGet)
		os.Exit(1)
	}

	defer resp.Body.Close()

	// Считываем тело ответа и записываем в буфер
	var page bytes.Buffer = bytes.Buffer{}
	bufFile := bufio.NewWriter(file)
	for {
		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)
		page.Write(bs)
		if n == 0 || err != nil {
			break
		}
	}
	bufFile.Flush()
	defer file.Close()

	var pageBytes []byte = page.Bytes()

	// Создаем регулярное выражение для поиска ссылок
	link := regexp.MustCompile("href=('|\\\")([\\s\\S]+?)\"")
	all := link.FindAllIndex(pageBytes, -1)

	// Меняем относительные ссылки на абсолютные
	for i := len(all) - 1; i > 0; i-- {
		var result []byte = downloadAssets(all[i][0], all[i][1], pageBytes, sourceUrl)
		if result != nil {
			pageBytes = result
		}
	}

	// Записываем данные в файл
	bufFile.Write(pageBytes)
	bufFile.Flush()
}

func downloadAssets(start int, end int, page []byte, sourceUrl *url.URL) []byte {
	var hrefByte []byte = make([]byte, len(page[start:end]))
	copy(hrefByte, page[start:end])
	var hrefString string
	hrefString = strings.TrimPrefix(hrefString, "href=\"")
	hrefString = strings.TrimSuffix(hrefString, "\"")
	href, err := url.Parse(hrefString)
	if err != nil {
		return nil
	}
	if href.Scheme == "" || href.Host == "" {
		href.Scheme = sourceUrl.Scheme
		href.Host = sourceUrl.Host
	}
	var absoluteLink []byte = bytes.Replace(hrefByte, []byte(hrefString), []byte(href.String()), 1)
	var result []byte = bytes.Replace(page, hrefByte, absoluteLink, 1)
	return result
}

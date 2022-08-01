package pattern

import "fmt"

/*
Описание:
	Фасад - паттерн, который инкапсулирует сложную логику и предоставляет упрощенный интерфейс.

Применение:
	Может применяться, когда, например, есть библиотека со множеством объектов,
	которые необходимо инициализировать, следить за правильным порядком зависимостей и т.п.,
	и вся эта логика оборачивается в простые методы, которые инкапсулируют всю сложную работу.

Плюсы:
	- Изоляция компонентов от сложной логики

Минусы:
	- Привязка объекта фасада ко всем классам программы
*/

// Фасад
type VideoConverter struct{}

func (v *VideoConverter) Convert(fileName string) string {
	file := videoFile{}
	cf := codecFactory{}
	sourceCodec := cf.extract(VideoConverter{})
	destinationCodec := mpeg4CompressionCodec{}
	fmt.Print(sourceCodec)
	br := bitrateReader{}
	buffer := br.read(file)
	fmt.Println(buffer)
	convert := br.convert(destinationCodec)

	fmt.Println(convert)
	ax := audioMixer{}
	result := ax.fix(fileName)

	return result
}

// Имитация сложной библиотеки

type videoFile struct{}

type mpeg4CompressionCodec struct{}

type codecFactory struct{}

func (c *codecFactory) extract(v VideoConverter) string {
	return "codec"
}

type bitrateReader struct{}

func (b *bitrateReader) read(fileName videoFile) string {
	return "buffer"
}

func (b *bitrateReader) convert(fileName mpeg4CompressionCodec) string {
	return "result"
}

type audioMixer struct{}

func (a *audioMixer) fix(fileName string) string {
	return "fix"
}

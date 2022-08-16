Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error

Данный вывод объясняется тем, что переменной err присваетвается конкретный тип.
То есть до присваиявания err указывает на конкретный тип nil и на значение nil, а после присваивания конкретный тип равен customError, а значение nil.
Соответственно, чтобы программа вывела ок, конкретный тип и значение должны быть nil. Как, например, до присваивания.

```
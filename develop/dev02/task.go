package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// fmt.Println(unpackString(""))
	stringToTrim := "yrdyd efwpow"
	fmt.Println(unpackString(stringToTrim))
}

func unpackString(str string) (string, error) {
	if len(strings.Split(str, " ")) != 1 {
		return "", fmt.Errorf("В строке присутствуют пробелы")
	}

	if len(str) == 0 {
		return str, nil
	}

	if _, err := strconv.Atoi(string(str[0])); err == nil {
		return "", fmt.Errorf("некорректная строка")
	}

	var sb strings.Builder
	// "a4bc2d5e" => "aaaabccddddde"
	for i := 0; i < len(str); {
		// если последний символ в строке
		if i == len(str)-1 {
			if _, err := strconv.Atoi(string(str[0])); err == nil {
				return "", fmt.Errorf("некорректная строка")
			}

			sb.WriteByte(str[i])
			i++
			continue
		}

		// если после символа не цифра то просто добавляем в итоговую строку
		v, err := strconv.Atoi(string(str[i+1]))
		if err != nil {
			sb.WriteByte(str[i])
			i++
			continue
		}

		// добавляем в итоговую строку символ в количестве в зависимости от цифры
		for j := 0; j < v; j++ {
			sb.WriteByte(str[i])
		}

		i++
		i++
	}

	return sb.String(), nil
}

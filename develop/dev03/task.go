package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки в файле по аналогии с консольной утилитой sort
(man sort — смотрим описание и основные параметры): на входе подается файл из несортированными строками,
на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительно

Реализовать поддержку утилитой следующих ключей:

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учетом суффиксов

*/

const (
	readPath  = "text.txt"
	writePath = "result.txt"
)

type flags map[string]interface{}

func main() {

	f := readFlags()

	str, err := readStrings(readPath)
	if err != nil {
		panic(err)
	}

	r, err := sortStrings(str, f)
	if err != nil {
		panic(err)
	}

	fmt.Println("res", r)

	err = writeStrings(writePath, r)
	if err != nil {
		panic(err)
	}
}

func readFlags() flags {
	var flags flags = make(flags)

	k := flag.Int("k", 1, "колонка для сортировки")
	n := flag.Bool("n", false, "сортировать по числовому значению")
	r := flag.Bool("r", false, "сортировать в обратном порядке")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")

	flag.Parse()

	if *k == 0 {
		*k = 1
	}

	flags["k"] = *k
	fmt.Println("k ", k)

	flags["n"] = *n
	fmt.Println("n ", n)

	flags["r"] = *r
	fmt.Println("r ", r)

	flags["u"] = *u
	fmt.Println("u ", u)

	fmt.Println(flags)

	return flags
}

// Вызывает функции в зависимости от флагов
func sortStrings(str []string, flags flags) ([]string, error) {
	var s []string = make([]string, len(str))
	copy(s, str)

	column := flags["k"].(int)

	isInt := flags["n"].(bool)

	if isInt {
		var err error = nil
		s, err = sortInt(column, s)
		if err != nil {
			return nil, err
		}
	} else {
		s = sortStr(column, s)
	}

	fmt.Println(s)

	if u := flags["u"]; u.(bool) {
		s = deleteRepeat(s)
	}

	if r := flags["r"]; r.(bool) {
		reverse(s)
	}

	return s, nil
}

func deleteRepeat(in []string) []string {
	var allKeys map[string]struct{} = make(map[string]struct{})
	var out []string

	for _, item := range in {
		if _, value := allKeys[item]; !value {
			allKeys[item] = struct{}{}
			out = append(out, item)
		}
	}

	return out
}

func reverse(in []string) {
	inputLen := len(in)
	inputMid := inputLen / 2

	for i := 0; i < inputMid; i++ {
		j := inputLen - i - 1

		in[i], in[j] = in[j], in[i]
	}
}

// Сортирует строки
func sortStr(column int, str []string) []string {
	// результирующий слайс
	var res []string = make([]string, 0, len(str))
	// мапа чтобы доставать строки по слову из колонки
	var m map[string]string = map[string]string{}
	// чтобы отсортировать слова из колонок
	var strs []string = make([]string, 0, 5)

	// слайс для сортировки строк у которых меньше колонок
	start := make([]string, 0, 10)

	// распределяем по слайсам в зависимости от количества колонок
	for _, v := range str {
		s := strings.Split(v, " ")
		if len(s) < column {
			start = append(start, v)
		} else {
			i := s[column-1]

			m[i] = v
			strs = append(strs, i)
		}
	}

	// сортируем начальный слайс и добавляем в результирующий
	sort.Strings(start)
	res = append(res, start...)

	// сортируем конечный слайс и добавляем в результирующий
	sort.Strings(strs)
	for _, v := range strs {
		res = append(res, m[v])
	}

	return res
}

func writeStrings(path string, str []string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Unable to create file: %s", err.Error())
	}

	defer file.Close()

	for i, v := range str {
		file.WriteString(v)

		if i != len(str)-1 {
			file.WriteString("\n")
		}
	}

	return nil
}

func readStrings(path string) ([]string, error) {
	var result []string

	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Error open file: %s", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		result = append(result, sc.Text())
	}

	return result, nil
}

// Сортируем числа
func sortInt(column int, str []string) ([]string, error) {
	// результирующий слайс
	var res []string = make([]string, 0, len(str))
	// мапа чтобы доставать строки по слову из колонки
	var m map[int]string = map[int]string{}
	// чтобы отсортировать слова из колонок
	var ints []int = make([]int, 0, 5)

	// слайс для сортировки строк у которых меньше колонок
	start := make([]string, 0, 10)

	// распределяем по слайсам в зависимости от количества колонок
	for _, v := range str {
		s := strings.Split(v, " ")
		if len(s) < column {
			start = append(start, v)
		} else {
			i, err := strconv.Atoi(s[column-1])
			if err != nil {
				return nil, err
			}

			m[i] = v
			ints = append(ints, i)
		}
	}

	// сортируем начальный слайс и добавляем в результирующий
	sort.Strings(start)
	res = append(res, start...)

	// сортируем конечный слайс и добавляем в результирующий
	sort.Ints(ints)
	for _, v := range ints {
		res = append(res, m[v])
	}

	return res, nil
}

package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// type

func main() {
	s := []string{"пяТак", "пятка", "пятка", "типок", "тяпка", "листок", "слиток", "столик", "тест"}

	// m := []string{"a", "b", "c", "d", "z", "w"}
	// for _, v := range m {
	// 	fmt.Println(v)
	// }
	r := findAnagrams(s)

	for i, v := range *r {
		fmt.Println(i, v)
	}
}

func findAnagrams(words []string) *map[string]*[]string {
	// результирующая мапа
	var result map[string]*[]string = map[string]*[]string{}
	// вспомогательная мапа для хранения байтов из слов
	r := make(map[string]*AnogramHelp)

	// такой же слайс как и входной, но в нижнем регистре и без повторов слов
	w := make([]string, 0, len(words))

	// заполняем слайс нижнего регистра и вспомогательную мапу
	for _, v := range words {
		lower := strings.ToLower(v)

		_, ok := r[lower]
		if ok {
			continue
		}

		w = append(w, lower)

		ah := NewAnagramHelp(lower)
		r[lower] = ah
	}

	// сортируем слайс в нижнем регистре
	sort.Strings(w)

	// итерация по слайсу со словами в нижнем регистре
	for _, word := range w {
		// переменная для проверки нужно ли создавать новое множество
		isNeedNewSlice := true

		// итерация по уже существующим ключам в результирующей мапе
		for k := range result {

			// проверяем длинну байтов ключа существующего множества с переданным словом
			if len(r[k].SortedByte) == len(r[word].SortedByte) {
				// проверяем равны ли байты ключа существующего множества с переданным словом
				if bytes.Equal(r[k].SortedByte, r[word].SortedByte) {
					isNeedNewSlice = false
					// добавляем новое значение в уже существующее множество
					r := *result[k]
					r = append(r, word)
					result[k] = &r

					break
				}
			}
		}

		// проверяем нужно ли создать новое множество под входящее слово
		if isNeedNewSlice {
			s := result[word]
			s = makeSlice(word)
			result[word] = s
		}
	}

	// убираем множества с одним словом
	for k, v := range result {
		s := *v

		if len(s) < 2 {
			delete(result, k)
		}
	}

	return &result
}

func makeSlice(v string) *[]string {
	s := make([]string, 0, 2)
	s = append(s, v)
	return &s
}

type AnogramHelp struct {
	SortedByte []byte
}

func NewAnagramHelp(s string) *AnogramHelp {
	// добавляем в структуру отсортированные байты
	sorting := []byte(s)
	sort.Slice(sorting, func(i, j int) bool {
		return sorting[i] < sorting[j]
	})

	return &AnogramHelp{SortedByte: sorting}
}

// func findAnagrams(words []string) *map[string]*[]string {
// 	// 1
// 	// existLetters := make(map[string]map[string]int)
// 	// result := make(map[string]*[]string)

// 	result := make(map[string][]*AnagramHelp)

// 	anagramHelps := make([]*AnagramHelp, 0, len(words))

// 	for _, v := range words {
// 		sLower := strings.ToLower(v)
// 		h := NewAnagramHelp(sLower)
// 		anagramHelps = append(anagramHelps, h)
// 	}

// 	// for i := 0; i < len(anagramHelps); i++ {
// 	// 	ah := anagramHelps[i]

// 	// }

// 	// "пятак", "пятак","пятка", "тяпка", "типок","листок", "слиток", "столик", "тест"
// 	for _, ah := range anagramHelps {
// 		if len(result) == 0 {
// 			result[ah.String()] = []*AnagramHelp{ah}
// 			continue
// 		}

// 		for _, v := range result {
// 			rAh := v[0]

// 			if rAh.AllRunesCount != ah.AllRunesCount {
// 				result[ah.String()] = []*AnagramHelp{ah}
// 				continue
// 			}

// 			for i, _ := range ah.EachRunesCount {
// 				_, ok := rAh.EachRunesCount[i]
// 				if !ok {
// 					result[ah.String()] = []*AnagramHelp{ah}
// 					continue
// 				}

// 				ahs := result[rAh.String()]
// 				ahs = append(ahs, ah)
// 				result[rAh.String()] = ahs
// 			}
// 		}
// 	}

// 	return nil
// }

// type AnagramHelp struct {
// 	AllRunesCount  int
// 	EachRunesCount map[rune]int
// 	runes          []rune
// }

// func NewAnagramHelp(s string) *AnagramHelp {
// 	allRunesCount := len([]rune(s))
// 	eachRunesCount := make(map[rune]int)

// 	for _, v := range s {
// 		_, ok := eachRunesCount[v]
// 		if !ok {
// 			eachRunesCount[v] = 0
// 		}

// 		eachRunesCount[v]++
// 	}

// 	return &AnagramHelp{AllRunesCount: allRunesCount, runes: []rune(s), EachRunesCount: eachRunesCount}
// }

// func (ah *AnagramHelp) String() string {
// 	return string(ah.runes)
// }

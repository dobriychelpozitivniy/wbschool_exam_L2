package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"flag"
	"fmt"
	"strings"
)

type Flags map[string]interface{}

func main() {
	// 	9 ::: kek eadiajodjeafh b
	// 12 ::: kek eadiajodjeafh b
	// 15 ::: kek eadiajodjeafh b
	// 18 ::: kek eadiajodjeafh b
	// 21 ::: kek eadiajodjeafh b
	// 24 ::: kek eadiajodjeafh b
	// 27 ::: kek eadiajodjeafh b

	s := []string{
		"aaaaaaaaaaaaaa",
		"bbbbbbbbbbbbb",
		"cccccccccccc",
		"wwwwwwwwwwwwwwwww",
		"qqqqqqqqqqqqqqq",
		"eeeeeeeeeeeeeeeeeee",
		"python",
		"python sefsfdues sbeifsbei dnjsf",
		"asiudhiwudsbqwdpython",
		"kek eadiajodjeafh b",
		"python sefsfdues sbeifsbei dnjsf",
		"asiudhiwudsbqwdpython",
		"KEK eadiajodjeafh b",
		"python sefsfdues sbeifsbei dnjsf",
		"asiudhiwudsbqwdpython",
		"kek eadiajodjeafh b",
		"python sefsfdues sbeifsbei dnjsf",
		"asiudhiwudsbqwdpython check",
		"kek eadiajodjeafh b",
		"python sefsfdues sbeifsbei dnjsf",
		"wkudhiqwudhqiwudhqiw",
		"KEK eadiajodjeafh b",
		"qdlwijoqwdjoqwuhe2189e",
		"wkudhiqwudhqiwudhqiw",
		"kek eadiajodjeafh b",
		"qdlwijoqwdjoqwuhe2189e",
		"wkudhiqwudhqiwudhqiw",
		"kek eadiajodjeafh b",
		"qdlwijoqwdjoqwuhe2189e",
		"wkudhiqwudhqiwudhqiw",
		"wkudhiqwudhqiwudhqiw",
		"wkudhiqwudhqiwudhqiw",
		"wkudhiqwudhqiwudhqiw",
		"wkudhiqwudhqiwudhqiw",
	}

	fs := initFlags()

	res := grep(s, "kek", fs)

	for _, v := range res {
		fmt.Println(v)
	}
}

func initFlags() Flags {
	var fs Flags = make(Flags)

	A := flag.Int("A", 0, "печатать +N строк после совпадения")
	B := flag.Int("B", 0, "печатать +N строк до совпадения")
	C := flag.Int("C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	c := flag.Int("c", 0, "(количество строк)")
	i := flag.Bool("i", false, "(игнорировать регистр)")
	v := flag.Bool("v", false, "(вместо совпадения, исключать)")
	F := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	n := flag.Bool("n", false, "печатать номер строки")

	flag.Parse()

	fs["A"] = *A
	fs["B"] = *B
	fs["C"] = *C
	fs["c"] = *c
	fs["i"] = *i
	fs["v"] = *v
	fs["F"] = *F
	fs["n"] = *n

	fmt.Println(fs)

	return fs
}

func grep(s []string, f string, flags Flags) []string {
	var indexes []int
	var strs []string = make([]string, len(s))
	var res []string = make([]string, 0)

	i := flags["i"].(bool)
	if i {
		f = strings.ToLower(f)

		for i, v := range s {
			strs[i] = strings.ToLower(v)
		}
	} else {
		copy(strs, s)
	}

	c := flags["c"].(int)
	v := flags["v"].(bool)
	F := flags["F"].(bool)

	if F {
		if v {
			indexes = findIndexesVF(strs, f, c)
		} else {
			indexes = findIndexesF(strs, f, c)
		}
	} else {
		if v {
			indexes = findIndexesV(strs, f, c)
		} else {
			indexes = findIndexes(strs, f, c)
		}
	}

	fmt.Println(indexes)

	C := flags["C"].(int)
	if C > 0 {
		indexes = after(indexes, C, len(s))
		indexes = before(indexes, C)
	} else {
		A := flags["A"].(int)
		if A > 0 {
			indexes = after(indexes, A, len(s))
		}

		B := flags["B"].(int)
		if B > 0 {
			indexes = before(indexes, B)
		}
	}

	res = filter(s, indexes)

	if len(res) == 0 {
		fmt.Println("Совпадения не найдены.")

		return nil
	}

	n := flags["n"].(bool)
	if n {
		for i, v := range res {
			res[i] = fmt.Sprintf("%d ::: %s", indexes[i]+1, v)
		}
	}

	return res
}

func findIndexesVF(strs []string, f string, count int) []int {
	var indexes []int = make([]int, 0, 5)

	j := 0
	for i, v := range strs {
		if count > 0 {
			if j == count {
				break
			}
		}

		if !(v == f) {
			indexes = append(indexes, i)
			j++
		}
	}

	return indexes
}

func findIndexesF(strs []string, f string, count int) []int {
	var indexes []int = make([]int, 0, 5)

	j := 0
	for i, v := range strs {
		if count > 0 {
			if j == count {
				break
			}
		}

		if v == f {
			indexes = append(indexes, i)
			j++
		}
	}

	return indexes
}

func findIndexes(strs []string, f string, count int) []int {
	var indexes []int = make([]int, 0, 5)

	j := 0
	for i, v := range strs {
		if count > 0 {
			if j == count {
				break
			}
		}

		if strings.Contains(v, f) {
			indexes = append(indexes, i)
			j++
		}
	}

	return indexes
}

func findIndexesV(strs []string, f string, count int) []int {
	var indexes []int = make([]int, 0, 5)

	j := 0
	for i, v := range strs {
		if count > 0 {
			if j == count {
				break
			}
		}

		if !strings.Contains(v, f) {
			indexes = append(indexes, i)
			j++
		}
	}

	return indexes
}

// func excludeFilter(s []string, idxs []int) []string {
// 	var res []string = make([]string, 0, len(idxs))
// 	var midxs map[int]struct{} = make(map[int]struct{})

// 	for _, v := range idxs {
// 		midxs[v] = struct{}{}
// 	}

// 	for i, v := range s {
// 		if _, ok := midxs[i]; !ok {
// 			res = append(res, v)
// 		}
// 	}

// 	return res
// }

func filter(s []string, idxs []int) []string {
	var res []string = make([]string, 0, len(idxs))
	var midxs map[int]struct{} = make(map[int]struct{})

	for _, v := range idxs {
		midxs[v] = struct{}{}
	}

	for i, v := range s {
		if _, ok := midxs[i]; ok {
			res = append(res, v)
		}
	}

	return res
}

func after(idxs []int, n int, maxIndex int) []int {
	res := make([]int, 0, len(idxs)+n)
	res = append(res, idxs...)

	iLast := idxs[len(idxs)-1]
	for i := 1; i < n+1; i++ {
		if iLast+i == maxIndex {
			break
		}

		res = append(res, iLast+i)
	}

	return res

}

func before(idxs []int, n int) []int {
	res := make([]int, 0, len(idxs)+n)

	iFirst := idxs[0] - n
	iEnd := idxs[0]
	for i := iFirst; i < iEnd; i++ {
		if i < 0 {
			continue
		}

		res = append(res, i)
	}

	res = append(res, idxs...)

	return res
}

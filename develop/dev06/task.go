package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Flags struct {
	f map[int]struct{}
	d string
	s bool
}

func main() {
	var s []string

	flags := initFlags()

	fmt.Println(fmt.Sprintf("%#v", flags))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	in := os.Stdin

	go func() {
		<-c
		in.Close()
	}()

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	res := cut(s, flags)
	for _, v := range res {
		fmt.Println(v)
	}

	// file, err := os.Open("text.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer func() {
	// 	if err = file.Close(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// fmt.Println(a, w)

	// b, err := ioutil.ReadAll(file)
	// s := string(b)

	// fmt.Println(name)

	// s := cut("text.txt", flags)
}

func cut(strs []string, flags Flags) []string {
	var res []string = make([]string, 0, 10)

	for _, v := range strs {
		s := getParsedString(v, flags)
		if s != "" {
			res = append(res, s)
		}
	}

	return res
}

func getParsedString(s string, flags Flags) string {
	var temp []string = make([]string, 0, 5)
	var result string

	ss := strings.Split(s, flags.d)

	if flags.s {
		if len(ss) == 1 {
			return ""
		}
	} else {
		if len(ss) == 1 {
			return ss[0]
		}
	}

	for i, v := range ss {
		if _, ok := flags.f[i+1]; ok {
			// fmt.Println(v)
			temp = append(temp, v)
		}
	}

	result = strings.Join(temp, flags.d)

	return result
}

func parseFlagF(f string) map[int]struct{} {
	var res map[int]struct{} = make(map[int]struct{})

	s := strings.Split(f, ",")

	for _, v := range s {
		ss := strings.Split(v, "-")

		if len(ss) == 1 {
			i, err := strconv.Atoi(v)
			if err != nil {
				panic(fmt.Sprintf("Error parse f flag: %s", err))
			}

			res[i] = struct{}{}
		} else {
			if len(ss) != 2 {
				panic(fmt.Sprintf("Error parse f flag"))
			}

			start, err := strconv.Atoi(ss[0])
			if err != nil {
				panic(fmt.Sprintf("Error parse f flag: %s", err))
			}

			end, err := strconv.Atoi(ss[1])
			if err != nil {
				panic(fmt.Sprintf("Error parse f flag: %s", err))
			}

			r := getRange(uint(start), uint(end))

			for _, v := range r {
				res[int(v)] = struct{}{}
			}
		}
	}

	return res
}

func getRange(start uint, end uint) (result []uint) {
	result = make([]uint, 0, end-start)
	for value := start; value <= end; value++ {
		result = append(result, value)
	}

	return result
}

func initFlags() Flags {
	f := flag.String("f", "", "выбрать поля (колонки)")
	d := flag.String("d", "\t", "использоваться другой разделитель")
	s := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	if len(*f) <= 0 {
		panic("Необходимо передать флаг -f ( выбор поля (колонки) )")
	}

	return Flags{
		f: parseFlagF(*f),
		d: *d,
		s: *s,
	}
}

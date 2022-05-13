package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

const downDir = "/site/"
const siteDir = "/site/src/"

func getLinks(url string) map[string][]string {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}

	body := resp.Body

	var HTMLlinks []string
	var fileLinks []string
	var links map[string][]string = make(map[string][]string)
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			//todo: links list shoudn't contain duplicates
			links["html"] = HTMLlinks
			links["file"] = fileLinks

			for i, v := range links {
				if i == "html" {
					for _, w := range v {
						fmt.Println("HTML URL:", w)
					}
				}
				if i == "file" {
					for _, w := range v {
						fmt.Println("FILE URL:", w)
					}
				}
			}

			return links
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			fmt.Println("TOKEN: ", token.Data)
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						HTMLlinks = append(HTMLlinks, attr.Val)
					}

				}
			}
			if "link" == token.Data {
				fmt.Println("LINK: ", token.Attr)
				for _, attr := range token.Attr {
					fmt.Println("ATTR: ", attr)
					if attr.Key == "href" {
						fileLinks = append(fileLinks, attr.Val)
					}

				}
			}
			if "script" == token.Data {

				for _, attr := range token.Attr {
					if attr.Key == "src" {
						fileLinks = append(fileLinks, attr.Val)
					}

				}
			}
			if "img" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "src" {
						fileLinks = append(fileLinks, attr.Val)
					}

				}
			}

		}
	}
}

func parseLinks(l []string) []string {
	var res []string

	for _, v := range l {
		if strings.HasPrefix(v, "http") ||
			strings.HasPrefix(v, "https") ||
			v == "/" ||
			strings.Contains(v, "#") {
			continue
		}

		res = append(res, v)
	}

	return res
}

func downloadHTML(wg *sync.WaitGroup, url string, path string) error {
	rPath := []rune(path)
	if path != "" {
		a := '/'
		if rPath[len(rPath)-1] != a {
			rPath = append(rPath, a)
		}
	}

	filename := "index.html"
	// fmt.Println("FILENAME", filename)
	// fmt.Println("DIR", dir)

	pwd, err := os.Getwd()
	err = os.MkdirAll(pwd+siteDir+string(rPath), 0755)
	if err != nil {
		fmt.Printf("ERR MKDIR: %s", err)
	}

	go func() error {
		defer wg.Done()

		resp, err := http.Get(url + path)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		f, err := os.Create(pwd + siteDir + string(rPath) + filename)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(f, resp.Body)

		return nil
	}()

	return nil
}

func downloadFile(wg *sync.WaitGroup, url string, path string) error {
	// var filename string

	// fmt.Println("IN URL: ", url)
	// fmt.Println("IN PATH: ", path)

	rPath := []rune(path)

	if rPath[0] == '/' {
		rPath = rPath[1:]
	}

	temp := []rune{}
	for i := len(rPath) - 1; ; i-- {
		if rPath[i] == '/' {
			break
		}
		temp = append(temp, rPath[i])
	}

	path = string(rPath)

	l := strings.Split(path, "/")
	// for i, v := range l {
	// 	fmt.Println("I:", i, v)
	// }

	// fmt.Println("PATH AFTER: ", path)

	filename := l[len(l)-1]
	dir := "/" + strings.Join(l[:len(l)-1], "/") + "/"

	// fmt.Println("FILENAME", filename)
	// fmt.Println("DIR", dir)

	pwd, err := os.Getwd()
	err = os.MkdirAll(pwd+siteDir+dir, 0755)
	if err != nil {
		fmt.Printf("ERR MKDIR: %s", err)
	}

	go func() error {
		defer wg.Done()

		resp, err := http.Get(url + path)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		f, err := os.Create(pwd + siteDir + dir + filename)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(f, resp.Body)

		return nil
	}()

	return nil
}

func main() {
	pUrl := flag.String("url", "https://turbosolution.ru/", "URL to be processed")
	flag.Parse()
	url := *pUrl
	if url == "" {
		fmt.Fprintf(os.Stderr, "Error: empty URL!\n")
		return
	}

	filename := "index.html"
	l := getLinks(url)

	// for i, v := range l {
	// 	fmt.Println("I:", i, v)
	// }

	for i, v := range l {
		l[i] = parseLinks(v)
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	downloadHTML(&wg, url, "")

	wg.Add(len(l["file"]))
	wg.Add(len(l["html"]))
	for i, v := range l {
		// wg.Add(len(v))
		if i == "file" {
			for _, fPath := range v {
				// fmt.Println("URL: ", url)
				// fmt.Println("URL: ", fPath)
				downloadFile(&wg, url, fPath)
				// fmt.Println(fPath)
			}
		}
		if i == "html" {
			for _, htmlPath := range v {
				// fmt.Println("URL: ", url)
				// fmt.Println("URL: ", htmlPath)
				downloadHTML(&wg, url, htmlPath)
				// fmt.Println(htmlPath)
			}
		}
	}

	wg.Wait()
	fmt.Println("Checking if " + filename + " exists ...")
	// if _, err := os.Stat(filename); os.IsNotExist(err) {
	// 	err := download(url, "", filename)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(filename + " saved!")
	// } else {
	// 	fmt.Println(filename + " already exists!")
	// }
}

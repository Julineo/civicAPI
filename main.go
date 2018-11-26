package main

import (
	"bufio"
	"fmt"
	"os"
	"net/http"
	"net/url"
	"log"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	key := os.Args[1]
	inputFile := os.Args[2]

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	address := ""

	//passing first string
	scanner.Scan()

	for scanner.Scan() {
		address = scanner.Text()


		fmt.Printf("%v,", address)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		p := url.Values{"fields": {"divisions,normalizedInput"}, "key": {key}, "address": {address}}
		resp, err := http.Get("https://www.googleapis.com/civicinfo/v2/representatives?" + p.Encode())

		if err != nil {
			log.Fatal(err)
		}

		rs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	//	fmt.Printf("%s\n", rs)

		srs := string(rs)
		t := strings.Index(srs, "congress")
		if t == -1 {
			fmt.Println("404")
			continue
		}
		s, e := -1, -1
		for i := t; i > 0; i-- {
			if srs[i] == 34 {
				s = i
				break
			}
		}
		for i := t; i < t + 100; i++ {
			if srs[i] == 34 {
				e = i
				break
			}
		}
		fmt.Println(string(rs[s+1 : e]))
		resp.Body.Close()
		time.Sleep(1 * time.Second)
	}
}


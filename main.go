package main

import (
	"bufio"
	"fmt"
	"os"
	"net/http"
	"net/url"
	"log"
	"io/ioutil"
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

	fmt.Println("[")
	for scanner.Scan() {
		address = scanner.Text()


		//fmt.Printf("%v\n", address)

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
		fmt.Printf("%s\n", rs)
		fmt.Println(",")
		resp.Body.Close()
	}
	fmt.Println("{}]")
}


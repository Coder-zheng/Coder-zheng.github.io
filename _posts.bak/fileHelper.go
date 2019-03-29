package main

import (
	"io/ioutil"
	"log"
	"strconv"
)

const (
	date  = "2019-4"
	count = 28
)

func main() {
	buffer, err := ioutil.ReadFile("./template")
	for index := 1; index <= count; index++ {
		err = ioutil.WriteFile(`./`+date+`/`+date+`-`+strconv.Itoa(index)+`.md`, buffer, 0644)
	}
	if err != nil {
		log.Fatal(err)
	}
}

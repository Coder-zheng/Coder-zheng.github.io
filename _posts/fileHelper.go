package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"
)

const DATE_FORMAT = "2006-01-02"

func main() {
	name := os.Args[1]
	categories := os.Args[2]
	tags := os.Args[3]

	buffer := []byte(`---

layout: post
title: "` + name + `"
categories: ` + categories + `
tags:  ` + tags + `
author: ant

---

* content
{:toc}

`)
	now := time.Now().Format(DATE_FORMAT)
	err := ioutil.WriteFile(now+`-`+name+`.md`, buffer, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

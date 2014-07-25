package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type Site struct {
	Name string
	Url  string
}

type SiteList struct {
	Sites []Site
}

func main() {
	os.Exit(realMain())
}

func realMain() int {
	args := os.Args[1:]

	if len(args) < 1 {
		usage()
		return 1
	}
	fileName := args[0]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return 1
	}
	defer file.Close()
	siteList := parseSites(file)
	fmt.Println("processing", len(siteList.Sites), "sites")

	var wg sync.WaitGroup
	wg.Add(len(siteList.Sites))

	for _, site := range siteList.Sites {
		fmt.Println("testing:", site.Name)
		go hammer(site, &wg)
	}

	wg.Wait()
	return 0
}

func parseSites(file *os.File) SiteList {
	decoder := json.NewDecoder(file)
	siteList := SiteList{}
	err := decoder.Decode(&siteList)

	if err != nil {
		fmt.Println("error:", err)
	}

	return siteList
}

func hammer(site Site, wg *sync.WaitGroup) {
	startTime := time.Now()
	response, err := http.Get(site.Url)

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		fmt.Println(site.Name, "responded in:", time.Since(startTime))
		defer response.Body.Close()
		_, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
	}

	wg.Done()
}

func usage() {
  fmt.Println("Usage: gogh config.json")
}

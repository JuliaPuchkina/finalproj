package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}

	newsAllPage("1")
}

func newsAllPage(n string) {
	link := "http://localhost:8081/news?page=" + url.QueryEscape(n)

	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

}

func newsTitleSearch(n string) {
	link := "http://localhost:8081/news?title=" + url.QueryEscape(n)

	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

func newsDetailed() {}

func addComment() {}

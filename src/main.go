package main

import (
	"os"
	"fmt"
	"net/url"
	"net/http"
	"io/ioutil"
	"path/filepath"
	"github.com/webview/webview"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	exe, err := os.Executable()
	if err != nil {
		os.Exit(1)
	}

	w := webview.New(true)
	htmlFile := filepath.Join(filepath.Dir(exe), "static", "index.html")
	resultFile := filepath.Join(filepath.Dir(exe), "result.txt")
	w.SetTitle("WebView Example")
	w.SetSize(800, 600, webview.HintNone)
	bytes, _ := ioutil.ReadFile(htmlFile)
	w.Bind("getTitle", func(url string) string {
		result, err := getTitle(url)
		if err != nil {
			w.Dispatch(func() {
				writeResult(fmt.Sprint(err), resultFile)
			})
			return ""
		} else {
			w.Dispatch(func() {
				writeResult(result, resultFile)
			})
			return result
		}
	})
	w.Navigate("data:text/html," + url.PathEscape(string(bytes)))
	w.Run()
}

func getTitle(url string)(string, error) {
	var err error
	var res string
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return res, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return res, err
	}
	return doc.Find("title").Text(), nil
}

func writeResult(result string, file string) {
	if err := ioutil.WriteFile(file, []byte(result), 0666); err != nil {
		os.Exit(1)
	}
}

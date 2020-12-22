package main

import (
	"os"
	"fmt"
	"net/url"
	"io/ioutil"
	"path/filepath"
	"github.com/webview/webview"
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
	w.Bind("response", func(s string) {
		w.Dispatch(func() {
			err := ioutil.WriteFile(resultFile, []byte("[test001] " + s), 0666)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		})
	})
	w.Navigate("data:text/html," + url.PathEscape(string(bytes)))
	w.Run()
}

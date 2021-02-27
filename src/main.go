package main

import (
	"os"
	"fmt"
	"embed"
	"net/url"
	"path/filepath"
	"github.com/webview/webview"
)

//go:embed static
var staticFS embed.FS

func main() {
	exe, err := os.Executable()
	if err != nil {
		os.Exit(1)
	}

	w := webview.New(true)
	resultFile := filepath.Join(filepath.Dir(exe), "result.txt")
	w.SetTitle("WebView Example")
	w.SetSize(800, 600, webview.HintNone)
	initFile := "static/init.js"
	initBytes, _ := staticFS.ReadFile(initFile)
	indexFile := "static/index.html"
	indexBytes, _ := staticFS.ReadFile(indexFile)
	w.Bind("response", func(s string) {
		w.Dispatch(func() {
			err := os.WriteFile(resultFile, []byte("[test001] " + s), 0666)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		})
	})
	w.Init(string(initBytes))
	w.Navigate("data:text/html," + url.PathEscape(string(indexBytes)))
	w.Run()
}

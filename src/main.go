package main

import (
	"os"
	"time"
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

	// Log Setting
	t := time.Now()
	os.Mkdir(filepath.Join(filepath.Dir(exe), "log"), 0777)
	logFilePath := filepath.Join(filepath.Dir(exe), "log", t.Format("20060102") + ".log")
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		os.Exit(1)
	}
	defer func() {
		logFile.Close()
	}()

	// Webview Setting
	w := webview.New(true)
	w.SetTitle("WebView Example")
	w.SetSize(800, 600, webview.HintNone)
	initFile := filepath.Join("static", "init.js")
	initBytes, _ := staticFS.ReadFile(initFile)
	indexFile := filepath.Join("static", "index.html")
	indexBytes, _ := staticFS.ReadFile(indexFile)
	w.Bind("log", func(s string) {
		w.Dispatch(func() {
			t := time.Now()
			logFile.WriteString("[" + t.Format("2006/01/02 15:04:05") + "] " + s + "\n")
			if err != nil {
				os.Exit(1)
			}
		})
	})
	w.Init(string(initBytes))
	w.Navigate("data:text/html," + url.PathEscape(string(indexBytes)))
	w.Run()
}

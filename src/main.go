package main

import (
	"os"
	"fmt"
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
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		os.Exit(1)
	}
	defer func() {
		logFile.Close()
	}()

	// Memo Setting
	memoFilePath := filepath.Join(filepath.Dir(exe), "log", "memo.log")

	// Webview Setting
	w := webview.New(true)
	w.SetTitle("WebView Example")
	w.SetSize(800, 600, webview.HintNone)
	initFilePath := filepath.Join("static", "init.js")
	initBytes, err := staticFS.ReadFile(initFilePath)
	if err != nil {
		os.Exit(1)
	}
	indexFilePath := filepath.Join("static", "index.html")
	indexBytes, err := staticFS.ReadFile(indexFilePath)
	if err != nil {
		os.Exit(1)
	}
	w.Bind("log", func(s string) {
		w.Dispatch(func() {
			write(logFile, "[" + time.Now().Format("2006/01/02 15:04:05") + "] " + s + "\n")
		})
	})
	w.Bind("navigate", func(url string) {
		w.Navigate(url)
	})
	w.Bind("local", func(html string) {
		htmlFilePath := filepath.Join("static", html)
		htmlBytes, err := staticFS.ReadFile(htmlFilePath)
		if err != nil {
			w.Dispatch(func() {
				write(logFile, "[" + time.Now().Format("2006/01/02 15:04:05") + "] " + fmt.Sprint(err) + "\n")
			})
		} else {
			w.Navigate("data:text/html," + url.PathEscape(string(htmlBytes)))
		}
	})
	w.Bind("write", func(s string) {
		w.Dispatch(func() {
			os.WriteFile(memoFilePath, []byte(s), os.ModePerm)
		})
	})
	w.Bind("read", func()(string) {
		memoBytes, err := os.ReadFile(memoFilePath)
		if err != nil {
			write(logFile, "[" + time.Now().Format("2006/01/02 15:04:05") + "] " + fmt.Sprint(err) + "\n")
			return ""
		} else {
			return string(memoBytes)
		}
	})
	w.Init(string(initBytes))
	w.Navigate("data:text/html," + url.PathEscape(string(indexBytes)))
	w.Run()
}

func write(f *os.File, s string) {
	f.WriteString(s)
}

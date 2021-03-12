package main

import (
	"os"
	"fmt"
	"time"
	"bytes"
	"embed"
	"regexp"
	"strings"
	"net/url"
	"html/template"
	"path/filepath"
	"github.com/webview/webview"
)

type TemplateData struct {
	Links []LinkData `json:"links"`
}

type LinkData struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

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

	// Template Data Setting
	var linkDataList []LinkData
	dataBytes, err := staticFS.ReadFile(filepath.Join("static", "data.csv"))
	if err != nil {
		writeLog(logFile, fmt.Sprint(err))
		os.Exit(1)
	}
	for _, w := range regexp.MustCompile("[\n]").Split(string(dataBytes), -1) {
		x := regexp.MustCompile("[,]").Split(w, -1)
		if len(x) > 1 {
			linkDataList = append(linkDataList, LinkData{Title: x[0], Url: x[1]})
		}
	}
	templateData := TemplateData{Links: linkDataList}

	// Webview Setting
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("WebView Example")
	w.SetSize(800, 600, webview.HintNone)

	// Webview Bind
	w.Bind("log", func(s string) {
		w.Dispatch(func() {
			writeLog(logFile, s)
		})
	})
	w.Bind("navigate", func(url string) {
		w.Navigate(url)
	})
	w.Bind("local", func(tpl string) {
		w.Navigate(getHtmlString(filepath.Join("static", tpl), templateData))
	})
	w.Bind("write", func(s string) {
		w.Dispatch(func() {
			os.WriteFile(memoFilePath, []byte(s), os.ModePerm)
		})
	})
	w.Bind("read", func()(string) {
		memoBytes, err := os.ReadFile(memoFilePath)
		if err != nil {
			writeLog(logFile, fmt.Sprint(err))
			return ""
		} else {
			return string(memoBytes)
		}
	})
	w.Bind("debug", func() {
		w.Dispatch(func() {
			logBytes, err := os.ReadFile(logFilePath)
			if err != nil {
				writeLog(logFile, fmt.Sprint(err))
			} else {
				for _, v := range regexp.MustCompile("[\n]").Split(string(logBytes), -1) {
					v_ := strings.TrimSpace(v)
					if len(v_) > 0 {
						w.Eval("console.log('" + v_ + "');")
					}
				}
			}
		})
	})

	// Webview Init
	initFilePath := filepath.Join("static", "init.js")
	initBytes, err := staticFS.ReadFile(initFilePath)
	if err != nil {
		os.Exit(1)
	}
	w.Init(string(initBytes))

	// Webview Navigate
	w.Navigate(getHtmlString(filepath.Join("static", "index.tpl"), templateData))
	w.Run()
}

func writeLog(f *os.File, s string) {
	f.WriteString("[" + time.Now().Format("2006/01/02 15:04:05") + "] " + s + "\n")
}

func getHtmlString(templateFilePath string, data TemplateData) string {
	tmp := template.New("tmp")
	var buf bytes.Buffer
	funcMap := template.FuncMap{
		"safehtml": func(text string) template.HTML { return template.HTML(text) },
	}
	tmp = template.Must(template.New("").Funcs(funcMap).ParseFS(staticFS, templateFilePath, filepath.Join("static", "css.tpl"), filepath.Join("static", "img.tpl")))
	if err := tmp.ExecuteTemplate(&buf, "base", data); err != nil {
		os.Exit(1)
	}
	return "data:text/html," + url.PathEscape(buf.String())
}

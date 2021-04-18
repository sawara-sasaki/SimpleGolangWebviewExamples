package main

import (
	"io"
	"os"
	"fmt"
	"time"
	"bytes"
	"embed"
	"regexp"
	"strings"
	"net/url"
	"net/http"
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

var logFile *os.File
var exe string
var err error

func main() {
	exe, err = os.Executable()
	if err != nil {
		os.Exit(1)
	}

	// Log Setting
	t := time.Now()
	os.Mkdir(filepath.Join(filepath.Dir(exe), "log"), 0777)
	logFilePath := filepath.Join(filepath.Dir(exe), "log", t.Format("20060102") + ".log")
	logFile, err = os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		os.Exit(1)
	}
	defer func() {
		logFile.Close()
	}()

	// Download Setting
	os.Mkdir(filepath.Join(filepath.Dir(exe), "download"), 0777)

	// Memo Setting
	memoFilePath := filepath.Join(filepath.Dir(exe), "log", "memo.log")

	// Bookmark Setting
	bookmarkFilePath := filepath.Join(filepath.Dir(exe), "log", "bookmark.log")

	// Template Data Setting
	var linkDataList []LinkData
	dataBytes, err := staticFS.ReadFile(filepath.Join("static", "data.csv"))
	if err != nil {
		writeLog("< Error > " + fmt.Sprint(err))
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
		go func() {
			writeLog(s)
		}()
	})
	w.Bind("navigate", func(url string) {
		w.Navigate(url)
	})
	w.Bind("local", func(tpl string) {
		w.Navigate(getHtmlString(filepath.Join("static", tpl), templateData))
	})
	w.Bind("write", func(s string) {
		go func() {
			os.WriteFile(memoFilePath, []byte(s), os.ModePerm)
		}()
	})
	w.Bind("read", func()(string) {
		memoBytes, err := os.ReadFile(memoFilePath)
		if err != nil {
			writeLog("< Error > " + fmt.Sprint(err))
			return ""
		} else {
			return string(memoBytes)
		}
	})
	w.Bind("debug", func() {
		go func() {
			logBytes, err := os.ReadFile(logFilePath)
			if err != nil {
				writeLog("< Error > " + fmt.Sprint(err))
			} else {
				for _, v := range regexp.MustCompile("[\n]").Split(string(logBytes), -1) {
					v_ := strings.TrimSpace(v)
					if len(v_) > 0 {
						w.Dispatch(func() {
							w.Eval("console.log('" + v_ + "');")
						})
					}
				}
			}
		}()
	})
	w.Bind("saveSource", func(url string, source string) {
		go func() {
			sourceFilePath := filepath.Join(filepath.Dir(exe), "log", "source_" + getDomain(url) + "_" + time.Now().Format("20060102150405") + ".log")
			os.WriteFile(sourceFilePath, []byte(source), os.ModePerm)
		}()
	})
	w.Bind("saveCookie", func(url string, cookie string) {
		go func() {
			cookieFilePath := filepath.Join(filepath.Dir(exe), "log", "cookie_" + getDomain(url) + "_" + time.Now().Format("20060102150405") + ".log")
			os.WriteFile(cookieFilePath, []byte(cookie), os.ModePerm)
		}()
	})
	w.Bind("download", func(url string) {
		go func() {
			err := zipDownload(url)
			if err != nil {
				writeLog("< Download Fail > " + url)
				writeLog("< Error > " + fmt.Sprint(err))
			}
		}()
	})
	w.Bind("onAddBookmark", func(url string) {
		go func() {
			bookmarkFile, err := os.OpenFile(bookmarkFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				os.Exit(1)
			}
			defer func() {
				bookmarkFile.Close()
			}()
			bookmarkFile.WriteString(url + "\n")
		}()
	})
	w.Bind("onShowBookmark", func() {
		go func() {
			bookmarkBytes, err := os.ReadFile(bookmarkFilePath)
			if err != nil {
				writeLog("< Error > " + fmt.Sprint(err))
			} else {
				var bookmarkDataList []LinkData
				for _, v := range regexp.MustCompile("[\n]").Split(string(bookmarkBytes), -1) {
					v_ := strings.TrimSpace(v)
					if len(v_) > 0 {
						bookmarkDataList = append(bookmarkDataList, LinkData{Title: v_, Url: v_})
					}
				}
				templateData := TemplateData{Links: bookmarkDataList}
				w.Navigate(getHtmlString(filepath.Join("static", "links.tpl"), templateData))
			}
		}()
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

func writeLog(s string) {
	logFile.WriteString("[" + time.Now().Format("2006/01/02 15:04:05") + "] " + s + "\n")
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

func getDomain(url string) string {
	w := regexp.MustCompile("[/]").Split(url, -1)
	if len(url) < 2 {
		return ""
	}
	return w[2]
}

func getFileName(url string) string {
	w := regexp.MustCompile("[/]").Split(url, -1)
	return w[len(w) - 1]
}

func zipDownload(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept-Encoding", "gzip")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	zipFilePath := filepath.Join(filepath.Dir(exe), "download", time.Now().Format("20060102150405") + "_" + getFileName(url))
	os.WriteFile(zipFilePath, body, os.ModePerm)
	return nil
}

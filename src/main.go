package main

import (
	"os"
	"net/http"
	"path/filepath"
	"github.com/webview/webview"
)

const port string = "8080"

func main() {
	exe, err := os.Executable()
	if err != nil {
		os.Exit(1)
	}

	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			var path string
			if len(r.URL.Path[1:]) == 0 {
				path = filepath.Join(filepath.Dir(exe), "static", "index.html")
			} else {
				path = filepath.Join(filepath.Dir(exe), "static", r.URL.Path[1:])
			}
			http.ServeFile(w, r, path)
		})
		http.ListenAndServe(":" + port, nil)
	}()

	w := webview.New(true)
	w.SetTitle("WebView Example")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://localhost:" + port)
	w.Run()
}

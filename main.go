package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("assets/"))                // staticというディレクトリにあるファイルがサーブされる
	http.Handle("/static/", http.StripPrefix("/static/", fs)) // URLのプレフィックスをStripPrefixで除外する

	http.ListenAndServe(":8080", nil)
}

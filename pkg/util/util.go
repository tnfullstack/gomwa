package util

import "net/http"

func LocalAssets(s string) {
	filePath := "/" + s + "/"
	fs := http.FileServer(http.Dir(s))
	http.Handle(filePath, http.StripPrefix(filePath, fs))
}

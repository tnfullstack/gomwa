package util

import "net/http"

func LocalAssets(s string) {
	filePath := "/" + s + "/"
	http.Handle(filePath, http.StripPrefix(filePath, http.FileServer(http.Dir(s))))
}

package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/mattes/go-asciibot"
)

// H handler function
func H(w http.ResponseWriter, r *http.Request) {
	robot(w, r)

	f := r.URL.Query().Get("khmer")

	if len(f) > 0 && f == "dev" {
		template(w)
	} else {
		http.Redirect(w, r, "https://khmer.dev/?angkor=dev", 307)
	}
}

func robot(w http.ResponseWriter, r *http.Request) {
	ua := r.Header.Get("User-Agent")

	if strings.Contains(ua, "curl") || strings.Contains(ua, "HTTPie") {
		bot := asciibot.Random()

		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", strconv.Itoa(len(bot)))

		w.Write([]byte(bot))
	}
}

func template(w http.ResponseWriter) {
	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>{ angkor.dev }</title>
		<style>
			body {
				background: #111;
				margin: 0;
				display: flex;
				justify-content: center;
				align-items: center;
				height: 100vh;
			}
			img {
				width: 200px;
				height: auto;
			}
		</style>
	</head>
	<body>
		<img src="/angkor.dev.png" alt="angkor.dev">
	</body>
	</html>
	`

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", strconv.Itoa(len(html)))

	w.Write([]byte(html))
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func main() {
	log.Println("start")
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/", handleRoot)

	http.ListenAndServe(":8080", nil)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	settings := GetZQuery()
	var err error
	if r.URL.Query().Has("response") {
		response := r.URL.Query().Get("response")
		if settings.Response, err = strconv.Atoi(response); err != nil {
			dropError("Response parameter must be a number.", w)
			return
		}
		settings.Body.RequestedReturnCode = settings.Response
	}

	if r.URL.Query().Has("delay") {
		delaystr := r.URL.Query().Get("delay")
		re := regexp.MustCompile(`^(\d+)(s|ms)$`)

		if !re.Match([]byte(delaystr)) {
			dropError("the delay must be given in a format like 1s or 1000ms", w)
			return
		}

		parts := re.FindStringSubmatch(delaystr)
		delay, _ := strconv.ParseInt(parts[1], 10, 64)

		var multiplier time.Duration
		if parts[2] == "s" {
			multiplier = time.Second
		} else {
			multiplier = time.Millisecond
		}

		time.Sleep(time.Duration(delay) * multiplier)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(settings.Response)
	jsonbody, _ := json.Marshal(settings.Body)
	w.Write(jsonbody)
}

func dropError(msg string, w http.ResponseWriter) {
	body := zErr{
		Message: msg,
	}
	jsonbody, _ := json.Marshal(body)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(400)
	w.Write(jsonbody)
}

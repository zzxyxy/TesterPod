package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"zxyxy.net/testerpod/utils"
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

		delay, err := utils.ParseDuration(delaystr)

		if err != nil {
			dropError(err.Error(), w)
			return
		}

		time.Sleep(delay)
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

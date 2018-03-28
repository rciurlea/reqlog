package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	rl := newReqlist()
	http.HandleFunc("/", handler(rl))
	log.Fatal(http.ListenAndServe(":2000", nil))
}

func handler(rl *reqlist) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("ok"))
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println(err)
			return
		}
		defer req.Body.Close()
		rl.add(&reqentry{
			When:    time.Now(),
			URL:     req.URL,
			Method:  req.Method,
			Headers: req.Header,
			Body:    b,
		})
	}
}

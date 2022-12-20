package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, getQuote())
}

func main() {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go simulateLogging(ticker, quit)

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", RequestLogger(mux))
}

func simulateLogging(t *time.Ticker, quit chan struct{}) {
	log.Println("")
	for {
		select {
		case <-t.C:
			time.Sleep(time.Duration(rand.Intn(30)) * time.Second)
			q := getQuote()
			log.Println(q)
		case <-quit:
			t.Stop()
			return
		}
	}
}

func getQuote() string {
	url := "https://get-me-a-quote.herokuapp.com"
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	return string(body)
}

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		targetMux.ServeHTTP(w, r)
		requesterIP := r.RemoteAddr
		log.Println(r.Method, r.RequestURI, requesterIP, time.Since(start))
	})
}

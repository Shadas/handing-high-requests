package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Business(payload Payload, wtag int) error {
	sleep_time := payload.SleepTime
	log.Println("run on worker ", wtag, " ...")
	time.Sleep(time.Duration(sleep_time) * time.Second)
	log.Println("run over...")
	return nil
}

func BusinessHandler1(w http.ResponseWriter, r *http.Request) {
	payload := Payload{}
	byteBody, _ := ioutil.ReadAll(r.Body)
	if ok := json.Unmarshal(byteBody, &payload); ok != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := Business(payload, 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Write([]byte(ret))
	// return
}

func BusinessHandler2(w http.ResponseWriter, r *http.Request) {
	payload := Payload{}
	byteBody, _ := ioutil.ReadAll(r.Body)
	if ok := json.Unmarshal(byteBody, &payload); ok != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	work := Job{Payload: payload}

	JobQueue <- work

	w.WriteHeader(http.StatusOK)
	return
}

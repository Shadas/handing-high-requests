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
	// log.Println("run on worker ", wtag, " ...")
	time.Sleep(time.Duration(sleep_time) * time.Second)
	// log.Println("run over...")
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

	// return
}

func BusinessHandler2(w http.ResponseWriter, r *http.Request) {
	payload := Payload{}
	byteBody, _ := ioutil.ReadAll(r.Body)
	if ok := json.Unmarshal(byteBody, &payload); ok != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("queue length:", len(JobQueue))
	log.Println("queue cap:", cap(JobQueue))
	work := Job{Payload: payload}
	if len(JobQueue) == cap(JobQueue) {
		log.Panicln("full and panic!!!!")
	}
	JobQueue <- work
	// log.Println("queue length:", len(JobQueue))

	w.WriteHeader(http.StatusOK)
	return
}

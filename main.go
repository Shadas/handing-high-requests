package main

import (
	"net/http"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	MaxWorker := 200
	MaxQueue := 1000

	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()
	JobQueue = make(chan Job, MaxQueue)

	mux := http.NewServeMux()
	mux.HandleFunc("/aaa", BusinessHandler1)
	mux.HandleFunc("/bbb", BusinessHandler2)

	http.ListenAndServe(":8099", mux)

}

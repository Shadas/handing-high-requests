package main

import (
	"net/http"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	MaxWorker := 5
	MaxQueue := 5

	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()
	JobQueue = make(chan Job, MaxQueue)

	mux := http.NewServeMux()
	mux.HandleFunc("/aaa", BusinessHandler1)
	mux.HandleFunc("/bbb", BusinessHandler2)

	http.ListenAndServe("127.0.0.1:8099", mux)

}

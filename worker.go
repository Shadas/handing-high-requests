package main

type Payload struct {
	SleepTime int64 `json:"sleep_time"`
}

type Job struct {
	Payload Payload
}

var JobQueue chan Job

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	num        int
	quit       chan bool
}

func NewWorker(workpool chan chan Job, num int) Worker {
	return Worker{
		WorkerPool: workpool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
		num:        num,
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				Business(job.Payload, w.num)
			case <-w.quit:
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func generate(from, to int) <-chan int {
	out := make(chan int)
	go func() {
		for i := from; i < to; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

type Worker struct{}

func (w *Worker) Work(jobs <-chan string, doneJobs chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Started working on job: <<%v>>\n", job)
		time.Sleep(2 * time.Second)
		doneJobs <- strings.ToUpper(job)
	}
}

type WorkerPool struct {
	workers  []*Worker
	jobs     chan string
	doneJobs chan string
	wg       *sync.WaitGroup
}

func (wp *WorkerPool) Start() {
	wp.wg.Add(len(wp.workers))
	for i := 0; i < len(wp.workers); i++ {
		wp.workers[i] = &Worker{}
		go wp.workers[i].Work(wp.jobs, wp.doneJobs, wp.wg)
	}
}

func (wp *WorkerPool) DoJob(job string) {
	wp.jobs <- job
}

func (wp *WorkerPool) DoneQueue() <-chan string {
	return wp.doneJobs
}

func (wp *WorkerPool) Stop() {
	close(wp.jobs)
	wp.wg.Wait()
}

func NewWorkerPool(workersCount int) *WorkerPool {
	return &WorkerPool{
		workers:  make([]*Worker, workersCount),
		jobs:     make(chan string, workersCount),
		doneJobs: make(chan string, workersCount),
		wg:       &sync.WaitGroup{},
	}
}

func main() {
	pool := NewWorkerPool(3)
	pool.Start()

	pool.DoJob("¿")
	pool.DoJob("Amiga")
	pool.DoJob(", ")
	pool.DoJob("cómo")
	pool.DoJob("Te")
	pool.DoJob("Llamas")
	pool.DoJob("?")

	go func() {
		for upperCaseStr := range pool.DoneQueue() {
			fmt.Println("Received ", upperCaseStr)
		}
	}()

	time.Sleep(10 * time.Second)
	pool.Stop()
}

package worker

import (
	"fmt"
	"hash/fnv"
	"log"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// create random string
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// create list of jobs
func CreateJobs(amount int) []string {
	var jobs []string

	for i := 0; i < amount; i++ {
		jobs = append(jobs, RandStringRunes(8))
	}
	return jobs
}

// mimics any type of job that can be run concurrently
func DoWork(word string, id int) {
	h := fnv.New32a()
	h.Write([]byte(word))
	time.Sleep(time.Second)
	fmt.Printf("worker [%d] - created hash [%d] from word [%s]\n", id, h.Sum32(), word)
}

type Work struct {
	Id  int
	Job string
}

type Worker struct {
	ID            int
	WorkerChannel chan chan Work
	Channel       chan Work
	End           chan struct{}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerChannel <- w.Channel
			select {
			case job := <-w.Channel:
				DoWork(job.Job, job.Id)
			case <-w.End:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	log.Printf("worker [%d] is stopping", w.ID)
	close(w.End)
}

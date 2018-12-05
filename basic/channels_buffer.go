package main

import (
	"fmt"
	rand2 "math/rand"
	"sync"
	"time"
)

func initbuff() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 1
	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))
}

//WaitGroup: wait for a collection of Goroutines to finish executing.
func wait_group() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process_wg(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}

func process_wg(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

//  a collection of threads which are waiting for tasks to be assigned
func worker_pool() {

}

type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
}

func sum_digits(number int) int {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit
		number /= 10
	}
	return sum
}
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, sum_digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

func create_worker_pool(num_worker int) {
	var wg sync.WaitGroup
	for i := 0; i < num_worker; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()

	// after wg countdown to 0, close channel results
	close(results)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func allocate(num_job int) {
	for i := 0; i < num_job; i++ {
		rand := rand2.Intn(100)
		job := Job{i, rand}
		jobs <- job
	}
	close(jobs)
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func main() {
	//initbuff()
	//wait_group()

	done := make(chan bool)
	go allocate(2)
	go result(done)
	create_worker_pool(3)
	<-done
}
